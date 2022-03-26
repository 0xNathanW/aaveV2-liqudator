package main

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	cli, err := ethclient.Dial("wss://eth-mainnet.alchemyapi.io/v2/6XlZuWxlBrwztE3riceGmPMPLQxP44jD")
	if err != nil {
		log.Fatal(err)
	}
	_ = cli
}
