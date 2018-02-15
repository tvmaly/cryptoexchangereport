package quotes

type BithumbQuotes struct {
	Data struct {
		AveragePrice string `json:"average_price"`
		BuyPrice     string `json:"buy_price"`
		ClosingPrice string `json:"closing_price"`
		Date         string `json:"date"`
		MaxPrice     string `json:"max_price"`
		MinPrice     string `json:"min_price"`
		OpeningPrice string `json:"opening_price"`
		SellPrice    string `json:"sell_price"`
		UnitsTraded  string `json:"units_traded"`
		Volume1day   string `json:"volume_1day"`
		Volume7day   string `json:"volume_7day"`
	} `json:"data"`
	Status string `json:"status"`
}
