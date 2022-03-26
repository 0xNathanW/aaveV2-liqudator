package bot

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/0xNathanW/aaveV2-liqOSS/contracts"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	dec "github.com/shopspring/decimal"
)

var (
	LendingPoolABI, _   = abi.JSON(strings.NewReader(string(contracts.LendingPoolABI)))
	lendingPoolAddr     = common.HexToAddress("0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9")
	FlashLoanPremium, _ = dec.NewFromString("0.09")
)

func (b *Bot) attemptLiq(user common.Address) error {

	var (
		debtAsset       common.Address
		collateralAsset common.Address
		debtAmount      *big.Int = big.NewInt(0)
		collateralBonus *big.Int = big.NewInt(0)

		debtString       string
		collateralString string
	)

	for sym, token := range b.assets.reserveAssets {

		// Retieve user info concerning current token.
		userReserveData, err := b.contracts.DataProvider.GetUserReserveData(nil, token.address, user)
		if err != nil {
			log.Printf("could not get user,%s ,reserve data for %s\n", user.Hex(), token.name)
			log.Println(err)
			continue
		}

		fmt.Println("\n", sym)

		// debtToCover = (userStableDebt + userVariableDebt) * LiquidationCloseFactorPercent
		totalDebt := big.NewInt(0).Add(userReserveData.CurrentStableDebt, userReserveData.CurrentVariableDebt)
		debtToCover := big.NewInt(0).Div(totalDebt, big.NewInt(2))

		// Continue if the user has no liquidatable debt, or collateral in this asset.
		if userReserveData.CurrentATokenBalance.Cmp(big.NewInt(0)) == 0 &&
			debtToCover.Cmp(big.NewInt(0)) == 0 {

			fmt.Println("\tno liquidatable debt or collateral")

			continue
		}

		// Retrieve price from oracle.
		price, err := b.contracts.PriceOracle.GetAssetPrice(nil, token.address)
		if err != nil {
			return fmt.Errorf("could not get price for %s: %w", token.name, err)
		}

		// debtToCover in terms of wei.
		debtToCover.Mul(debtToCover, price)

		var assetBonus *big.Int
		// If the user has asset as collateral, calculate the collateral bonus.
		if userReserveData.UsageAsCollateralEnabled && userReserveData.CurrentATokenBalance.Cmp(big.NewInt(0)) > 0 {

			// collat bonus = (userCollatBalance * price) * CollateralBonusPercent
			bonusDec := dec.NewFromBigInt(token.liqBonus, -4)
			assetBonus = dec.NewFromBigInt(userReserveData.CurrentATokenBalance, 0).
				Mul(bonusDec).
				Mul(dec.NewFromBigInt(price, 0)).
				BigInt()
			fmt.Println("\tCollat bonus:", BigPretty(assetBonus, 18))
		}

		// Print values.
		fmt.Println("\n\tDecimals:", token.decimals.Uint64())
		fmt.Println("\tAsset liq bonus:", token.liqBonus.String())
		fmt.Println("\tDebt to cover:", BigPretty(debtToCover, token.decimals.Uint64()))
		fmt.Println("\tCollateral bonus:", BigPretty(assetBonus, 0))

		// Compare with current stats.  If better, update.
		if debtToCover.Cmp(debtAmount) > 0 {
			debtAsset = token.address
			debtAmount = debtToCover
			debtString = sym
		}
		if assetBonus.Cmp(collateralBonus) >= 0 {
			collateralAsset = token.address
			collateralBonus = assetBonus
			collateralString = sym
		}
	}

	if debtAmount.Cmp(big.NewInt(0)) == 0 {
		return fmt.Errorf("no liquidatable debt found for %s, despite being under health factor", user.Hex())
	}

	fmt.Printf("\nAttempting to liquidate %s:"+
		"\n\tDebt asset: %s"+
		"\n\tDebt amount: %s"+
		"\n\tCollateral asset: %s"+
		"\n\tCollateral bonus: %s\n",
		user.Hex(), debtString, BigPretty(debtAmount, b.assets.reserveAssets[debtString].decimals.Uint64()),
		collateralString, BigPretty(collateralBonus, b.assets.reserveAssets[collateralString].decimals.Uint64()))

	suggestedGasPrice, err := b.client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("could not get gas price: %w", err)
	}

	// Pack extra data.
	params := byteEncodeParams(collateralAsset.Bytes(), user.Bytes())

	// ABI pack whole method.
	packedParams, err := LendingPoolABI.Pack(
		"flashLoan",
		b.contractAddress,
		[]common.Address{debtAsset},
		[]*big.Int{debtAmount},
		[]*big.Int{big.NewInt(0)},
		b.contractAddress,
		params,
		uint16(0),
	)
	if err != nil {
		return fmt.Errorf("could not pack params: %w", err)
	}
	fmt.Println("\nPacked params:", hex.EncodeToString(packedParams))

	// Calculate gas usage.
	usedGas, err := b.client.EstimateGas(
		context.Background(),
		ethereum.CallMsg{
			From:     b.txOpts.From,
			To:       &lendingPoolAddr,
			Gas:      0, // max gas, 0 = near-infinite.
			GasPrice: suggestedGasPrice,
			Value:    big.NewInt(0),
			Data:     packedParams,
		},
	)
	if err != nil {
		return fmt.Errorf("could not estimate gas: %w", err)
	}

	// Check if liquidation is worth it.  Both values are in wei.
	var totalCost *big.Int
	totalCost.Mul(big.NewInt(int64(usedGas)), suggestedGasPrice)
	var totalCostDec = dec.NewFromBigInt(totalCost, 0)
	totalCost = totalCostDec.Mul(FlashLoanPremium).BigInt()

	if totalCost.Cmp(collateralBonus) > 0 {
		log.Println("\nNot worth it, skipping...")
		log.Println("\n\tTotal cost:", totalCost.String())
		log.Println("\n\tCollateral bonus:", collateralBonus.String())
		return nil
	}

	// Calc max gas price.
	var maxGasPrice *big.Int
	maxGasPrice.Div(collateralBonus, big.NewInt(int64(usedGas)))

	txOpts := b.txOpts
	txOpts.GasPrice = suggestedGasPrice
	txOpts.GasLimit = uint64(usedGas)

	// Send the transaction.
	tx, err := b.contracts.LendingPool.FlashLoan(
		txOpts,
		b.contractAddress,
		[]common.Address{debtAsset},
		[]*big.Int{debtAmount},
		[]*big.Int{big.NewInt(0)},
		b.contractAddress,
		params,
		0,
	)
	if err != nil {
		return fmt.Errorf("transaction error: %w", err)
	}

	// Collect receipt.
	receipt, err := bind.WaitMined(context.Background(), b.client, tx)
	if err != nil {
		return fmt.Errorf("transaction mining error: %w", err)
	}
	log.Println("\nTransaction mined:", tenderlyLink(receipt.TxHash.Hex()))

	return nil
}

func byteEncodeParams(params ...[]byte) []byte {
	var result []byte
	for _, param := range params {
		param = common.LeftPadBytes(param, 32)
		result = append(result, param...)
	}
	return result
}

func tenderlyLink(txHash string) string {
	return "https://dashboard.tenderly.co/tx/mainnet/" + txHash
}

func BigPretty(i *big.Int, decimals uint64) string {
	return dec.NewFromBigInt(i, -int32(decimals)).String()
}
