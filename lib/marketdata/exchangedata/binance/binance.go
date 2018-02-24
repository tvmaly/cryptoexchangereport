package binance

import (
	"cryptoexchangereport/marketdata/exchangedata"
)

type Binance struct {
	exchangedata.GenericExchange
}

func New() *Binance {
	return NewWithRequestSender(exchangedata.NewRequestSender())
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
