package binance

import (
    "fmt"
    "log"
//    "cryptopairs"
    "gopkg.in/resty.v1"
)

const MarketsUrl = "https://bittrex.com/api/v1.1/public/getmarkets"

func getCryptopairsNormalizer () cryptopairs.Normalizer {

    pairs, err := getMarkets(MarketsUrl)

    if err != nil {
        log.Fatal(fmt.Sprintf("Error while getting markets from bittrex: %v", err))
    }

//    return cryptopairs.NewNormalizer(pairs, "-")
    return nil
}

func getMarkets (url string) map[string]bool {

    resp, err := resty.R().Get(MarketsUrl)

    if err != nil {
        return _, err
    }

    type Market struct {
        MarketName string `json:MarketName`
    }

    type Result struct {
        Result []*Market `json:result`
    }

    var markets []*Market
    var result Result{Result: markets}

    err := json.Unmarshall(resp.Body(), &result)

    if err != nil {
        return nil, err
    }

    for _, market := range marketsREsp.markets {
        markets[market] = true
    }

    return markets, nil
}
