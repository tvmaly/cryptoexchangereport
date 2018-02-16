package quotes

import (
	"errors"
	"github.com/tvmaly/cryptoexchangereport/lib/constants"
	"github.com/tvmaly/cryptoexchangereport/lib/parsing/utility"
)

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

func (t *BithumbQuote) GenericQuote() *GenericQuote {

	return &GenericQuote{
		Exchange:  constants.BITFINEX,
		TimeStamp: t.Data.Date,
		BidPrice:  t.Data.BuyPrice,
		AskPrice:  t.Data.SellPrice,
	}
}

func (t *BitStampQuote) GenericQuote() *GenericQuote {

	return &GenericQuote{
		Exchange:  constants.BITSTAMP,
		TimeStamp: t.Timestamp,
		BidPrice:  t.Bid,
		AskPrice:  t.Ask,
	}
}

func (t *BitTrexQuote) GenericQuote() *GenericQuote {

	return &GenericQuote{
		Exchange: constants.BITTREX,
		BidPrice: t.Result.Bid,
		AskPrice: t.Result.Ask,
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

func (t *GeminiQuote) GenericQuote() *GenericQuote {

	return &GenericQuote{
		Exchange:  constants.GEMINI,
		TimeStamp: t.Timestamp(),
		BidPrice:  t.Bid,
		AskPrice:  t.Ask,
	}
}
