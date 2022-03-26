package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/0xNathanW/aaveV2-liqOSS/contracts"
	testcontracts "github.com/0xNathanW/aaveV2-liqOSS/tests/test-contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	//"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

type Frame struct {
	client    *ethclient.Client
	address   common.Address
	txOpts    *bind.TransactOpts
	contracts *testContractData
}

type testContractData struct {
	v3SwapRouter *contracts.SwapRouter
	v2SwapRouter *testcontracts.UniswapV2Router02
	lendingPool  *contracts.LendingPool
	daiLinkPool  *testcontracts.Pool
	usdc         *testcontracts.Erc20
	weth         *testcontracts.Erc20
	dai          *testcontracts.Erc20
	link         *testcontracts.Erc20
}

var (
	kovanUSDC   = common.HexToAddress("0x0af08696cb51e81456dc0a1dee7f8bfad8d82a22") // Risdale USDC.
	kovanWETH   = common.HexToAddress("0x4470e84ce5c1fe89ea35a19560b6f2b1ed71507f")
	kovanDai    = common.HexToAddress("0xff795577d9ac8bd7d90ee22b6c1703490b6512fd") // Using AAVE Dai.
	kovanLink   = common.HexToAddress("0xa36085F69e2889c224210F603D836748e7dC0088")
	kovanV3Swap = common.HexToAddress("0xE592427A0AEce92De3Edee1F18E0157C05861564")
	DaiLinkPool = common.HexToAddress("0x6dd927415800150397cf8868668ba31fd3716481")
	kovanV2Swap = common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")
	kovanLP     = common.HexToAddress("0xE0fBa4Fc209b4948668006B2bE61711b7f465bAe")
)

func main() {

	var (
		endPoint = os.Getenv("KOVAN_RPC_URL")
		address  = os.Getenv("KOVAN_ADDRESS")
		privKey  = os.Getenv("KOVAN_PRIV_KEY")
	)

	cli, err := ethclient.Dial(endPoint)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Kovan.")

	p, err := crypto.HexToECDSA(privKey)
	if err != nil {
		log.Fatal(err)
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(p, big.NewInt(42))
	if err != nil {
		log.Fatal(err)
	}

	gas, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	_ = gas
	txOpts.GasLimit = uint64(3000000)

	fmt.Printf("\n%++v\n", txOpts)

	contracts, err := getContracts(cli)
	if err != nil {
		log.Fatal(err)
	}

	frame := &Frame{
		client:    cli,
		address:   common.HexToAddress(address),
		txOpts:    txOpts,
		contracts: contracts,
	}

	frame.swapTest()
}

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		fmt.Println("Error loading .env file")
	}
}

func (f *Frame) swapTest() error {

	if err := f.printAccoutBal(); err != nil {
		return err
	}

	// -- Approval -- //
	time.Sleep(time.Second * 3)
	_, err := f.contracts.dai.Approve( // approve router for dai.
		f.txOpts,
		kovanV3Swap,
		big.NewInt(0).Mul(big.NewInt(100), big.NewInt(1e18)), // 100 Dai.
	)
	if err != nil {
		return fmt.Errorf("failed to approve swap router: %v", err)
	}

	time.Sleep(time.Second * 5)
	allowedRout, err := f.contracts.dai.Allowance(nil, f.address, kovanV3Swap)
	if err != nil {
		return fmt.Errorf("failed to get allowance: %v", err)
	}

	fmt.Printf("Router allowed to spend: %s\n", allowedRout.String())
	time.Sleep(time.Second * 5)

	// -- Transaction -- //
	tx1, err := f.contracts.v3SwapRouter.ExactOutputSingle(
		f.txOpts,
		contracts.ISwapRouterExactOutputSingleParams{
			TokenIn:           kovanDai,                                             // DAI.
			TokenOut:          kovanLink,                                            // LINK.
			Fee:               big.NewInt(3000),                                     // 0.3% fee.
			Recipient:         f.address,                                            // Recipient us.
			Deadline:          big.NewInt(time.Now().Unix() + 3600),                 // 1 hour.
			AmountOut:         big.NewInt(1e18),                                     // 1 LINK.
			AmountInMaximum:   big.NewInt(0).Mul(big.NewInt(200), big.NewInt(1e18)), // 200 DAI.
			SqrtPriceLimitX96: big.NewInt(0),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to get exact output single: %v", err)
	}
	prettyTx(tx1)

	recepit, err := bind.WaitMined(context.Background(), f.client, tx1)
	if err != nil {
		return fmt.Errorf("failed to wait for mined: %v", err)
	}

	fmt.Printf("\nTransaction mined: %s\n", recepit.TxHash.String())
	if recepit.Status == 1 {
		fmt.Println("Transaction status: success")
	}
	if recepit.Status == 0 {
		fmt.Println("Transaction status: failed")
	}

	return nil
}

func testAAVEflash() {

}

func (f *Frame) printAccoutBal() error {

	usdcBalance, err := f.contracts.usdc.BalanceOf(nil, f.address)
	if err != nil {
		return err
	}

	wethBalance, err := f.contracts.weth.BalanceOf(nil, f.address)
	if err != nil {
		return err
	}

	daiBalance, err := f.contracts.dai.BalanceOf(nil, f.address)
	if err != nil {
		return err
	}

	linkBalance, err := f.contracts.link.BalanceOf(nil, f.address)
	if err != nil {
		return err
	}

	fmt.Println("\n=== Account Balances ===")
	fmt.Printf("DAI: %s\n", usdcBalance.String())
	fmt.Printf("WETH: %s\n", wethBalance.String())
	fmt.Printf("DAI: %s\n", daiBalance.String())
	fmt.Printf("LINK: %s\n", linkBalance.String())
	fmt.Println("=======================")
	return nil
}

func getContracts(cli *ethclient.Client) (*testContractData, error) {

	router, err := contracts.NewSwapRouter(kovanV3Swap, cli)
	if err != nil {
		return nil, err
	}
	usdc, err := testcontracts.NewErc20(kovanUSDC, cli)
	if err != nil {
		return nil, err
	}
	weth, err := testcontracts.NewErc20(kovanWETH, cli)
	if err != nil {
		return nil, err
	}
	pool, err := testcontracts.NewPool(DaiLinkPool, cli)
	if err != nil {
		return nil, err
	}
	link, err := testcontracts.NewErc20(kovanLink, cli)
	if err != nil {
		return nil, err
	}
	dai, err := testcontracts.NewErc20(kovanDai, cli)
	if err != nil {
		return nil, err
	}
	V2SwapRouter, err := testcontracts.NewUniswapV2Router02(kovanV2Swap, cli)
	if err != nil {
		return nil, err
	}
	lendingPool, err := contracts.NewLendingPool(kovanLP, cli)
	if err != nil {
		return nil, err
	}

	return &testContractData{
		v3SwapRouter: router,
		usdc:         usdc,
		weth:         weth,
		daiLinkPool:  pool,
		link:         link,
		dai:          dai,
		v2SwapRouter: V2SwapRouter,
		lendingPool:  lendingPool,
	}, nil
}

func etherscanLink(txHash string) string {
	return "https://kovan.etherscan.io/tx/" + txHash + " "
}

func prettyTx(tx *types.Transaction) {
	fmt.Println("\n=== Transaction Details ===")
	fmt.Printf("Hash: %v\n", tx.Hash().Hex())
	fmt.Printf("Value: %v\n", tx.Value().String())
	fmt.Printf("Gas: %v\n", tx.Gas())
	fmt.Printf("GasPrice: %v\n", tx.GasPrice().String())
	fmt.Printf("Nonce: %v\n", tx.Nonce())
	fmt.Println("Etherscan Link:", etherscanLink(tx.Hash().Hex()))
	fmt.Println("=======================")
}
