package hitbtc

import (
	"cryptoexchangereport/marketdata/assets"
	"encoding/json"
	"gopkg.in/resty.v1"
	"log"
)

func NewSymbolMap() *assets.SymbolMap {
	symbols, err := getSymbols(ApiBase + SymbolsEndpoint)

	if err != nil {
		log.Fatalf("Error while getting symbols from hitbtc: %v", err)
	}

	return assets.NewSymbolMap("hitbtc", "", true, symbols)
}

func getSymbols(url string) (map[string][2]string, error) {
	var symbols map[string][2]string = map[string][2]string{}

	resp, err := resty.R().Get(url)

	if err != nil {
		return symbols, err
	}

	type Market struct {
		ID            string `json:id`
		BaseCurrency  string `json:baseCurrency`
		QuoteCurrency string `json:quoteCurrency`
	}

	var markets []*Market

	err = json.Unmarshal(resp.Body(), &markets)

	if err != nil {
		return symbols, err
	}

	for _, market := range markets {
		symbols[(*market).ID] = [2]string{
			(*market).BaseCurrency,
			(*market).QuoteCurrency,
		}
	}

	return symbols, nil
}
