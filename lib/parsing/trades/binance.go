package trades

type BinanceTrades []struct {
	ID           int64  `json:"id"`
	IsBestMatch  bool   `json:"isBestMatch"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
	Price        string `json:"price"`
	Qty          string `json:"qty"`
	Time         int64  `json:"time"`
}
