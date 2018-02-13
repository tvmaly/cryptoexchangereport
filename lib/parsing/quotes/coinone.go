package quotes

type CoinoneQuote struct {
	Currency        string `json:"currency"`
	ErrorCode       string `json:"errorCode"`
	First           string `json:"first"`
	High            string `json:"high"`
	Last            string `json:"last"`
	Low             string `json:"low"`
	Result          string `json:"result"`
	Timestamp       string `json:"timestamp"`
	Volume          string `json:"volume"`
	YesterdayFirst  string `json:"yesterday_first"`
	YesterdayHigh   string `json:"yesterday_high"`
	YesterdayLast   string `json:"yesterday_last"`
	YesterdayLow    string `json:"yesterday_low"`
	YesterdayVolume string `json:"yesterday_volume"`
}
