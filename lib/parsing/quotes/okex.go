package quotes

type OkexQuote struct {
	Date   string `json:"date"`
	Ticker struct {
		Buy  string `json:"buy"`
		High string `json:"high"`
		Last string `json:"last"`
		Low  string `json:"low"`
		Sell string `json:"sell"`
		Vol  string `json:"vol"`
	} `json:"ticker"`
}
