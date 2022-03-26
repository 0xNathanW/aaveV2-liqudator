package bot

import (
	"log"
	"math/big"
	"os"
	"os/signal"
	"syscall"

	"github.com/0xNathanW/aaveV2-liqOSS/tests"
	"github.com/ethereum/go-ethereum/common"
)

func (b *Bot) Run() {

	go b.checkInterupt()

	//go b.listener.Listen(b.err, b.client, b.db, b.contracts.LendingPool)

	b.iterating = true
	for {
		if !b.iterating { // Breaks out of loop if interupted.
			break
		}
		iter := b.db.NewIterator(nil, nil)

	iterateDB:
		for iter.Next() {

			userAddress := common.BytesToAddress(iter.Key())
			// Retrieve data associated with the user from lendingPool.
			accountData, err := b.contracts.LendingPool.GetUserAccountData(nil, userAddress)
			if err != nil {
				log.Printf("Error getting user account data: %s\n", err)
				continue iterateDB
			}

			//log.Printf("User %s health factor: %s\n", userAddress.Hex(), dec.NewFromBigInt(accountData.HealthFactor, -18).String())

			// If the user is below threshold, attempt to liquidate.
			if accountData.HealthFactor.Cmp(big.NewInt(1e18)) < 0 {
				b.db.Put(userAddress.Bytes(), []byte("liquidatable"), nil)
				log.Println("User", userAddress.Hex(), "is below threshold, liquidation possible")
				if err := b.attemptLiq(userAddress); err != nil {
					log.Println("Error attempting liquidation:", err)
					continue iterateDB
				}
				break iterateDB
			}
		}
		iter.Release() // For testing
		return         // For testing
	}
}

// Cleans up when user interrupts or error occurs.
func (b *Bot) checkInterupt() {
	interupt := make(chan os.Signal, 1)
	defer close(interupt)
	defer signal.Stop(interupt)
	signal.Notify(interupt, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-interupt:
		log.Println("Interupted by user")
		b.shutdown()
	case e := <-b.err:
		log.Println("Error occured:", e.Error())
		b.shutdown()
	}
}

func (b *Bot) Setup() {

	accounts, err := tests.InitAccounts()
	if err != nil {
		log.Fatal(err)
	}

	if err := tests.SetupTarget(b.client, b.contracts.LendingPool, accounts[4]); err != nil {
		log.Fatal(err)
	}
}
