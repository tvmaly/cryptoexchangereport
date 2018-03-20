package binance

import (
	"cryptoexchangereport/marketdata/exchangedata"
)

type Binance struct {
	assets.ExchangeAssetsChecker
	name string
    limit exchangedata.LimitCounter
    requestSender exchangedata.RequestSender
}

func New() exchangedata.Exchange {
	return
}

func (b *Binance) Name() string {
	return b.name
}

func (b *Binance) OrderBook(string, string) ([]byte, error) {
}

func (b *Binance) Ticker(string, string) ([]byte, error) {
}

func (b *Binance) LastTrades(string, string) ([]byte, error) {
}

func (b *Binance) prepareURL (url, coin1, coin2 string) string {
    pair, err := b.Pair(coin1, coin2)
    url := fmt.Sprintf(url, pair)

    response, err := b.SendRequest(urlTemplate
}

func NewWithRequestSender(requestSender exchangedata.RequestSender) *Binance {
	var URL map[string]string = map[string]string{
		"OrderBookURLTemplate":  ApiBase + OrderBookEndpoint,
		"TickerURLTemplate":     ApiBase + TickerEndpoint,
		"LastTradesURLTemplate": ApiBase + LastTradesEndpoint,
	}

	genericExchange := exchangedata.NewGenericExchange("binance", 500, NewSymbolMap(), URL, requestSender)

	return &Binance{*genericExchange}
}
