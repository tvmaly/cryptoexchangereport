package book

type BithumbBook struct {
	Data struct {
		Asks []struct {
			Price    string `json:"price"`
			Quantity string `json:"quantity"`
		} `json:"asks"`
		Bids []struct {
			Price    string `json:"price"`
			Quantity string `json:"quantity"`
		} `json:"bids"`
		OrderCurrency   string `json:"order_currency"`
		PaymentCurrency string `json:"payment_currency"`
		Timestamp       string `json:"timestamp"`
	} `json:"data"`
	Status string `json:"status"`
}
