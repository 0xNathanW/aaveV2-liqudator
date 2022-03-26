package main

import (
	"fmt"
	"log"

	"liquidator/bot"

	"github.com/joho/godotenv"
)

func main() {

	liquidator, err := bot.NewBot()
	if err != nil {
		log.Println("Error initialising bot!")
		log.Fatal(err)
	}
	fmt.Println("Bot initialised")

	if err := liquidator.Run(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	// loads values from .env into the system.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
