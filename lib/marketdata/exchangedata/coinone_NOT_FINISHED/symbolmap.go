package coinone

import (
	"cryptoexchangereport/marketdata/assets"
	"encoding/json"
	"gopkg.in/resty.v1"
	"log"
)

func NewSymbolMap() *assets.SymbolMap {
	symbols, err := getSymbols(ApiBase + SymbolsEndpoint)

	if err != nil {
		log.Fatalf("Error while getting symbols from coinone: %v", err)
	}

	return assets.NewSymbolMap("coinone", "-", true, symbols)
}

func getSymbols(url string) (map[string][2]string, error) {
	var symbols map[string][2]string = map[string][2]string{}

	resp, err := resty.R().Get(url)

	if err != nil {
		return symbols, err
	}

	type Market struct {
		MarketName     string `json:MarketName`
		MarketCurrency string `json:MarketCurrency`
		BaseCurrency   string `json:BaseCurrency`
	}

	type Markets struct {
		Result []*Market `json:result`
	}

	var markets []*Market
	var result Markets = Markets{Result: markets}

	err = json.Unmarshal(resp.Body(), &result)

	if err != nil {
		return symbols, err
	}

	for _, market := range result.Result {
		symbols[(*market).MarketName] = [2]string{
			(*market).MarketCurrency,
			(*market).BaseCurrency,
		}
	}

	return symbols, nil
}
