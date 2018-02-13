package quotes

import (
	"encoding/json"
)

type BitZQuote struct {
	Code int64 `json:"code"`
	Data struct {
		Buy  string `json:"buy"`
		Date int64  `json:"date"`
		High string `json:"high"`
		Last string `json:"last"`
		Low  string `json:"low"`
		Sell string `json:"sell"`
		Vol  string `json:"vol"`
	} `json:"data"`
	Msg string `json:"msg"`
}
