package txlistener

import (
	"fmt"

	"github.com/0xNathanW/aaveV2-liqOSS/contracts"
	"github.com/syndtr/goleveldb/leveldb"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Listener struct {
	SubsOn             bool
	Borrow             *borrow
	Deposit            *deposit
	FlashLoan          *flashLoan
	Liquidation        *liquidation
	Rebalance          *rebalance
	Repay              *repay
	ReserveUpdate      *reserveUpdate
	CollateralDisabled *collateralDisabled
	CollateralEnabled  *collateralEnabled
	Swap               *swap
	Withdraw           *withdraw
}

func NewListener() *Listener {
	return &Listener{
		Borrow:             &borrow{Incoming: make(chan *contracts.LendingPoolBorrow)},
		Deposit:            &deposit{Incoming: make(chan *contracts.LendingPoolDeposit)},
		FlashLoan:          &flashLoan{Incoming: make(chan *contracts.LendingPoolFlashLoan)},
		Liquidation:        &liquidation{Incoming: make(chan *contracts.LendingPoolLiquidationCall)},
		Rebalance:          &rebalance{Incoming: make(chan *contracts.LendingPoolRebalanceStableBorrowRate)},
		Repay:              &repay{Incoming: make(chan *contracts.LendingPoolRepay)},
		ReserveUpdate:      &reserveUpdate{Incoming: make(chan *contracts.LendingPoolReserveDataUpdated)},
		CollateralDisabled: &collateralDisabled{Incoming: make(chan *contracts.LendingPoolReserveUsedAsCollateralDisabled)},
		CollateralEnabled:  &collateralEnabled{Incoming: make(chan *contracts.LendingPoolReserveUsedAsCollateralEnabled)},
		Swap:               &swap{Incoming: make(chan *contracts.LendingPoolSwap)},
		Withdraw:           &withdraw{Incoming: make(chan *contracts.LendingPoolWithdraw)},
	}
}

func (l *Listener) initiateSubscriptions(cli *ethclient.Client, lendingPool *contracts.LendingPool) error {

	var err error
	// Subscribe to lending pool events.
	if l.Borrow.Sub, err = lendingPool.WatchBorrow(nil, l.Borrow.Incoming, nil, nil, nil); err != nil {
		return err
	}
	if l.Deposit.Sub, err = lendingPool.WatchDeposit(nil, l.Deposit.Incoming, nil, nil, nil); err != nil {
		return err
	}
	if l.FlashLoan.Sub, err = lendingPool.WatchFlashLoan(nil, l.FlashLoan.Incoming, nil, nil, nil); err != nil {
		return err
	}
	if l.Liquidation.Sub, err = lendingPool.WatchLiquidationCall(nil, l.Liquidation.Incoming, nil, nil, nil); err != nil {
		return err
	}
	if l.Rebalance.Sub, err = lendingPool.WatchRebalanceStableBorrowRate(nil, l.Rebalance.Incoming, nil, nil); err != nil {
		return err
	}
	if l.Repay.Sub, err = lendingPool.WatchRepay(nil, l.Repay.Incoming, nil, nil, nil); err != nil {
		return err
	}
	if l.ReserveUpdate.Sub, err = lendingPool.WatchReserveDataUpdated(nil, l.ReserveUpdate.Incoming, nil); err != nil {
		return err
	}
	if l.CollateralDisabled.Sub, err = lendingPool.WatchReserveUsedAsCollateralDisabled(nil, l.CollateralDisabled.Incoming, nil, nil); err != nil {
		return err
	}
	if l.CollateralEnabled.Sub, err = lendingPool.WatchReserveUsedAsCollateralEnabled(nil, l.CollateralEnabled.Incoming, nil, nil); err != nil {
		return err
	}
	if l.Swap.Sub, err = lendingPool.WatchSwap(nil, l.Swap.Incoming, nil, nil); err != nil {
		return err
	}
	if l.Withdraw.Sub, err = lendingPool.WatchWithdraw(nil, l.Withdraw.Incoming, nil, nil, nil); err != nil {
		return err
	}
	l.SubsOn = true
	return nil
}

// Unsubscribe from all subscriptions.
func (l *Listener) Cleanup() {
	l.Borrow.Sub.Unsubscribe()
	l.Deposit.Sub.Unsubscribe()
	l.FlashLoan.Sub.Unsubscribe()
	l.Liquidation.Sub.Unsubscribe()
	l.Rebalance.Sub.Unsubscribe()
	l.Repay.Sub.Unsubscribe()
	l.ReserveUpdate.Sub.Unsubscribe()
	l.CollateralDisabled.Sub.Unsubscribe()
	l.CollateralEnabled.Sub.Unsubscribe()
	l.Swap.Sub.Unsubscribe()
	l.Withdraw.Sub.Unsubscribe()
}

// Runs loop for listening to events and stores users in database.
func (l *Listener) Listen(
	e chan<- error,
	cli *ethclient.Client,
	db *leveldb.DB,
	lendingPool *contracts.LendingPool,
) {

	if err := l.initiateSubscriptions(cli, lendingPool); err != nil {
		e <- fmt.Errorf("failed to initiate subscriptions: %w", err)
		return
	}

	go func() {
		if err := l.listenSubError(); err != nil {
			e <- fmt.Errorf("subscription error: %w", err)
			return
		}
	}()

	fmt.Println("Listening for events...")
	for {
		select {
		case borrow := <-l.Borrow.Incoming:
			handleBorrow(db, borrow)
		case deposit := <-l.Deposit.Incoming:
			handleDeposit(db, deposit)
		case flashLoan := <-l.FlashLoan.Incoming:
			handleFlashLoan(db, flashLoan)
		case liquidation := <-l.Liquidation.Incoming:
			handleLiquidation(db, liquidation)
		case rebalance := <-l.Rebalance.Incoming:
			handleRebalance(db, rebalance)
		case repay := <-l.Repay.Incoming:
			handleRepay(db, repay)
		case reserveUpdate := <-l.ReserveUpdate.Incoming:
			handleReserveUpdate(reserveUpdate)
		case collateralDisabled := <-l.CollateralDisabled.Incoming:
			handleCollateralDisabled(db, collateralDisabled)
		case collateralEnabled := <-l.CollateralEnabled.Incoming:
			handleCollateralEnabled(db, collateralEnabled)
		case swap := <-l.Swap.Incoming:
			handleSwap(db, swap)
		case withdraw := <-l.Withdraw.Incoming:
			handleWithdraw(db, withdraw)
		}
	}
}

// Returns subscription error.
func (l *Listener) listenSubError() error {
	// Unblocks if an error occurs.
	select {
	case err := <-l.Borrow.Sub.Err():
		return err
	case err := <-l.Deposit.Sub.Err():
		return err
	case err := <-l.FlashLoan.Sub.Err():
		return err
	case err := <-l.Liquidation.Sub.Err():
		return err
	case err := <-l.Rebalance.Sub.Err():
		return err
	case err := <-l.Repay.Sub.Err():
		return err
	case err := <-l.ReserveUpdate.Sub.Err():
		return err
	case err := <-l.CollateralDisabled.Sub.Err():
		return err
	case err := <-l.CollateralEnabled.Sub.Err():
		return err
	case err := <-l.Swap.Sub.Err():
		return err
	case err := <-l.Withdraw.Sub.Err():
		return err
	}

}
