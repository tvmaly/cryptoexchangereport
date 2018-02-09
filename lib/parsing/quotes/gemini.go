package quotes

type GDaxBTCUSDQuote struct {
	Ask    string `json:"ask"`
	Bid    string `json:"bid"`
	Last   string `json:"last"`
	Volume struct {
		BTC       string `json:"BTC"`
		USD       string `json:"USD"`
		Timestamp int64  `json:"timestamp"`
	} `json:"volume"`
}

type GDaxETHBTCQuote struct {
	Ask    string `json:"ask"`
	Bid    string `json:"bid"`
	Last   string `json:"last"`
	Volume struct {
		ETH       string `json:"ETH"`
		BTC       string `json:"BTC"`
		Timestamp int64  `json:"timestamp"`
	} `json:"volume"`
}

type GDaxETHUSDQuote struct {
	Ask    string `json:"ask"`
	Bid    string `json:"bid"`
	Last   string `json:"last"`
	Volume struct {
		ETH       string `json:"ETH"`
		USD       string `json:"USD"`
		Timestamp int64  `json:"timestamp"`
	} `json:"volume"`
}
