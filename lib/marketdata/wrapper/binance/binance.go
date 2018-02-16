package binance

import (
    "fmt"
    "log"
    "gopkg.in/resty.v1"
    "cryptoexchangereport/marketdata/symbolmap"
)

type Adapter struct {
    symbolMap *symbolmap.SymbolMap
    limit int // just a placeholder
    counter int // just a placeholder
    errorCodes map[int]string // just a placeholder
}

func New() Adapter {

    return Adapter{
        symbolMap: getSymbolMap(),
        limit: 500, // this is just a placeholder value
        counter: 0,
        errorCodes: map[int]string{404: "not found"}, // this is just a placeholder value
    }
}

func (a Adapter) OrderBook (s1, s2 string) []byte {
    requestUrl := a.composeWithSymbol(OrderBookUrl, s1, s2)
    return a.SendRequest(requestUrl).Body()
}

func (a Adapter) Ticker (s1, s2 string) []byte {
    requestUrl := a.composeWithSymbol(TickerUrl, s1, s2)
    return a.SendRequest(requestUrl).Body()
}

func (a Adapter) LastTrades (s1, s2 string) []byte {
    requestUrl := a.composeWithSymbol(LastTradesUrl, s1, s2)
    return a.SendRequest(requestUrl).Body()
}

func (a Adapter) SendRequest (url string) *resty.Response {

    if a.counter >= a.limit {
        log.Fatal("You hit the limit on bittrex")
    }

    resp, err := resty.R().Get(url)

    if err != nil {
        log.Fatalf("Error while requsting data from %s: %v", url, err)
    }

    return resp
}

func (a Adapter) getSymbol (s1, s2 string) string {
    symbol, err := a.symbolMap.GetSymbol(s1, s2)

    if err != nil {
        log.Fatal(err)
    }

    return symbol
}

func (a Adapter) composeWithSymbol (url, s1, s2 string) string {
    symbol := a.getSymbol(s1, s2)
    return fmt.Sprintf(url, symbol)
}
