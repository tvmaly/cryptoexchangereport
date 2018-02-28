package binance

import (
	"cryptoexchangereport/marketdata/assets"
	"encoding/json"
	"gopkg.in/resty.v1"
	"log"
)

func NewSymbolMap() *assets.SymbolMap {
	symbols, err := getSymbols(ApiBase + SymbolsEndpoint)

	if err != nil {
		log.Fatalf("Error while getting symbols from binance: %v", err)
	}

	return assets.NewSymbolMap("binance", "", true, symbols)
}

func getSymbols(url string) (map[string][2]string, error) {
	var symbols map[string][2]string = map[string][2]string{}

	resp, err := resty.R().Get(url)

	if err != nil {
		return symbols, err
	}

	type SymbolJSON struct {
		Symbol     string `json:symbol`
		BaseAsset  string `json:baseAsse`
		QuoteAsset string `json:quoteAsset`
	}

	type SymbolsJSON struct {
		Symbols []*SymbolJSON `json:symbols`
	}

	var symbolJSON []*SymbolJSON
	var symbolsJSON = SymbolsJSON{Symbols: symbolJSON}

	err = json.Unmarshal(resp.Body(), &symbolsJSON)

	if err != nil {
		return symbols, err
	}

	for _, sym := range symbolsJSON.Symbols {
		symbols[(*sym).Symbol] = [2]string{
			(*sym).BaseAsset,
			(*sym).QuoteAsset,
		}
	}

	return symbols, nil
}
