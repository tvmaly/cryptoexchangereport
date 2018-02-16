package book

type CoinoneBook struct {
	Ask []struct {
		Price float64 `json:"price,string"`
		Qty   float64 `json:"qty,string"`
	} `json:"ask"`
	Bid []struct {
		Price float64 `json:"price,string"`
		Qty   float64 `json:"qty,string"`
	} `json:"bid"`
	Currency  string  `json:"currency"`
	ErrorCode string  `json:"errorCode"`
	Result    string  `json:"result"`
	Timestamp float64 `json:"timestamp"`
}
