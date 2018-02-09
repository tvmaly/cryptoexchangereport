package quotes

type HuobiQuote struct {
	Ch     string `json:"ch"`
	Status string `json:"status"`
	Tick   struct {
		Amount  float64   `json:"amount"`
		Ask     []float64 `json:"ask"`
		Bid     []float64 `json:"bid"`
		Close   int64     `json:"close"`
		Count   int64     `json:"count"`
		High    int64     `json:"high"`
		ID      int64     `json:"id"`
		Low     float64   `json:"low"`
		Open    float64   `json:"open"`
		Version int64     `json:"version"`
		Vol     float64   `json:"vol"`
	} `json:"tick"`
	Ts int64 `json:"ts"`
}
