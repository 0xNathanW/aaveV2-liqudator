package listener

import (
	"fmt"
	"liquidator/assets"
	"liquidator/contracts/aaveV2"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
)

type aaveV2Events struct {
	borrow   chan *aaveV2.LendingPoolV2Borrow
	deposit  chan *aaveV2.LendingPoolV2Deposit
	withdraw chan *aaveV2.LendingPoolV2Withdraw
	repay    chan *aaveV2.LendingPoolV2Repay
	enable   chan *aaveV2.LendingPoolV2ReserveUsedAsCollateralEnabled
	disable  chan *aaveV2.LendingPoolV2ReserveUsedAsCollateralDisabled

	lendingPool *aaveV2.LendingPoolV2
	assetMap    []string // Asset map mirrors the user asset config bitmask.
}

func initAaveSubs(
	cli *ethclient.Client,
	lp *aaveV2.LendingPoolV2,
	assets map[common.Address]*assets.Asset,
) (*aaveV2Events, []event.Subscription, error) {

	borrow := make(chan *aaveV2.LendingPoolV2Borrow)
	borrowSub, err := lp.WatchBorrow(nil, borrow, nil, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	deposit := make(chan *aaveV2.LendingPoolV2Deposit)
	depositSub, err := lp.WatchDeposit(nil, deposit, nil, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	withdraw := make(chan *aaveV2.LendingPoolV2Withdraw)
	withdrawSub, err := lp.WatchWithdraw(nil, withdraw, nil, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	repay := make(chan *aaveV2.LendingPoolV2Repay)
	repaySub, err := lp.WatchRepay(nil, repay, nil, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	enable := make(chan *aaveV2.LendingPoolV2ReserveUsedAsCollateralEnabled)
	enableSub, err := lp.WatchReserveUsedAsCollateralEnabled(nil, enable, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	disable := make(chan *aaveV2.LendingPoolV2ReserveUsedAsCollateralDisabled)
	disableSub, err := lp.WatchReserveUsedAsCollateralDisabled(nil, disable, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	reservesList, err := lp.GetReservesList(nil)
	if err != nil {
		return nil, nil, err
	}

	assetMap := make([]string, len(reservesList))
	for i, asset := range reservesList {
		assetMap[i] = assets[asset].Symbol
	}

	return &aaveV2Events{
			borrow:   borrow,
			deposit:  deposit,
			withdraw: withdraw,
			repay:    repay,
			enable:   enable,
			disable:  disable,

			lendingPool: lp,
		}, []event.Subscription{
			borrowSub,
			depositSub,
			withdrawSub,
			repaySub,
			enableSub,
			disableSub,
		}, nil
}

func (l *EventListener) listenAaveV2(lp *aaveV2.LendingPoolV2) {

	for {
		select {
		case borrow := <-l.aave.borrow:
			l.handleAave(borrow.OnBehalfOf, lp)
		case deposit := <-l.aave.deposit:
		case withdraw := <-l.aave.withdraw:
		case repay := <-l.aave.repay:
		case enable := <-l.aave.enable:
		case disable := <-l.aave.disable:

		}
	}
}

func (l *EventListener) handleAave(user common.Address, lendingPool *aaveV2.LendingPoolV2) {

	userConfig, err := lendingPool.GetUserConfiguration(nil, user)
	if err != nil {
		log.Println(fmt.Errorf("failed to get user %s config: %v", user.Hex(), err))
		return
	}
	bitmask := userConfig.Data

	for i := bitmask.BitLen() + 1; i >= 0; i -= 2 {

		collateral := bitmask.Bit(i)
		var borrowed uint
		if i-1 < 0 {
			borrowed = 0
		} else {
			borrowed = bitmask.Bit(i - 1)
		}
		name := assets.FormatSymbol(l.aave.assetMap[i])

		if borrowed == 1 {
			if err := l.db.Put(
				[]byte(fmt.Sprintf("%s-:%s", name, user.Hex())),
				[]byte(fmt.Sprintf("%s:%s\n", "aaveV2", bitmask.String())),
				nil,
			); err != nil {
				l.errLog.Println(fmt.Errorf("failed to write to db: %v", err))
			}
		}
		if collateral == 1 {
			if err := l.db.Put(
				[]byte(fmt.Sprintf("%s+:%s", name, user.Hex())),
				[]byte(fmt.Sprintf("%s:%s\n", "aaveV2", bitmask.String())),
				nil,
			); err != nil {
				l.errLog.Println(fmt.Errorf("failed to write to db: %v", err))
			}
		}

	}

}
