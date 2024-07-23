package internals

import (
	"fmt"
	"github.com/thiago-s-silva/go-websocket-stock-prices/pkg"
	"log"
)

type Server interface {
	Run() error
	Stop() error
}

type server struct {
	coinbaseService pkg.Coinbase
}

func NewServer(coinbaseService pkg.Coinbase) *server {
	return &server{coinbaseService: coinbaseService}
}

func (s *server) Run() error {
	fmt.Println("Starting server...")

	// Open the websocket connection
	if err := s.coinbaseService.Connect(); err != nil {
		return err
	}
	fmt.Println("Websocket connection established")

	// Subscribe to ETH-USD channel
	if err := s.coinbaseService.Subscribe("ETH-USD"); err != nil {
		return err
	}
	fmt.Println("Subscribed to ETH-USD channel")

	// Start listening for the messages
	for {
		message, err := s.coinbaseService.Listen()
		if err != nil {
			return err
		}
		fmt.Println(string(message))
	}
}

func (s *server) Stop() {
	fmt.Println("Disconnecting from Websocket")

	// Close the Websocket connection
	err := s.coinbaseService.Disconnect()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server stopped")
}
