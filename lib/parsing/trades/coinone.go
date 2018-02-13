package trades

type CoinoneTrades struct {
	CompleteOrders []struct {
		Price     string `json:"price"`
		Qty       string `json:"qty"`
		Timestamp string `json:"timestamp"`
	} `json:"completeOrders"`
	Currency  string `json:"currency"`
	ErrorCode string `json:"errorCode"`
	Result    string `json:"result"`
	Timestamp string `json:"timestamp"`
}
