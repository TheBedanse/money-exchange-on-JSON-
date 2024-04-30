package main

import (
	exchangeHendler "github.com/vintrinsics/money-exchange/internal/handler"
	"log"
)

func main() {
	router := exchangeHendler.ExchangeHandler()
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
