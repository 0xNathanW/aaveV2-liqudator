package contracts

import (
	"fmt"
	"liquidator/contracts/aaveV2"
	"liquidator/contracts/chainlink"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Contracts struct {
	AaveV2 *AaveV2
}

type Chainlink struct {
	Aggregator *chainlink.Aggregator
}

type AaveV2 struct {
	LendingPool  *aaveV2.LendingPoolV2
	DataProvider *aaveV2.DataProviderV2
}

func RetrieveContracts(cli *ethclient.Client) (*Contracts, error) {

	var (
		dataProviderAddress = common.HexToAddress("0xd05e3E715d945B59290df0ae8eF85c1BdB684744")
	)

	addressProvider, err := aaveV2.NewAddressProviderV2(dataProviderAddress, cli)
	if err != nil {
		return nil, fmt.Errorf("failed to create address provider v2: %v", err)
	}

	lendingPoolAddress, err := addressProvider.GetLendingPool(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get lending pool v2 address: %v", err)
	}

	lendingPool, err := aaveV2.NewLendingPoolV2(lendingPoolAddress, cli)
	if err != nil {
		return nil, fmt.Errorf("failed to create lending pool v2: %v", err)
	}

	dataProviderAddress, err = addressProvider.GetAddress(nil, [32]byte{0x01})
	if err != nil {
		return nil, fmt.Errorf("failed to get data provider v2 address: %v", err)
	}

	dataProvider, err := aaveV2.NewDataProviderV2(dataProviderAddress, cli)
	if err != nil {
		return nil, fmt.Errorf("failed to create data provider v2: %v", err)
	}

	return &Contracts{
		AaveV2: &AaveV2{
			LendingPool:  lendingPool,
			DataProvider: dataProvider,
		},
	}, nil
}
