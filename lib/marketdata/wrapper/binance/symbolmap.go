package binance

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

    return symbolmap.New("binance", "", symbols)
}

func getSymbols (url string) (map[string]bool, error) {

    var symbols map[string]bool = map[string]bool{}

    resp, err := resty.R().Get(url)

    if err != nil {
        return symbols, err
    }

    type SymbolJSON struct {
        Symbol string `json:symbol`
    }

    type SymbolsJSON struct {
        Symbols []*SymbolJSON `json:symbols`
    }

    var symbolJSON []*SymbolJSON
    var  symbolsJSON = SymbolsJSON{Symbols: symbolJSON}

    err = json.Unmarshal(resp.Body(), &symbolsJSON)

    if err != nil {
        return symbols, err
    }

    for _, sym := range symbolsJSON.Symbols {
        symbols[(*sym).Symbol] = true
    }

    return symbols, nil
}
