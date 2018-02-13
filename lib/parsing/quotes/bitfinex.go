package quotes

import (
	"encoding/json"
)

type BitFinexQuote struct {
	Ask       float64 `json:"ask,string"`
	Bid       float64 `json:"bid,string"`
	High      float64 `json:"high,string"`
	LastPrice float64 `json:"last_price,string"`
	Low       float64 `json:"low,string"`
	Mid       float64 `json:"mid,string"`
	Timestamp float64 `json:"timestamp,string"`
	Volume    float64 `json:"volume,string"`
}
