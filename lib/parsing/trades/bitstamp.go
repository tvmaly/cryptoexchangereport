package trades

type BitStampTrades []struct {
	Amount string `json:"amount"`
	Date   string `json:"date"`
	Price  string `json:"price"`
	Tid    string `json:"tid"`
	Type   string `json:"type"`
}
