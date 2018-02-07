package parsing

import (
	"encoding/json"
)

type Quote interface {
	MidPoint() float64
}

type BinanceQuote struct {
	Symbol      string  `json:symbol`
	BidPrice    float64 `json:"bidPrice,string"`
	BidQuantity float64 `json:"bidQty,string"`
	AskPrice    float64 `json:"askPrice,string"`
	AskQuantity float64 `json:"askQty,string"`
}

func NewQuotes(exchange string, data []byte) ([]BinanceQuote, error) {

	if exchange == "biance" {

		var quotes []BinanceQuote

		err := json.Unmarshal(data, &quotes)

		return quotes, err

	}

	return nil, nil
}

func (q *BinanceQuote) MidPoint() float64 {
	return (q.BidPrice*q.BidQuantity + q.AskPrice*q.AskQuantity) / (q.BidQuantity + q.AskQuantity)
}
