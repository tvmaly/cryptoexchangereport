package quotes

type KrakenQuote struct {
	Error  []interface{} `json:"error"`
	Result struct {
		Bcheur struct {
			A []string `json:"a"`
			B []string `json:"b"`
			C []string `json:"c"`
			H []string `json:"h"`
			L []string `json:"l"`
			O string   `json:"o"`
			P []string `json:"p"`
			T []int64  `json:"t"`
			V []string `json:"v"`
		} `json:"BCHEUR"`
	} `json:"result"`
}
