package internals

import (
	"fmt"
	"github.com/thiago-s-silva/go-websocket-stock-prices/pkg/Coinbase"
)

func processCrypto(id int, c <-chan []byte) {
	// Start listening for message on crypto channel
	for msg := range c {
		// Log the received message
		fmt.Printf("worker %d received: %s\n", id, msg)

		// Unmarshal the JSON to the Crypto Message entity
		_, err := Coinbase.NewCryptoMessageFromJSON(msg)
		if err != nil {
			fmt.Printf("worker %d error: %s\n", id, err)
			continue
		}
	}

	// Call the finish message at the end
	fmt.Printf("worker %d finished\n", id)
}
