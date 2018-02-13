package bittrex

import (
    "fmt"
    "gopkg.in/resty.v1"
)

type Markets struct {
    markets []*Market `json:result`
}

type Market struct {
    MarketName string `json:MarketName`
}

func BuildMarkets (url string) map[string]bool {
    resp, err := resty.R().Get(Url)A

    var markets map[string]bool

    var marketsResp := Markets

    err := json.Unmarshall(resp.Body, &marketsResp.Body())

    if err != nil {
        panic(fmt.Sprintf("%v", err))
    }

    for _, market := range marketsREsp.markets {
        markets[market] = true
    }

    return markets
}
