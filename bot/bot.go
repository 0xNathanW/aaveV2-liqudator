package bot

import (
	"fmt"
	"math/big"

	crt "github.com/0xNathanW/aaveV2-liqOSS/contracts"
	txlistener "github.com/0xNathanW/aaveV2-liqOSS/tx-listener"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/syndtr/goleveldb/leveldb"
)

type Bot struct {
	client          *ethclient.Client
	db              *leveldb.DB
	iterating       bool
	listener        *txlistener.Listener
	contracts       *contractData
	contractAddress common.Address
	assets          *assetsData
	txOpts          *bind.TransactOpts
	err             chan error
}

type contractData struct {
	LendingPool   *crt.LendingPool
	PriceOracle   *crt.AaveOracle
	LendingOracle *crt.LendingRateOracle
	DataProvider  *crt.AaveProtocolDataProvider
	FlashLiq      *crt.FlashLiq
}

func NewBot(endPoint, dbPath, privKey, contractAddress string) (*Bot, error) {

	_ = endPoint
	// Connect to an Ethereum node. - Ganache for test.
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to Ethereum node.")

	// Open the database, or create if non-existant.
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		return nil, fmt.Errorf("could not open db: %w", err)
	}

	contracts, err := getContracts(client, common.HexToAddress(contractAddress))
	if err != nil {
		return nil, err
	}

	assets, err := getAssetData(contracts.DataProvider)
	if err != nil {
		return nil, err
	}

	txOpts, err := buildTxOpts(privKey)
	if err != nil {
		return nil, err
	}

	return &Bot{
		client:          client,
		db:              db,
		listener:        txlistener.NewListener(),
		contracts:       contracts,
		assets:          assets,
		txOpts:          txOpts,
		contractAddress: common.HexToAddress(contractAddress),
		err:             make(chan error),
	}, nil
}

func (b *Bot) shutdown() {
	b.iterating = false
	if b.listener.SubsOn {
		b.listener.Cleanup()
	}
	b.client.Close()
	b.db.Close()
	fmt.Println("Bot shutdown complete.")
}

func getContracts(cli *ethclient.Client, contractAddr common.Address) (*contractData, error) {

	strct := &contractData{}
	addrProvider, err := crt.NewLendingPoolAddressesProvider(
		common.HexToAddress("0xB53C1a33016B2DC2fF3653530bfF1848a515c8c5"), cli,
	)
	if err != nil {
		return nil, fmt.Errorf("could not get lending pool addresses: %w", err)
	}
	fmt.Println("Lending pool addresses provider retrieved.")

	// Get lending pool.
	lpAddress, err := addrProvider.GetLendingPool(nil)
	if err != nil {
		return nil, err
	}
	if strct.LendingPool, err = crt.NewLendingPool(lpAddress, cli); err != nil {
		return nil, err
	}
	fmt.Println("Lending pool contract retrieved.")

	// Get price oracle.
	priceOracleAddr, err := addrProvider.GetPriceOracle(nil)
	if err != nil {
		return nil, err
	}
	if strct.PriceOracle, err = crt.NewAaveOracle(priceOracleAddr, cli); err != nil {
		return nil, err
	}
	fmt.Println("Price oracle contract retrieved.")

	// Get lending oracle.
	lendingOracleAddr, err := addrProvider.GetLendingRateOracle(nil)
	if err != nil {
		return nil, err
	}
	if strct.LendingOracle, err = crt.NewLendingRateOracle(lendingOracleAddr, cli); err != nil {
		return nil, err
	}
	fmt.Println("Lending oracle contract retrieved.")

	// Get data provider.
	dataProvAddress, err := addrProvider.GetAddress(nil, [32]byte{0x01})
	if err != nil {
		return nil, err
	}
	if strct.DataProvider, err = crt.NewAaveProtocolDataProvider(dataProvAddress, cli); err != nil {
		return nil, err
	}
	fmt.Println("Data provider contract retrieved.")

	// Get flash liq.
	// if strct.FlashLiq, err = crt.NewFlashLiq(contractAddr, cli); err != nil {
	// 	return nil, err
	// }

	fmt.Println("Contracts retrieved.")

	return strct, nil
}

func buildTxOpts(privKey string) (*bind.TransactOpts, error) {

	key, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return nil, fmt.Errorf("could not get private key: %w", err)
	}
	txOpts, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1))
	if err != nil {
		return nil, fmt.Errorf("could not get tx opts: %w", err)
	}

	return txOpts, nil
}
