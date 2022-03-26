package bot

import (
	"bytes"
	"context"
	"fmt"
	"liquidator/assets"
	"liquidator/contracts/chainlink"
	"liquidator/listener"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	dec "github.com/shopspring/decimal"
)

const MAX_WORKERS = 7

type priceUpdate struct {
	asset  common.Address
	change *big.Int
}

func (b *Bot) Run() error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer b.ethCli.Close()
	defer b.db.Close()

	// Setup and run the event listeners.
	eventListener, err := listener.NewEventListener(b.db, b.ethCli, b.contracts)
	if err != nil {
		return fmt.Errorf("could not create event listener: %w", err)
	}
	eventListener.Run(ctx, b.contracts)

	// Retrieve the aggregator ABI.
	aggABI, err := abi.JSON(strings.NewReader(chainlink.AggregatorMetaData.ABI))
	if err != nil {
		return fmt.Errorf("could not get aggregator abi: %w", err)
	}

	newTx := make(chan common.Hash, 20000)
	updates := make(chan *priceUpdate)

	// Subscribe to new transactions.
	sub, err := b.rpcCli.EthSubscribe(
		context.Background(), newTx, "newPendingTransactions",
	)
	if err != nil {
		return fmt.Errorf("could not subscribe to new txs: %w", err)
	}
	defer sub.Unsubscribe()

	// Start workers.
	for i := 0; i < MAX_WORKERS; i++ {
		go b.worker(newTx, updates, &aggABI)
	}

	go func() { // Handle new tx subscription errors.
		for err := range sub.Err() {
			log.Fatal("subscription error:", err) // TODO: fix
		}
	}()

	for update := range updates {

		fmt.Println("\n=========== New Update ===========")
		log.Println("asset:", b.assets[update.asset].Name)
		log.Println("price change:", BigPretty(update.change, uint64(b.assets[update.asset].Decimals)))
		fmt.Println("===================================")

	}

	return nil
}

func (b *Bot) worker(newTx <-chan common.Hash, updates chan<- *priceUpdate, aggAbi *abi.ABI) {
	fmt.Println("worker started..")
	for hash := range newTx {

		tx, pending, err := b.ethCli.TransactionByHash(context.Background(), hash)
		if err != nil || !pending || tx.To() == nil {
			continue
		}

		if _, ok := assets.PolyDataFeeds[*tx.To()]; ok {
			// ↓ sig for "trasmit" function.
			if bytes.Equal(tx.Data()[:4], []byte{201, 128, 117, 57}) {

				update, err := b.calculatePriceChange(tx, aggAbi)
				if err != nil {
					log.Println("could not calculate price change:", err)
					continue
				}

				updates <- update
			}
		}
	}
}

func (b *Bot) calculatePriceChange(tx *types.Transaction, aggABI *abi.ABI) (*priceUpdate, error) {

	b.logTx(tx)

	method, err := aggABI.MethodById(tx.Data()[:4])
	if err != nil {
		return nil, fmt.Errorf("could not get method: %w", err)
	}

	args := make(map[string]interface{})
	if err = method.Inputs.UnpackIntoMap(args, tx.Data()[4:]); err != nil {
		return nil, fmt.Errorf("could not unpack args: %w", err)
	}

	if report, ok := args["_report"].([]byte); ok {
		var observations []*big.Int
		// First 2 bytes32 are rawReportContext and rawObservers.
		// For some reason next two are also taken (num of observations maybe?).
		for i := 4; i < (len(report) / 32); i++ {
			obv := new(big.Int).SetBytes(report[i*32 : (i+1)*32])
			observations = append(observations, obv)
		}
		// Observations are already sorted, new price is median.
		newPrice := observations[len(observations)/2]

		offchainAggregator, err := chainlink.NewAggregator(*tx.To(), b.ethCli)
		if err != nil {
			return nil, fmt.Errorf("could not get offchain aggregator, %s: %w", tx.To(), err)
		}

		// Get current asset price.
		currentPrice, err := offchainAggregator.LatestAnswer(nil)
		if err != nil {
			return nil, fmt.Errorf("could not get asset price, aggregator %s: %w", tx.To(), err)
		}

		priceChange := new(big.Int).Sub(newPrice, currentPrice)

		return &priceUpdate{
			asset:  assets.PolyDataFeeds[*tx.To()],
			change: priceChange,
		}, nil

	} else {
		return nil, fmt.Errorf("could not get _report from args")
	}

}

func BigPretty(i *big.Int, decimals uint64) string {
	return dec.NewFromBigInt(i, -int32(decimals)).String()
}
