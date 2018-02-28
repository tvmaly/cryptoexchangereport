package gemini

import (
	"encoding/json"
	"log"

	"gopkg.in/resty.v1"

	"cryptoexchangereport/marketdata/assets"
	"github.com/agmr/allcoin"
)

func NewSymbolMap() *assets.SymbolMap {
	symbols, err := getSymbols(ApiBase + SymbolsEndpoint)

	if err != nil {
		log.Fatalf("Error while getting symbols from gemini: %v", err)
	}

	return assets.NewSymbolMap("gemini", "", false, symbols)
}

func getSymbols(url string) (map[string][2]string, error) {
	var symbols map[string][2]string = map[string][2]string{}

	allCoins := GetAllCoins()

	resp, err := resty.R().Get(url)

	if err != nil {
		return symbols, err
	}

	var markets []string

	err = json.Unmarshal(resp.Body(), &markets)

	if err != nil {
		return symbols, err
	}

	for _, market := range markets {
		coins, err := allCoins.SliceSymbolOnCoins(market)

		if err != nil {
			return symbols, err
		}

		symbols[market] = [2]string{
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

	return allCoins
}
