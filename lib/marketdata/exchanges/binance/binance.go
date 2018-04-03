package binance

import (
	"fmt"
	"log"
	"strings"

	"cryptoexchangereport/marketdata/assets"
	"cryptoexchangereport/marketdata/errors"
	"cryptoexchangereport/marketdata/exchanges"
	"cryptoexchangereport/marketdata/requestcounter"
)

type Binance struct {
	assets.ExchangeAssetsChecker
	exchanges.RequestSender
	requestcounter.AutoRefreshedRequestCounter
	name string
}

func New() *Binance {
	limit := requestcounter.New(60, requestcounter.NewSimpleRefresher(1))

	return &Binance{
		NewExchangeAssetsEnum(),
		exchanges.NewRequestSender(),
		limit,
		"binance",
	}
}

func (b *Binance) Name() string {
	return b.name
}

func (b *Binance) OrderBook(c1, c2 string) (exchanges.Response, error) {
	url := b.prepareUrl(b.getOrderBookUrl(), c1, c2)
	return b.sendRequest(url, 1)
}

func (b *Binance) getOrderBookUrl() string {
	return strings.Join([]string{ApiBase, OrderBookEndpoint}, "")
}

func (b *Binance) Ticker(c1, c2 string) (exchanges.Response, error) {
	url := b.prepareUrl(b.getTickerUrl(), c1, c2)
	return b.sendRequest(url, 1)
}

func (b *Binance) getTickerUrl() string {
	return strings.Join([]string{ApiBase, TickerEndpoint}, "")
}

func (b *Binance) LastTrades(c1, c2 string) (exchanges.Response, error) {
	url := b.prepareUrl(b.getLastTradesUrl(), c1, c2)
	return b.sendRequest(url, 1)
}

func (b *Binance) getLastTradesUrl() string {
	return strings.Join([]string{ApiBase, LastTradesEndpoint}, "")
}

func (b *Binance) prepareUrl(url, coin1, coin2 string) string {
	pair, err := b.Pair(coin1, coin2)

	if err != nil {
		log.Fatalf("Could not find pair for coins %s - %s: %v\n", coin1, coin2, err)
	}

	return fmt.Sprintf(url, pair)
}

func (b *Binance) sendRequest(url string, requestScore int) (exchanges.Response, error) {
	if b.IsLimit(requestScore) {
		err := errors.NewLimitError(b.Name())
		return exchanges.Response{}, err
	}
	b.Hit(1)
	return b.SendRequest(url)
}
