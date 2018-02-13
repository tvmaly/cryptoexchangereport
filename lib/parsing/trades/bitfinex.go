package trades

type BitFinexTrades []struct {
	Amount    string `json:"amount"`
	Exchange  string `json:"exchange"`
	Price     string `json:"price"`
	Tid       int64  `json:"tid"`
	Timestamp int64  `json:"timestamp"`
	Type      string `json:"type"`
}
