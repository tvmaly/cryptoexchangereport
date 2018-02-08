package quotes

import (
	"encoding/json"
)

type BitZTicker struct {
	Code    int       `json:code`
	Message string    `json:msg`
	Quote   BitZQuote `json:data`
}

type BitZQuote struct {
	Symbol    string  `json:symbol`
	Date      int64   `json:"date,string"`
	LastPrice float64 `json:"last,string"`
	BidPrice  float64 `json:"buy,string"`
	AskPrice  float64 `json:"sell,string"`
	High      float64 `json:"high,string"`
	Low       float64 `json:"low,string"`
	Volume    float64 `json:"vol,string"`
}

func NewBitZQuote(symbol string, data []byte) (*BitZQuote, error) {
	var ticker BitZTicker

	err := json.Unmarshal(data, &ticker)

	if err == nil {
		ticker.Quote.Symbol = symbol
	}

	return &ticker.Quote, err
}

func (q *BitZQuote) MidPoint() float64 {
	return (q.BidPrice + q.AskPrice) / 2.0
}

func (q *BitZQuote) IsWeighted() bool {
	return false
}
