package Coinbase

import "encoding/json"

type CryptoMessage struct {
	TickerCode            string  `json:"s,omitempty"`
	LastPrice             string  `json:"p,omitempty"`
	Quantity              string  `json:"q,omitempty"`
	DailyChangePercentage string  `json:"dc,omitempty"`
	DailyDiffPercentage   string  `json:"dd,omitempty"`
	Timestamp             float32 `json:"t,omitempty"`
}

func NewCryptoMessageFromJSON(data []byte) (*CryptoMessage, error) {
	var cryptoMessage CryptoMessage

	err := json.Unmarshal(data, &cryptoMessage)
	if err != nil {
		return nil, err
	}

	return &cryptoMessage, nil
}
