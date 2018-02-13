package quotes

import (
	"errors"
	"github.com/tvmaly/cryptoexchangereport/lib/parsing/exchange"
	"github.com/tvmaly/cryptoexchangereport/lib/parsing/utility"
)

type GenericQuote struct {
	Exchange    string  `json:exchange`
	Symbol      string  `json:symbol`
	Timestamp   string  `json:timestamp`
	BidPrice    float64 `json:"bidprice,string"`
	BidQuantity float64 `json:"bidquantity,string"`
	AskPrice    float64 `json:"askprice,string"`
	AskQuantity float64 `json:"askquantity,string"`
}

func (q *GenericQuote) MidPoint() float64 {

	if q.Weighted() {
		return (q.BidPrice*q.BidQuantity + q.AskPrice*q.AskQuantity) / (q.BidQuantity + q.AskQuantity)
	} else {
		return (q.BidPrice + q.AskPrice) / 2
	}
}

func (q *GenericQuote) Weighted() bool {
	return q.BidQuantity != 0 && q.AskQuantity != 0
}

func NewGenericQuote(exchange string, symbol string, rawquote interface{}) (*GenericQuote, error) {

	if exchange == "gdax" {
		return NewGenericQuoteFromGDax(exchange, symbol, rawquote)
	}

}

func (t *BinanceQuote) GenericQuote() *GenericQuote {

	return &GenericQuote{
		Exchange:    exchange.BINANCE,
		TimeStamp:   t.Timestamp,
		BidPrice:    t.BidPrice,
		BidQuantity: t.BidQuantity,
		AskPrice:    t.AskPrice,
		AskQuantity: t.AskQuantity,
	}
}

func (t *BitFinexQuote) GenericQuote() *GenericQuote {

	return &GenericQuote{
		Exchange:  exchange.BITFINEX,
		TimeStamp: t.Timestamp,
		BidPrice:  t.Bid,
		AskPrice:  t.Ask,
	}
}

func (t *BitZQuote) GenericQuote() *GenericQuote {

	return &GenericQuote{
		Exchange:  exchange.BITZ,
		TimeStamp: t.Data.Date,
		BidPrice:  t.Data.Buy,
		AskPrice:  t.Data.Sell,
	}
}

func (t *GDaxQuote) GenericQuote() *GenericQuote {

	return &GenericQuote{
		Exchange:  exchange.GDAX,
		TimeStamp: t.Time,
		BidPrice:  t.Bid,
		AskPrice:  t.Ask,
	}
}
