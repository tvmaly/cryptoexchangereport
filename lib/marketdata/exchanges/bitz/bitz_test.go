package bitz

import (
	"testing"
)

var exchange *BitZ = New()

func TestOrderBook(t *testing.T) {
	orderBookUrl := exchange.OrderBookURLTemplate()

	expectedUrl := ApiBase + OrderBookEndpoint

	if orderBookUrl != expectedUrl {
		t.Errorf("Orderbook url isn't correct %s - %s\n", orderBookUrl, expectedUrl)
	}
}

func TestTicker(t *testing.T) {
	tickerUrl := exchange.TickerURLTemplate()

	expectedUrl := ApiBase + TickerEndpoint

	if tickerUrl != ApiBase+TickerEndpoint {
		t.Errorf("Ticker url is not correct %s - %s\n", tickerUrl, expectedUrl)
	}
	/*
		type Ticker struct {
			Data interface{} `json:data`
		}

		url := fmt.Sprintf(tickerUrl, "ETH_BTC")
		resp, _ := exchange.SendRequest(url, "ETH_BTC")

		var tick Ticker

		json.Unmarshal(resp.Response, &tick)

		fmt.Println(tick)
	*/
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
