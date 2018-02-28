package huobi

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
		log.Fatalf("Error while getting symbols from huobi: %v", err)
	}

	return assets.NewSymbolMap("huobi", "", false, symbols)
}

func getSymbols(url string) (map[string][2]string, error) {
	var symbols map[string][2]string = map[string][2]string{}

	var base string
	var quote string
	var symbol string
	var coins [2]string

	resp, err := resty.R().Get(url)

	if err != nil {
		return symbols, err
	}

	type Market struct {
		Data []map[string]interface{} `json:data`
	}

	var markets Market

	err = json.Unmarshal(resp.Body(), &markets)

	if err != nil {
		return symbols, err
	}

	for _, market := range markets.Data {
		base = market["base-currency"].(string)
		quote = market["quote-currency"].(string)

		coins = [2]string{base, quote}
		symbol = strings.Join(coins[:], "")

		symbols[symbol] = coins
	}

	return symbols, nil
}
