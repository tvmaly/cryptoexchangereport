package trades

type BitZTrades struct {
	Code int64 `json:"code"`
	Data struct {
		D []struct {
			N string `json:"n"`
			P string `json:"p"`
			S string `json:"s"`
			T string `json:"t"`
		} `json:"d"`
		Max string `json:"max"`
		Min string `json:"min"`
		Sum string `json:"sum"`
	} `json:"data"`
	Msg string `json:"msg"`
}
