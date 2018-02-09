package quotes

type HitBTCQuote struct {
	Ask         string `json:"ask"`
	Bid         string `json:"bid"`
	High        string `json:"high"`
	Last        string `json:"last"`
	Low         string `json:"low"`
	Open        string `json:"open"`
	Symbol      string `json:"symbol"`
	Timestamp   string `json:"timestamp"`
	Volume      string `json:"volume"`
	VolumeQuote string `json:"volumeQuote"`
}
