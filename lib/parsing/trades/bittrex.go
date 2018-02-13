package trades

type BitTrexTrades struct {
	Message string `json:"message"`
	Result  []struct {
		FillType  string  `json:"FillType"`
		ID        int64   `json:"Id"`
		OrderType string  `json:"OrderType"`
		Price     float64 `json:"Price"`
		Quantity  float64 `json:"Quantity"`
		TimeStamp string  `json:"TimeStamp"`
		Total     float64 `json:"Total"`
	} `json:"result"`
	Success bool `json:"success"`
}
