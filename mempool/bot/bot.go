package bot

import (
	"fmt"
	"liquidator/assets"
	"liquidator/contracts"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/syndtr/goleveldb/leveldb"
)

type Bot struct {
	rpcCli    *rpc.Client
	ethCli    *ethclient.Client
	contracts *contracts.Contracts
	assets    map[common.Address]*assets.Asset
	db        *leveldb.DB
	auth      *bind.TransactOpts
	txLog     *log.Logger
}

func NewBot() (*Bot, error) {

	rpcCli, err := rpc.Dial(os.Getenv("POLY_RPC"))
	if err != nil {
		return nil, err
	}

	ethCli := ethclient.NewClient(rpcCli)

	contracts, err := contracts.RetrieveContracts(ethCli)
	if err != nil {
		return nil, err
	}

	assets, err := assets.FetchAssets(ethCli, contracts)
	if err != nil {
		return nil, err
	}

	db, err := leveldb.OpenFile(os.Getenv("DB_PATH"), nil)
	if err != nil {
		return nil, fmt.Errorf("could not open db: %w", err)
	}

	auth, err := newAuth(os.Getenv("PRIV_KEY"))
	if err != nil {
		return nil, err
	}

	f, err := os.OpenFile("logs/tx.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	logger := log.New(f, "", log.LstdFlags)

	return &Bot{
		rpcCli:    rpcCli,
		ethCli:    ethCli,
		contracts: contracts,
		assets:    assets,
		db:        db,
		auth:      auth,
		txLog:     logger,
	}, nil
}

func newAuth(privKey string) (*bind.TransactOpts, error) {

	key, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return nil, fmt.Errorf("could not get private key: %w", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1))
	if err != nil {
		return nil, fmt.Errorf("could not get tx opts: %w", err)
	}

	return auth, nil
}

func (b *Bot) logTx(tx *types.Transaction) {

	b.txLog.Println("----- NEW TX -----")
	b.txLog.Println("hash:", tx.Hash().Hex())
	b.txLog.Println("to:", tx.To().Hex())
	b.txLog.Print("gas:", tx.Gas(), "\n\n")

}
