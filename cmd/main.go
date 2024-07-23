package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/thiago-s-silva/go-websocket-stock-prices/internals"
	"github.com/thiago-s-silva/go-websocket-stock-prices/pkg/Coinbase"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Define the crypto websocket endpoint
	cryptoEndpoint := os.Getenv("CRYPTO_WS_ENDPOINT") + os.Getenv("API_KEY")
	if cryptoEndpoint == "" {
		log.Fatal("CRYPTO_WS_ENDPOINT environment variable is not set")
	}

	// Init a new coinbase service instance
	coinbaseService := Coinbase.NewCoinbase(cryptoEndpoint)

	// Init the server instance
	server := internals.NewServer(coinbaseService)

	// Run the server in a Go Routine to be able to handle graceful shutdowns
	go func() {
		if err := server.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	// Declare a buffered shutdown channel to receive OS signals
	shutdown := make(chan os.Signal, 1)

	// Notify shutdowns and interruptions to shut down a channel
	signal.Notify(shutdown, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	// Block the thread until getting a shutdown signal
	<-shutdown

	fmt.Println("Stopping the server...")

	// Stop the server
	server.Stop()
}
