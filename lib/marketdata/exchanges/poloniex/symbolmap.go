package poloniex

import (
	"cryptoexchangereport/marketdata/assets"
	"encoding/json"
	"gopkg.in/resty.v1"
	"log"
	"strings"
)

func NewSymbolMap() *assets.SymbolMap {
	symbols, err := getSymbols(ApiBase + SymbolsEndpoint)

	if err != nil {
		log.Fatalf("Error while getting symbols from poloniex: %v", err)
	}

	return assets.NewSymbolMap("poloniex", "_", true, symbols)
}

func getSymbols(url string) (map[string][2]string, error) {
	var symbols map[string][2]string = map[string][2]string{}

	resp, err := resty.R().Get(url)

	if err != nil {
		return symbols, err
	}

	var markets map[string]interface{}

	err = json.Unmarshal(resp.Body(), &markets)

	if err != nil {
		return symbols, err
	}

	for symbol, _ := range markets {
		coins := strings.Split(symbol, "_")

		symbols[symbol] = [2]string{
			coins[0],
			coins[1],
		}
	}

	return symbols, nil
}
