package bittrex

import (
    "fmt"
    "gopkg.in/resty.v1"
)

type BittrexAdapter struct {
    marketsUrl string
    orderBookUrl string
    lastTradesUrl string
    tickerUrl string
    limit int
    counter int
    errorCodes map[int]string
}

func New() BittrexAdapter {
    return BittrexAdapter{
        marketsUrl: "https://bittrex.com/api/v1.1/public/getmarkets",
        orderBookUrl: "https://bittrex.com/api/v1.1/public/getorderbook?market=%s&type=both",
        lastTradesUrl: "https://bittrex.com/api/v1.1/public/getmarkethistory?market=%s",
        tickerUrl: "https://bittrex.com/api/v1.1/public/getticker?market=%s",
        limit: 500, // this is just a placeholder value
        errorCodes: map[int]string{404: "not found"}, // this is just a placeholder value
    }
}

func (b BittrexAdapter) OrderBook (market string) []byte {
    requestUrl := b.compose(b.orderBookUrl, market)
    return b.SendRequest(requestUrl)
}

func (b BittrexAdapter) Ticker (market string) []byte {
    requestUrl := b.compose(b.tickerUrl, market)
    return b.SendRequest(requestUrl)
}

func (b BittrexAdapter) LastTrades (market string) []byte {
    requestUrl := b.compose(b.lastTradesUrl, market)
    return b.SendRequest(requestUrl)
}

func (b BittrexAdapter) SendRequest (url string) []byte {

    if b.counter >= b.limit {
        panic("You hit the limit on bittrex") 
    }

    resp, err := resty.R().Get(url)

    if err != nil {
        fmt.Println("Error while requesting data for: ", url)
        panic(fmt.Sprintf("%v", err))
    }

    return resp.Body()
}

func (b BittrexAdapter) compose (url, market string) string {
    return fmt.Sprintf(url, market)
}
