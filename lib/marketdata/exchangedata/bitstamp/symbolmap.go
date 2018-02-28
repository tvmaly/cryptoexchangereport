package bitstamp

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
		log.Fatalf("Error while getting symbols from bitstamp: %v", err)
	}

	return assets.NewSymbolMap("bitstamp", "", false, symbols)
}

func getSymbols(url string) (map[string][2]string, error) {
	var symbols map[string][2]string = map[string][2]string{}

	resp, err := resty.R().Get(url)

	if err != nil {
		return symbols, err
	}

	type TradingPair struct {
		Name       string `json:name`
		Url_Symbol string `json:url_symbol`
	}

	var tradpairs []*TradingPair

	err = json.Unmarshal(resp.Body(), &tradpairs)

	if err != nil {
		return symbols, err
	}

	for _, tp := range tradpairs {

		coins := strings.Split(tp.Name, "/")

		symbols[tp.Url_Symbol] = [2]string{
			coins[0],
			coins[1],
		}
	}

	return symbols, nil
}
