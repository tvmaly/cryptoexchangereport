package quotes

type GDaxQuote struct {
	Ask     float64 `json:"ask,string"`
	Bid     float64 `json:"bid,string"`
	Price   float64 `json:"price,string"`
	Size    float64 `json:"size,string"`
	Time    string  `json:"time"`
	TradeID int64   `json:"trade_id,string"`
	Volume  float64 `json:"volume,string"`
}
