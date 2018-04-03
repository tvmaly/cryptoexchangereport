package binance

import (
	"testing"
)

var exchange *Binance = New()

func TestOrderBookUrl(t *testing.T) {
	orderBookUrl := exchange.getOrderBookUrl()

	expectedUrl := ApiBase + OrderBookEndpoint

	if orderBookUrl != expectedUrl {
		t.Errorf("Orderbook url is not correct %s - %s\n", orderBookUrl, expectedUrl)
	}
}

func TestTickerUrl(t *testing.T) {
	tickerUrl := exchange.getTickerUrl()

	expectedUrl := ApiBase + TickerEndpoint

	if tickerUrl != ApiBase+TickerEndpoint {
		t.Errorf("Ticker url is not correct %s - %s\n", tickerUrl, expectedUrl)
	}
}

func TestLastTradesUrl(t *testing.T) {
	lastTrades := exchange.getLastTradesUrl()

	expectedUrl := ApiBase + LastTradesEndpoint

	if lastTrades != expectedUrl {
		t.Errorf("Last trades url is not correct %s - %s\n", lastTrades, expectedUrl)
	}
}

func TestLimit(t *testing.T) {
	if exchange.IsLimit(0) {
		t.Error("Should not give an error. 0 count.")
	}

	if !exchange.IsLimit(500) {
		t.Error("Should be an error. 500 count.")
	}

	exchange.OrderBook("BTC", "LTC")

	if exchange.GetCounter() != 1 {
		t.Errorf("Counter should be 1, as we requested OrderBook. We have: %d\n", exchange.GetCounter())
	}
}
