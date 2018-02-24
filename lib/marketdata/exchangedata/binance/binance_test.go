package binance

import (
	"testing"
)

var exchange *Binance = New()

func TestOrderBook(t *testing.T) {
	orderBookUrl := exchange.OrderBookURLTemplate()

	expectedUrl := ApiBase + OrderBookEndpoint

	if orderBookUrl != expectedUrl {
		t.Errorf("Orderbook url is not correct %s - %s\n", orderBookUrl, expectedUrl)
	}
}

func TestTicker(t *testing.T) {
	tickerUrl := exchange.TickerURLTemplate()

	expectedUrl := ApiBase + TickerEndpoint

	if tickerUrl != ApiBase+TickerEndpoint {
		t.Errorf("Ticker url is not correct %s - %s\n", tickerUrl, expectedUrl)
	}
}

func TestLastTrades(t *testing.T) {
	lastTrades := exchange.LastTradesURLTemplate()

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
}
