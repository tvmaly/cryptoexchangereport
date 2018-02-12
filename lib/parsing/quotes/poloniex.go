package quotes

import (
	"encoding/json"
	"strings"
)

type PoloniexQuote struct {
	ID            int64   `json:"id"`
	Symbol        string  `json:symbol`
	LastPrice     float64 `json:"last,string"`
	BaseVolume    float64 `json:"baseVolume,string"`
	High24hr      float64 `json:"high24hr,string"`
	BidPrice      float64 `json:"highestBid,string"`
	IsFrozen      float64 `json:"isFrozen,string"`
	Low24hr       float64 `json:"low24hr,string"`
	AskPrice      float64 `json:"lowestAsk,string"`
	PercentChange float64 `json:"percentChange,string"`
	QuoteVolume   float64 `json:"quoteVolume,string"`
}

func NewPoloniexQuotes(data []byte) ([]*PoloniexQuote, error) {

	var quotemap map[string]*PoloniexQuote

	var quotes []*PoloniexQuote

	err := json.Unmarshal(data, &quotemap)

	if err != nil {
		return quotes, err
	}

	for symbol, quote := range quotemap {
		quote.Symbol = CorrectPoloniexSymbol(symbol)
		quotes = append(quotes, quote)
	}

	return quotes, err
}

func (q *PoloniexQuote) MidPoint() float64 {
	return (q.BidPrice + q.AskPrice) / 2
}

func (q *PoloniexQuote) IsWeighted() bool {
	return false
}

func CorrectPoloniexSymbol(poloniex_symbol string) string {

	parts := strings.Split(poloniex_symbol, "_")

	return parts[1] + parts[0]

}
