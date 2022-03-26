package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/0xNathanW/aaveV2-liqOSS/bot"
	"github.com/joho/godotenv"
)

func main() {

	var ( // Get variables from .env file.
		endPoint        = os.Getenv("KOVAN_RPC_URL")
		contractAddress = os.Getenv("CONTRACT_ADDRESS")
		privKey         = os.Getenv("PRIV_KEY")
		dbPath          = os.Getenv("DB_PATH")
	)

	bot, err := bot.NewBot(
		endPoint,
		dbPath,
		privKey,
		contractAddress,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Bot initialised.")
	bot.Run()
	// Sleep to allow for graceful shutdown.
	time.Sleep(time.Second * 5)
}

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
}
