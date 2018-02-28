package liqui

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
		log.Fatalf("Error while getting symbols from liqui: %v", err)
	}

	return assets.NewSymbolMap("liqui", "_", false, symbols)
}

func getSymbols(url string) (map[string][2]string, error) {
	var symbols map[string][2]string = map[string][2]string{}

	resp, err := resty.R().Get(url)

	if err != nil {
		return symbols, err
	}

	type Market struct {
		Pairs map[string]interface{} `json:pairs`
	}

	var market Market

	err = json.Unmarshal(resp.Body(), &market)

	if err != nil {
		return symbols, err
	}

	for symbol, _ := range market.Pairs {

		coins := strings.Split(symbol, "_")

		symbols[symbol] = [2]string{
			coins[0],
			coins[1],
		}
	}

	return symbols, nil
}
