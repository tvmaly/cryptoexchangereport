package book

type BinanceBook struct {
	Asks         [][]interface{} `json:"asks"`
	Bids         [][]interface{} `json:"bids"`
	LastUpdateID int64           `json:"lastUpdateId"`
}
