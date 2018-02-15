package quotes

import (
	"errors"
	"github.com/tvmaly/cryptoexchangereport/lib/constants"
	"github.com/tvmaly/cryptoexchangereport/lib/parsing/utility"
)

// GenericQuote this represents the basic set of fields that
// we wish to assume a quote has
// the one exception is the quantity fields
// these are not always present and this will affect calculations
// in respect to them being weighted by quantity values
type GenericQuote struct {
	Exchange    string  `json:exchange`
	Symbol      string  `json:symbol`
	Timestamp   string  `json:timestamp`
	BidPrice    float64 `json:"bidprice,string"`
	BidQuantity float64 `json:"bidquantity,string"`
	AskPrice    float64 `json:"askprice,string"`
	AskQuantity float64 `json:"askquantity,string"`
}

// MidPoint calculate the mid point price of a quote
// if quantity is available, it will be weighted
// otherwise it will be a simple midpoint
func (q *GenericQuote) MidPoint() float64 {

	if q.Weighted() {
		return (q.BidPrice*q.BidQuantity + q.AskPrice*q.AskQuantity) / (q.BidQuantity + q.AskQuantity)
	} else {
		return (q.BidPrice + q.AskPrice) / 2
	}
}

// Weighted return true if quantity is set otherwise false
func (q *GenericQuote) Weighted() bool {
	return q.BidQuantity != 0 && q.AskQuantity != 0
}

func (t *BinanceQuote) GenericQuote() *GenericQuote {

	return &GenericQuote{
		Exchange:    constants.BINANCE,
		TimeStamp:   t.Timestamp,
		BidPrice:    t.BidPrice,
		BidQuantity: t.BidQuantity,
		AskPrice:    t.AskPrice,
		AskQuantity: t.AskQuantity,
	}
}

func (t *BitZQuote) GenericQuote() *GenericQuote {

	return &GenericQuote{
		Exchange:  constants.BITZ,
		TimeStamp: t.Data.Date,
		BidPrice:  t.Data.Buy,
		AskPrice:  t.Data.Sell,
	}
}

func (t *BitFinexQuote) GenericQuote() *GenericQuote {

	return &GenericQuote{
		Exchange:  constants.BITFINEX,
		TimeStamp: t.Timestamp,
		BidPrice:  t.Bid,
		AskPrice:  t.Ask,
	}
}

func (t *BithumbQuotes) GenericQuote() *GenericQuote {

	return &GenericQuote{
		Exchange:  constants.BITFINEX,
		TimeStamp: t.Data.Date,
		BidPrice:  t.Data.BuyPrice,
		AskPrice:  t.Data.SellPrice,
	}
}

func (t *GDaxQuote) GenericQuote() *GenericQuote {

	return &GenericQuote{
		Exchange:  constants.GDAX,
		TimeStamp: t.Time,
		BidPrice:  t.Bid,
		AskPrice:  t.Ask,
	}
}
