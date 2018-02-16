package quotes

type BithumbQuote struct {
	Data struct {
		AveragePrice float64 `json:"average_price,string"`
		BuyPrice     float64 `json:"buy_price,string"`
		ClosingPrice float64 `json:"closing_price,string"`
		Date         float64 `json:"date,string"`
		MaxPrice     float64 `json:"max_price,string"`
		MinPrice     float64 `json:"min_price,string"`
		OpeningPrice float64 `json:"opening_price,string"`
		SellPrice    float64 `json:"sell_price,string"`
		UnitsTraded  float64 `json:"units_traded,string"`
		Volume1day   float64 `json:"volume_1day,string"`
		Volume7day   float64 `json:"volume_7day,string"`
	} `json:"data"`
	Status string `json:"status"`
}
