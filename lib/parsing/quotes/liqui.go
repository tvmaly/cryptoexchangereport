package quotes

type LiquiQuote struct {
	EthBtc struct {
		Avg     float64 `json:"avg"`
		Buy     float64 `json:"buy"`
		High    float64 `json:"high"`
		Last    float64 `json:"last"`
		Low     float64 `json:"low"`
		Sell    float64 `json:"sell"`
		Updated int64   `json:"updated"`
		Vol     float64 `json:"vol"`
		VolCur  float64 `json:"vol_cur"`
	} `json:"eth_btc"`
}
