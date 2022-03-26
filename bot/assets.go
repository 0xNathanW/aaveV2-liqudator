package bot

import (
	"fmt"
	"math/big"

	crt "github.com/0xNathanW/aaveV2-liqOSS/contracts"
	"github.com/ethereum/go-ethereum/common"
)

type assetsData struct {
	reserveAssets map[string]*reserveToken
	reverseAssets map[common.Address]*reserveToken
	aTokens       map[string]common.Address
}

type reserveToken struct {
	address  common.Address
	name     string
	decimals *big.Int
	liqBonus *big.Int
	liqThres *big.Int
}

func getAssetData(dataProv *crt.AaveProtocolDataProvider) (*assetsData, error) {

	reserves := make(map[string]*reserveToken)
	reverse := make(map[common.Address]*reserveToken)
	aTokens := make(map[string]common.Address)

	reservesRaw, err := dataProv.GetAllReservesTokens(nil)
	if err != nil {
		return nil, fmt.Errorf("could not get reserves: %w", err)
	}

	aTokensRaw, err := dataProv.GetAllATokens(nil)
	if err != nil {
		return nil, err
	}

	for _, r := range reservesRaw {

		config, err := dataProv.GetReserveConfigurationData(nil, r.TokenAddress)
		if err != nil {
			return nil, fmt.Errorf("could not get reserve config for %s: %w", r.Symbol, err)
		}

		asset := &reserveToken{
			address:  r.TokenAddress,
			name:     r.Symbol,
			decimals: config.Decimals,
			liqBonus: config.LiquidationBonus,
			liqThres: config.LiquidationThreshold,
		}

		reserves[r.Symbol] = asset
		reverse[r.TokenAddress] = asset
	}

	for _, a := range aTokensRaw {
		aTokens[a.Symbol] = a.TokenAddress
	}

	fmt.Println("Asset data retrieved.")
	return &assetsData{
		reserveAssets: reserves,
		aTokens:       aTokens,
	}, nil
}

// Prints token info in a readable format.
func (t *assetsData) prettyAssetData() {

	for _, r := range t.reserveAssets {
		fmt.Printf("%s: %s\n", r.name, r.address.String())
		fmt.Println("\tDecimals:", r.decimals.String())
		fmt.Println("\tLiquidation Bonus:", r.liqBonus.String())
		fmt.Println("\tLiquidation Threshold:", r.liqThres.String())
	}

	fmt.Printf("\n\nA Tokens\n")

	for sym, a := range t.aTokens {
		fmt.Printf("%s: %s\n", sym, a.String())
	}
}
