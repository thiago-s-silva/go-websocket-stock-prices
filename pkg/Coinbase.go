package pkg

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type Coinbase interface {
	Connect() error
	Disconnect() error
	Subscribe(symbol string) error
	Listen() ([]byte, error)
}

type coinbase struct {
	dealer websocket.Dialer
	ws     string
	conn   *websocket.Conn
}

func NewCoinbase(ws string) *coinbase {
	return &coinbase{dealer: websocket.Dialer{}, ws: ws}
}

func (c *coinbase) Connect() error {
	conn, _, err := c.dealer.Dial(c.ws, nil)
	if err != nil {
		return err
	}
	c.conn = conn

	return nil
}

func (c *coinbase) Disconnect() error {
	return c.conn.Close()
}

func (c *coinbase) Subscribe(symbol string) error {
	subscribeMessage := []byte(fmt.Sprintf(`{"action":"subscribe", "symbols":"%s"}`, symbol))

	return c.conn.WriteMessage(websocket.TextMessage, subscribeMessage)
}

func (c *coinbase) Listen() ([]byte, error) {
	_, message, err := c.conn.ReadMessage()
	if err != nil {
		return nil, err
	}

	return message, nil
}
