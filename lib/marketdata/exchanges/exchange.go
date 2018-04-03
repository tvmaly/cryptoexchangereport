package exchanges

import (
	"cryptoexchangereport/marketdata/assets"
)

type Exchange interface {
	Name() string
	OrderBook(string, string) (Response, error)
	Ticker(string, string) (Response, error)
	LastTrades(string, string) (Response, error)
	assets.ExchangeAssetsChecker
}
