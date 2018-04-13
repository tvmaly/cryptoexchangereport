package exchanges

import (
	"cryptoexchangereport/marketdata/exchangemarkets"
)

type Exchange interface {
	Name() string
	OrderBook(string, string) (Response, error)
	Ticker(string, string) (Response, error)
	LastTrades(string, string) (Response, error)
	exchangemarkets.ExchangeMarkets
	GetExchangeMarkets() exchangemarkets.ExchangeMarkets
}
