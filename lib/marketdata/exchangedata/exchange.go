package exchangedata

import (
	"cryptoexchangereport/marketdata/assets"
)

type Exchange interface {
	Name() string
	OrderBook(string, string) ([]byte, error)
	Ticker(string, string) ([]byte, error)
	LastTrades(string, string) ([]byte, error)
	Limit() int
	assets.ExchangeAssetsChecker
}
