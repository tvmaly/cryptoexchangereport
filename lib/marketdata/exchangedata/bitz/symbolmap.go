package bitz

import (
	"encoding/json"
	"log"
	"strings"

	"gopkg.in/resty.v1"

	"cryptoexchangereport/marketdata/assets"
)

func NewSymbolMap() *assets.SymbolMap {
	symbols, err := getSymbols(ApiBase + SymbolsEndpoint)

	if err != nil {
		log.Fatalf("Error while getting symbols from bitz: %v", err)
	}

	return assets.NewSymbolMap("bitz", "_", true, symbols)
}

func getSymbols(url string) (map[string][2]string, error) {
	var symbols map[string][2]string = map[string][2]string{}

	restyResp, err := resty.R().Get(url)

	if err != nil {
		return symbols, err
	}

	type Symbols struct {
		Data map[string]interface{} `json:data`
	}

	var resp Symbols

	err = json.Unmarshal(restyResp.Body(), &resp)

	if err != nil {
		return symbols, err
	}

	for symb, _ := range resp.Data {
		coins := strings.Split(symb, "_")
		symbols[symb] = [2]string{
			coins[0],
			coins[1],
		}
	}

	return symbols, nil
}
