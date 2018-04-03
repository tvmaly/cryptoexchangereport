package cointracker

import (
	"cryptoexchangereport/marketdata/assets"
	"cryptoexchangereport/marketdata/exchanges/assetsenumfactory"
	"fmt"
)

type CoinTracker interface {
	CheckForAddedCoins() *AddedCoinsStorage
	AddExchangeAssetsEnum(string, assets.ExchangeAssetsChecker) error
	RemoveExchangeAssetsEnum(string)
}

type ConcreteCoinTracker struct {
	assetsCheckers map[string]assets.ExchangeAssetsChecker
}

func New() *ConcreteCoinTracker {
	return &ConcreteCoinTracker{}
}

func (cct *ConcreteCoinTracker) CheckForAddedCoins() *AddedCoinsStorage {
	var addedCoinsStorage *AddedCoinsStorage = NewAddedCoinsStorage()

	for exchange, currentExchangeAssetLookup := range cct.assetsCheckers {
		newExchangeAssetLookup := assetsenumfactory.NewExchangeAssetsEnum(exchange)
		newCoins := Compare(newExchangeAssetLookup, currentExchangeAssetLookup)
		if newCoins != nil {
			addedCoinsStorage.Add(newCoins)
		}
	}

	return addedCoinsStorage
}

func (cct *ConcreteCoinTracker) AddExchangeAssetsEnum(exchange string) error {
	if _, ok := cct.assetsCheckers[exchange]; ok {
		err := fmt.Errorf("We already tracking %s")
		return err
	} else {
		cct.assetsCheckers[exchange] = assetsenumfactory.NewExchangeAssetsEnum(exchange)
	}
}

func (cct *ConcreteCoinTracker) RemoveExchangeAssetsEnum(exchange string) {
	delete(cct.assetsCheckers, exchange)
}

func Compare() {
	// placeholder
}
