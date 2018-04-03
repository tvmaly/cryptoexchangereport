package assetsenumfactory

import (
	"cryptoexchangereport/marketdata/assets"
	"cryptoexchangereport/marketdata/exchanges/binance"
)

var exchangeAssetsEnumMap map[string]func() assets.ExchangeAssetsChecker = map[string]func() assets.ExchangeAssetsChecker{
	"binance": binance.NewExchangeAssetsEnum,
}

func NewExchangeAssetsEnum(exchange string) assets.ExchangeAssetsChecker {
	if method, ok := exchangeAssetsEnumMap[exchange]; ok {
		return method()
	} else {
		return nil
	}
}
