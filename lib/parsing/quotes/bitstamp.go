package quotes

import (
	"encoding/json"
)

type BitStampQuote struct {
	Ask       float64 `json:"ask,string"`
	Bid       float64 `json:"bid,string"`
	High      float64 `json:"high,string"`
	Last      float64 `json:"last,string"`
	Low       float64 `json:"low,string"`
	Open      float64 `json:"open,string"`
	Timestamp float64 `json:"timestamp,string"`
	Volume    float64 `json:"volume,string"`
	Vwap      float64 `json:"vwap,string"`
}
