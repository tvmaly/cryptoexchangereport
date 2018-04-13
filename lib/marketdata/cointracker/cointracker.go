package cointracker

import (
	"cryptoexchangereport/marketdata/exchangemarkets"
	"cryptoexchangereport/marketdata/exchanges/assetsenumfactory"
	"fmt"
)

type CoinTracker interface {
	CheckForAddedCoins() *AddedCoinsStorage
	AddExchangeAssetsEnum(string, exchangmarkets.ExchangeAssetsChecker) error
	RemoveExchangeAssetsEnum(string)
}

type ConcreteCoinTracker struct {
	marketsRepository map[string]exchangemarkets.ExchangeAssetsChecker
}

func New() *ConcreteCoinTracker {
	return &ConcreteCoinTracker{}
}

func (cct *ConcreteCoinTracker) CheckForAddedCoins() *AddedCoinsStorage {
	var addedCoinsStorage *AddedCoinsStorage = NewAddedCoinsStorage()

	for exchange, currentExchangeAssetLookup := range cct.marketsRepository {
		newExchangeAssetLookup := assetsenumfactory.NewExchangeAssetsEnum(exchange)
		newCoins := Compare(newExchangeAssetLookup, currentExchangeAssetLookup)
		if newCoins != nil {
			addedCoinsStorage.Add(newCoins)
		}
	}

	return addedCoinsStorage
}

func (cct *ConcreteCoinTracker) AddExchangeAssetsEnum(exchange string) error {
	if _, ok := cct.marketsRepository[exchange]; ok {
		err := fmt.Errorf("We already tracking %s")
		return err
	} else {
		cct.marketsRepository[exchange] = assetsenumfactory.NewExchangeAssetsEnum(exchange)
	}
}

func (cct *ConcreteCoinTracker) RemoveExchangeAssetsEnum(exchange string) {
	delete(cct.marketsRepository, exchange)
}

func Compare() {
	// placeholder
}
