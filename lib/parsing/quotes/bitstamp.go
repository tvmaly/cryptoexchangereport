package quotes

import (
	"encoding/json"
)

type BitStampQuote struct {
	Symbol    string  `json:symbol`
	HighPrice float64 `json:"high,string"`
	LastPrice float64 `json:"last,string"`
	Epoch     float64 `json:"timestamp,string"`
	BidPrice  float64 `json:"bid,string"`
	VWAP      float64 `json:"vwap,string"`
	Volume    float64 `json:"volume,string"`
	LowPrice  float64 `json:"low,string"`
	AskPrice  float64 `json:"ask,string"`
	OpenPrice float64 `json:"open,string"`
}

func NewBitStampQuote(symbol string, data []byte) (*BitStampQuote, error) {
	var quote BitStampQuote
	err := json.Unmarshal(data, &quote)
	return &quote, err
}

func (q *BitStampQuote) MidPoint() float64 {
	return (q.BidPrice + q.AskPrice) / 2.0
}

func (q *BitStampQuote) IsWeighted() bool {
	return false
}
