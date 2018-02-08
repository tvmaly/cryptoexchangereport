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
}

func NewBinanceQuotes(data []byte) ([]*BinanceQuote, error) {
	var quotes []*BinanceQuote
	err := json.Unmarshal(data, &quotes)
	return quotes, err
}

func (q *BinanceQuote) MidPoint() float64 {
	return (q.BidPrice*q.BidQuantity + q.AskPrice*q.AskQuantity) / (q.BidQuantity + q.AskQuantity)
}

func (q *BinanceQuote) IsWeighted() bool {
	return true
}
