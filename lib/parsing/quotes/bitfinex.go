package quotes

import (
	"encoding/json"
)

type BitFinexQuote struct {
	Symbol    string  `json:symbol`
	MidPrice  float64 `json:"mid,string"`
	BidPrice  float64 `json:"bid,string"`
	AskPrice  float64 `json:"ask,string"`
	LastPrice float64 `json:"last_price,string"`
	LowPrice  float64 `json:"low,string"`
	Volume    float64 `json:"volume,string"`
	Epoch     float64 `json:"timestamp,string"`
}

func NewBitFinexQuote(symbol string, data []byte) (*BitFinexQuote, error) {
	var quote BitFinexQuote
	err := json.Unmarshal(data, &quote)
	return &quote, err
}

func (q *BitFinexQuote) MidPoint() float64 {
	return (q.BidPrice + q.AskPrice) / 2.0
}

func (q *BitFinexQuote) IsWeighted() bool {
	return false
}
