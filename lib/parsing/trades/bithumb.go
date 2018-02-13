package trades

type BithumbTrades struct {
	Data []struct {
		ContNo          string `json:"cont_no"`
		Price           string `json:"price"`
		Total           string `json:"total"`
		TransactionDate string `json:"transaction_date"`
		Type            string `json:"type"`
		UnitsTraded     string `json:"units_traded"`
	} `json:"data"`
	Status string `json:"status"`
}
