package tests

import (
	"context"
	"fmt"
	"math/big"

	"github.com/0xNathanW/aaveV2-liqOSS/contracts"
	testcontracts "github.com/0xNathanW/aaveV2-liqOSS/tests/test-contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Fake mnemonic = tumble mail vacant pelican quick brave gallery wide embark inch crop utility

type FakeAccount struct {
	Address common.Address
	Auth    *bind.TransactOpts
}

var AccountsRaw = [][2]string{
	{"0x72A6Ff1614b9bf35B236b7393Df5628dE78E8A91", "88b49e25316d6dad17e5fd27d9fe5393beadd1a7c8ef04be5637f83505d46a39"},
	{"0x332263E1B275F0CdF291C3e4f2336ab0999DF4Fc", "0c42036f7fa63a1e4eb81dad906ed45bf941bb97531cbbba388d8d12d097dd05"},
	{"0xD964a180A598915D64c90446D23077EF4c4fB492", "db34c39597f79ca539adc07b73ec4c275ec74c20e2ed8db6dc773f348c6365ac"},
	{"0x43c73ED869EDDaf5Fe007115c139887A05661878", "1c4c014d33b5729cdad008f05b7fca9b7204d1085a0ebef152b2f9e9aac06f27"},
	{"0x7582F9E3e13E7161c19334D4BCD31bc4634620CE", "0eda8959e73a2bbc514fccfb590918b809b7e815e8eb071186b3a0f6d72ffbdd"}, // Target
}

var (
	WETH               = common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2") // WETH
	CollateralAddress  = common.HexToAddress("0x514910771AF9Ca656af840dff83E8264EcF986CA") // LINK
	LendingPoolAddress = common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D") // aave lending pool

	dec18 = big.NewInt(1e18)
)

func NewFakeAccount(address string, privKey string) (*FakeAccount, error) {

	key, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return nil, fmt.Errorf("failed to load private key: %v", err)
	}

	auth := bind.NewKeyedTransactor(key)

	return &FakeAccount{
		Address: common.HexToAddress(address),
		Auth:    auth,
	}, nil
}

func InitAccounts() ([]*FakeAccount, error) {
	var accountsList []*FakeAccount
	for _, account := range AccountsRaw {
		fakeAccount, err := NewFakeAccount(account[0], account[1])
		if err != nil {
			return nil, err
		}
		accountsList = append(accountsList, fakeAccount)
	}
	return accountsList, nil
}

func SetupTarget(cli *ethclient.Client, lendingPool *contracts.LendingPool, target *FakeAccount) error {

	ethBal, err := cli.BalanceAt(context.Background(), target.Address, nil)
	if err != nil {
		return err
	}
	ethBal.Div(ethBal, big.NewInt(1e18))
	fmt.Println("Target ETH balance:", ethBal)

	// Exchange for collateral.
	router, err := testcontracts.NewUniswapV2Router02(common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"), cli)
	if err != nil {
		return fmt.Errorf("failed to instantiate uniswap router: %v", err)
	}

	suggestedGas, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get gas price: %v", err)
	}

	txOpts := target.Auth
	txOpts.GasPrice = suggestedGas.Mul(suggestedGas, big.NewInt(50)) // Adjust Gas
	txOpts.Value = big.NewInt(0).Mul(big.NewInt(10), big.NewInt(1e18))
	_, err = router.SwapExactETHForTokens(
		txOpts,        // Txopts.
		big.NewInt(1), // Amount out min.
		[]common.Address{WETH, CollateralAddress}, // Path.
		target.Address, // To.
		big.NewInt(10).Mul(big.NewInt(10), big.NewInt(1e18)), // Deadline.
	)
	if err != nil {
		return fmt.Errorf("transaction failed: %v", err)
	}

	link, err := testcontracts.NewErc20(CollateralAddress, cli)
	if err != nil {
		return err
	}
	bal, err := link.BalanceOf(nil, target.Address)
	if err != nil {
		return err
	}
	fmt.Println("LINK balance:", ethVal(bal))
	ethBal, err = cli.BalanceAt(context.Background(), target.Address, nil)
	if err != nil {
		return err
	}
	ethBal.Div(ethBal, big.NewInt(1e18))
	fmt.Println("Target ETH balance:", ethBal)

	txOpts.Value = big.NewInt(0) // Reset value.
	// Approve LendingPool to spend LINK.
	_, err = link.Approve(txOpts, LendingPoolAddress, new(big.Int).Mul(big.NewInt(20), dec18))
	if err != nil {

		return err
	}
	_, err = link.Approve(txOpts, common.HexToAddress("0xC6845a5C768BF8D7681249f8927877Efda425baf"), new(big.Int).Mul(big.NewInt(20), dec18))
	if err != nil {
		return err
	}
	fmt.Println("LINK approved")

	// Print allowance
	allowance, err := link.Allowance(nil, target.Address, LendingPoolAddress)
	if err != nil {
		return err
	}
	fmt.Println("LINK allowance:", allowance.String())

	// Deposit LINK.
	_, err = lendingPool.Deposit(
		txOpts,
		common.HexToAddress("0x514910771AF9Ca656af840dff83E8264EcF986CA"),
		big.NewInt(0).Mul(big.NewInt(10), big.NewInt(1e18)), // 10 LINK.
		target.Address,
		0,
	)
	if err != nil {
		return fmt.Errorf("deposit failed: %v", err)
	}

	// Set Link as collateral.
	_, err = lendingPool.SetUserUseReserveAsCollateral(
		txOpts,
		CollateralAddress,
		true,
	)
	if err != nil {
		return fmt.Errorf("set user use reserve as collateral failed: %v", err)
	}

	printAAVEuserData(lendingPool, target)

	return nil
}

func printAAVEuserData(lp *contracts.LendingPool, target *FakeAccount) error {

	userData, err := lp.GetUserAccountData(nil, target.Address)
	if err != nil {
		return err
	}

	fmt.Println("User data:")
	fmt.Println("\tTotal collateral:", new(big.Int).Div(userData.TotalCollateralETH, dec18).String())
	fmt.Println("\tTotal debt:", new(big.Int).Div(userData.TotalDebtETH, dec18).String())
	fmt.Println("\tAvailable borrow:", new(big.Int).Div(userData.AvailableBorrowsETH, dec18).String())
	fmt.Println("\tHealth factor:", new(big.Int).Div(userData.HealthFactor, dec18).String())
	return nil
}

func ethVal(num *big.Int) int {
	return int(num.Div(num, big.NewInt(1e18)).Int64())
}
