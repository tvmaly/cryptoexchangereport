package bittrex

import (
    "log"
    "encoding/json"
    "cryptoexchangereport/marketdata/symbolmap"
    "gopkg.in/resty.v1"
)

func getSymbolMap() *symbolmap.SymbolMap {

    symbols, err := getSymbols(MarketsUrl)

    if err != nil {
        log.Fatalf("Error while getting symbols from bittrex: %v", err)
    }
    
    return symbolmap.New("bittrex", "-", symbols)
}

func getSymbols (url string) (map[string]bool, error) {

    var symbols map[string]bool = map[string]bool{}

    resp, err := resty.R().Get(url)

    if err != nil {
        return symbols, err
    }

    type Market struct {
        MarketName string `json:MarketName`
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
        symbols[(*market).MarketName] = true
    }

    return symbols, nil
}
