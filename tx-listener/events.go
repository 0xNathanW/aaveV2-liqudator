package txlistener

import (
	"log"

	"github.com/0xNathanW/aaveV2-liqOSS/contracts"
	ev "github.com/ethereum/go-ethereum/event"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type borrow struct {
	Sub      ev.Subscription
	Incoming chan *contracts.LendingPoolBorrow
}

func handleBorrow(db *leveldb.DB, borrow *contracts.LendingPoolBorrow) {
	log.Printf("%s borrowed %s\n", borrow.User.Hex(), borrow.Amount.String())
	db.Put(borrow.User[:], borrow.User[:], &opt.WriteOptions{Sync: false})
}

type deposit struct {
	Sub      ev.Subscription
	Incoming chan *contracts.LendingPoolDeposit
}

func handleDeposit(db *leveldb.DB, deposit *contracts.LendingPoolDeposit) {
	log.Printf("%s deposited %s\n", deposit.User.Hex(), deposit.Amount.String())
	db.Put(deposit.User[:], deposit.User[:], &opt.WriteOptions{Sync: false})
}

type flashLoan struct {
	Sub      ev.Subscription
	Incoming chan *contracts.LendingPoolFlashLoan
}

func handleFlashLoan(db *leveldb.DB, flashLoan *contracts.LendingPoolFlashLoan) {
	log.Printf("%s flash loaned %s\n", flashLoan.Initiator.Hex(), flashLoan.Amount.String())
}

type liquidation struct {
	Sub      ev.Subscription
	Incoming chan *contracts.LendingPoolLiquidationCall
}

func handleLiquidation(db *leveldb.DB, liquidation *contracts.LendingPoolLiquidationCall) {
	log.Printf(
		"%s liquidated %s, paid %s debt, recieved %s collateral\n",
		liquidation.Liquidator.Hex(),
		liquidation.User.String(),
		liquidation.DebtToCover.String(),
		liquidation.LiquidatedCollateralAmount.String())
}

type rebalance struct {
	Sub      ev.Subscription
	Incoming chan *contracts.LendingPoolRebalanceStableBorrowRate
}

func handleRebalance(db *leveldb.DB, rebalance *contracts.LendingPoolRebalanceStableBorrowRate) {
	log.Printf("%s rebalanced\n", rebalance.User.Hex())
	db.Put(rebalance.User[:], rebalance.User[:], &opt.WriteOptions{Sync: false})
}

type repay struct {
	Sub      ev.Subscription
	Incoming chan *contracts.LendingPoolRepay
}

func handleRepay(db *leveldb.DB, repay *contracts.LendingPoolRepay) {
	log.Printf("%s repaid %s\n", repay.Repayer.Hex(), repay.Amount.String())
	db.Put(repay.User[:], repay.User[:], &opt.WriteOptions{Sync: false})
}

type reserveUpdate struct {
	Sub      ev.Subscription
	Incoming chan *contracts.LendingPoolReserveDataUpdated
}

func handleReserveUpdate(reserveUpdate *contracts.LendingPoolReserveDataUpdated) {
	log.Printf("%s reserve updated\n", reserveUpdate.Reserve.Hex())
}

type collateralDisabled struct {
	Sub      ev.Subscription
	Incoming chan *contracts.LendingPoolReserveUsedAsCollateralDisabled
}

func handleCollateralDisabled(db *leveldb.DB, collateralDisabled *contracts.LendingPoolReserveUsedAsCollateralDisabled) {
	log.Printf("%s disabled %s as collateral\n",
		collateralDisabled.User.Hex(),
		collateralDisabled.Reserve.Hex())
	db.Put(collateralDisabled.User[:], collateralDisabled.User[:], &opt.WriteOptions{Sync: false})
}

type collateralEnabled struct {
	Sub      ev.Subscription
	Incoming chan *contracts.LendingPoolReserveUsedAsCollateralEnabled
}

func handleCollateralEnabled(db *leveldb.DB, collateralEnabled *contracts.LendingPoolReserveUsedAsCollateralEnabled) {
	log.Printf("%s enabled %s as collateral\n",
		collateralEnabled.User.Hex(),
		collateralEnabled.Reserve.Hex())
	db.Put(collateralEnabled.User[:], collateralEnabled.User[:], &opt.WriteOptions{Sync: false})
}

type swap struct {
	Sub      ev.Subscription
	Incoming chan *contracts.LendingPoolSwap
}

func handleSwap(db *leveldb.DB, swap *contracts.LendingPoolSwap) {
	log.Printf("%s swapped rate to %s\n", swap.User.Hex(), swap.RateMode.String())
	db.Put(swap.User[:], swap.User[:], &opt.WriteOptions{Sync: false})
}

type withdraw struct {
	Sub      ev.Subscription
	Incoming chan *contracts.LendingPoolWithdraw
}

func handleWithdraw(db *leveldb.DB, withdraw *contracts.LendingPoolWithdraw) {
	log.Printf("%s withdrew %s\n", withdraw.User.Hex(), withdraw.Amount.String())
	db.Put(withdraw.User[:], withdraw.User[:], &opt.WriteOptions{Sync: false})
}
