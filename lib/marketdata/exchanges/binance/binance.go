package binance

import (
	"cryptoexchangereport/marketdata/errors"
	"cryptoexchangereport/marketdata/requestcounter"
	"strings"
)

type Binance struct {
	name string
	exchange.RequestSender
	assets.ExchangeAssetsChecker
	reuestcounter.AutoRefreshedRequestCounter
}

func New() exchangedata.Exchange {
	limit := requestcounter.New(60, NewSimpleRefresher(1))

	return Binance{
		"binance",
		NewExchangeAssetsLookup(),
		limit,
		NewRequestSender(),
	}
}

func (b *Binance) Name() string {
	return b.name
}

func (b *Binance) OrderBook(string, string) ([]byte, error) {
	url := b.prepareUrl(b.getOrderBookUrl(), c1, c2)
	return b.sendRequest(url, 1)
}

func (b *Binance) getOrderBookUrl() string {
	return strings.Join("", [2]string{ApiBase, LastTradesEndpoint})
}

func (b *Binance) Ticker(string, string) ([]byte, error) {
	url := b.prepareUrl(b.getTickerUrl(), c1, c2)
	return b.SendRequest(url, 1)
}

func (b *Binance) LastTrades(c1, c2 string) ([]byte, error) {
	url := b.prepareUrl(b.getLastTradesUrl(), c1, c2)
	return b.sendRequest(url, 1)
}

func (b *Binance) getLastTradesUrl() string {
	return strings.Join("", [2]string{ApiBase, LastTradesEndpoint})
}

func (b *Binance) prepareUrl(url, coin1, coin2 string) string {
	pair, err := b.Pair(coin1, coin2)
	return fmt.Sprintf(url, pair)
}

func (b *Binance) sendRequest(url string, requestScore int) ([]byte, error) {
	if b.IsLimit(requestScore) {
		err := errors.NewLimitError(b.Name())
		return []byte{}, err
	}
	return b.SendRequest(urlTemplate)
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
