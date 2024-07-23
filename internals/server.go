package internals

import (
	"fmt"
	"github.com/thiago-s-silva/go-websocket-stock-prices/pkg/Coinbase"
	"log"
	"os"
	"strconv"
)

type Server interface {
	Run() error
	Stop() error
}

type server struct {
	coinbaseService Coinbase.Coinbase
	cryptoChannel   chan []byte
	numberOfWorkers int
}

func NewServer(coinbaseService Coinbase.Coinbase) *server {
	numberOfWorkers, err := strconv.Atoi(os.Getenv("NUMBER_OF_WORKERS"))
	if err != nil {
		numberOfWorkers = 1
	}

	return &server{
		coinbaseService: coinbaseService,
		cryptoChannel:   make(chan []byte, numberOfWorkers),
		numberOfWorkers: numberOfWorkers,
	}
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

	// Init the workers to start processing the received crypto messages
	s.initWorkers()

	// Start listening for the messages
	for {
		message, err := s.coinbaseService.Listen()
		if err != nil {
			return err
		}

		// Send the received message to the crypto channel
		s.cryptoChannel <- message
	}
}

func (s *server) Stop() {
	fmt.Println("Disconnecting from Websocket")

	// Close the Websocket connection
	err := s.coinbaseService.Disconnect()
	if err != nil {
		log.Fatal(err)
	}

	// Close the crypto channel
	close(s.cryptoChannel)

	fmt.Println("Server stopped")
}

func (s *server) initWorkers() {
	for i := 0; i < s.numberOfWorkers; i++ {
		go processCrypto(i, s.cryptoChannel)
	}

	fmt.Println("Workers initialized")
}
