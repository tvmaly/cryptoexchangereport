package binance

import (
    "fmt"
//    "cryptopairs"
    "gopkg.in/resty.v1"
)

type Adapter struct {
    orderBookUrl string
    lastTradesUrl string
    tickerUrl string
    cryptopairsNormalizer cryptopairs.Normalizer
    limit int // just a placeholder
    counter int // just a placeholder
    errorCodes map[int]string // just a placeholder
}

func New() Adapter {
    return Adapter{
        orderBookUrl: "https://bittrex.com/api/v1.1/public/getorderbook?market=%s&type=both",
        lastTradesUrl: "https://bittrex.com/api/v1.1/public/getmarkethistory?market=%s",
        tickerUrl: "https://bittrex.com/api/v1.1/public/getticker?market=%s",
        crytopairsNormalizer: getCryptopairsNormalizer(),
        limit: 500, // this is just a placeholder value
        counter: 0,
        errorCodes: map[int]string{404: "not found"}, // this is just a placeholder value
    }
}

func (a Adapter) OrderBook (cryptoPair string) []byte {
    requestUrl := a.composeWithCurrencyPaira.orderBookUrl, cryptoPair)
    return a.SendRequest(requestUrl)
}

func (a Adapter) Ticker (cryptoPair string) []byte {
    requestUrl := a.composea.tickerUrl, cryptoPair)
    return a.SendRequest(requestUrl)
}

func (a Adapter) LastTrades (cryptoPair string) []byte {
    requestUrl := a.composea.lastTradesUrl, cryptoPair)
    return a.SendRequest(requestUrl)
}

func (a Adapter) SendRequest (url string) []byte {

    if a.counter >= a.limit {
        panic("You hit the limit on bittrex") 
    }

    resp, err := resty.R().Get(url)

    if err != nil {
        fmt.Println("Error while requesting data for: ", url)
        panic(fmt.Sprintf("%v", err))
    }

    return resp.Body()
}

func (a Adapter) getValidCurrencyPair (cryptoPair string) string {
    return a.cryptopairsNormalizer.Normalize(cryptoPair)
}

func (a Adapter) composeWithCurrencyPair (url, cryptoPair string) string {
    normalizedCryptopair := a.getValidCurrencyPair(cryptoPair)
    return fmt.Sprintf(url, normalizedCryptopair)
}
