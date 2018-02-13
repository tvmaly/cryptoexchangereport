package quotes

import (
	"encoding/json"
)

type BinanceQuote struct {
	Symbol      string  `json:symbol`
	BidPrice    float64 `json:"bidPrice,string"`
	BidQuantity float64 `json:"bidQty,string"`
	AskPrice    float64 `json:"askPrice,string"`
	AskQuantity float64 `json:"askQty,string"`
	Timestamp   float64 `json:"timestamp,string"`
}

func NewBinanceQuotes(data []byte, timestamp float64) ([]*BinanceQuote, error) {
	var quotes []*BinanceQuote
	err := json.Unmarshal(data, &quotes)

	if err != nil {
		return quotes, err
	}

	for _, v := range quotes {
		v.Timestamp = timestamp
	}

	return quotes, err
}
