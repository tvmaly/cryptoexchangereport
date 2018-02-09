package quotes

type BitTrexQuote struct {
	Message string `json:"message"`
	Result  struct {
		Ask  float64 `json:"Ask"`
		Bid  float64 `json:"Bid"`
		Last float64 `json:"Last"`
	} `json:"result"`
	Success bool `json:"success"`
}
