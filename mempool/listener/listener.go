package listener

import (
	"context"
	"liquidator/contracts"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/syndtr/goleveldb/leveldb"
)

type EventListener struct {
	db     *leveldb.DB
	aave   *aaveV2Events
	subs   []event.Subscription
	errLog *log.Logger
}

func NewEventListener(
	db *leveldb.DB,
	cli *ethclient.Client,
	contracts *contracts.Contracts,
) (*EventListener, error) {

	var subs []event.Subscription

	aaveV2Events, aaveV2Subs, err := initAaveSubs(cli, contracts.AaveV2.LendingPool)
	if err != nil {
		return nil, err
	}

	subs = append(subs, aaveV2Subs...)

	return &EventListener{
		db:   db,
		aave: aaveV2Events,
	}, nil
}

func (l *EventListener) Run(ctx context.Context, contracts *contracts.Contracts) {

	for _, sub := range l.subs {
		go subscriptionHandler(ctx, sub, l.errLog)
	}

}

func subscriptionHandler(ctx context.Context, sub event.Subscription, log *log.Logger) {

	select {
	case err := <-sub.Err():
		log.Println("Error in subscription: ", err)
		return
	case <-ctx.Done():
		sub.Unsubscribe()
		return
	}

}
