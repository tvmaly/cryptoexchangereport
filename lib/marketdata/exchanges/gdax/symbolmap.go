package gdax

import (
	"cryptoexchangereport/marketdata/assets"
	"encoding/json"
	"gopkg.in/resty.v1"
	"log"
)

func NewSymbolMap() *assets.SymbolMap {
	symbols, err := getSymbols(ApiBase + SymbolsEndpoint)

	if err != nil {
		log.Fatalf("Error while getting symbols from gdax: %v", err)
	}

	return assets.NewSymbolMap("gdax", "-", true, symbols)
}

func getSymbols(url string) (map[string][2]string, error) {
	var symbols map[string][2]string = map[string][2]string{}

	resp, err := resty.R().Get(url)

	if err != nil {
		return symbols, err
	}

	type Market struct {
		ID             string `json:id`
		Base_Currency  string `json:base_currency`
		Quote_Currency string `json:quote_currency`
	}

	var markets []*Market

	err = json.Unmarshal(resp.Body(), &markets)

	if err != nil {
		return symbols, err
	}

	for _, market := range markets {
		symbols[(*market).ID] = [2]string{
			(*market).Base_Currency,
			(*market).Quote_Currency,
		}
	}

	return symbols, nil
}
