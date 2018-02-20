package bitfinex

import (
	"cryptoexchangereport/marketdata/allcoin"
	"cryptoexchangereport/marketdata/assets"
	"encoding/json"
	"gopkg.in/resty.v1"
	"log"
)

func GetSymbolMap() *assets.SymbolMap {

	symbols, err := getSymbols(SymbolsUrl)

	if err != nil {
		log.Fatalf("Error while getting symbols from bitfinex: %v", err)
	}

	return assets.NewSymbolMap("bitfinex", "", symbols)
}

func getSymbols(url string) (map[string][2]string, error) {
	var symbols map[string][2]string = map[string][2]string{}

	allCoins := GetAllCoins()

	resp, err := resty.R().Get(url)

	if err != nil {
		return symbols, err
	}

	var jsonSymbols []string

	err = json.Unmarshal(resp.Body(), &jsonSymbols)

	if err != nil {
		return symbols, err
	}

	for _, sym := range jsonSymbols {

		coins, err := allCoins.SliceSymbolOnCoins(sym)

		if err != nil {
			return symbols, err
		}

		symbols[sym] = [2]string{
			coins[0],
			coins[1],
		}
	}

	return symbols, nil
}

func GetAllCoins() allcoin.Coins {
	allCoins, err := allcoin.NewFromAPI()

	if err != nil {
		log.Fatal(err)
	}

	allCoins.Add("USD", "US Dollar")
	allCoins.Add("EUR", "Euro")

	// Workaround as Bitfinex represents some coins differently on they API:
	// QTUM - QTM, QASH - QSH, YOYOW - YYW, MANA - MNA, SNGLS - SNG
	allCoins.Add("QTM", "QTUM")
	allCoins.Add("QSH", "QASH")
	allCoins.Add("YYW", "YOYOW")
	allCoins.Add("MNA", "MANA")
	allCoins.Add("SNG", "SNGLS")

	return allCoins
}
