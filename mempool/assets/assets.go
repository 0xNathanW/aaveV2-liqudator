package assets

import (
	"fmt"
	"liquidator/contracts"
	"liquidator/contracts/tokens"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Asset struct {
	Name         string
	Symbol       string
	Address      common.Address
	Decimals     uint8
	AaveV2params *aave2params
}

type aave2params struct {
	LiqBonus     *big.Int
	LiqThreshold *big.Int
}

func FetchAssets(ethCli *ethclient.Client, contracts *contracts.Contracts) (map[common.Address]*Asset, error) {

	assets := make(map[common.Address]*Asset)

	for _, address := range PolyDataFeeds {

		erc20, err := tokens.NewErc20(address, ethCli)
		if err != nil {
			return nil, fmt.Errorf("failed to create ERC20 contract for %s: %v", address.Hex(), err)
		}

		name, err := erc20.Name(nil)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve name for %s: %v", address.Hex(), err)
		}

		symbol, err := erc20.Symbol(nil)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve symbol for %s: %v", address.Hex(), err)
		}

		decimals, err := erc20.Decimals(nil)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve decimals for %s: %v", address.Hex(), err)
		}

		aaveData, err := contracts.AaveV2.DataProvider.GetReserveConfigurationData(nil, address)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve aave data for %s: %v", address.Hex(), err)
		}

		assets[address] = &Asset{
			Name:     name,
			Symbol:   symbol,
			Address:  address,
			Decimals: decimals,
			AaveV2params: &aave2params{
				LiqBonus:     aaveData.LiquidationBonus,
				LiqThreshold: aaveData.LiquidationThreshold,
			},
		}
	}
	return assets, nil
}

// Format symbol to be exactly 4 characters.
func FormatSymbol(name string) string {
	if len(name) != 4 {
		if len(name) > 4 {
			name = name[:4]
		} else {
			name = name + strings.Repeat("#", 4-len(name))
		}
	}
	return name
}
