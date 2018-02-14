package bittrex

import (
    "fmt"
    "log"
    "encoding/json"
//    "cryptopairs"
    "gopkg.in/resty.v1"
)

const MarketsUrl = "https://bittrex.com/api/v1.1/public/getmarkets"

//func getCryptopairsNormalizer () cryptopairs.Normalizer {
func getCryptopairsNormalizer() {

    pairs, err := getMarkets(MarketsUrl)

    if err != nil {
        log.Fatal(fmt.Sprintf("Error while getting markets from bittrex: %v", err))
    }
    
    fmt.Println(pairs)

//    return cryptopairs.NewNormalizer(pairs, "-")
}

func getMarkets (url string) (map[string]bool, error) {

    var currencyPairs map[string]bool

    resp, err := resty.R().Get(MarketsUrl)

    if err != nil {
        return nil, err
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
        return nil, err
    }

    for _, market := range result.Result {
        currencyPairs[(*market).MarketName] = true
    }

    return currencyPairs, nil
}
