package bitfinex

import (
	"cryptoexchangereport/marketdata/exchangedata"
)

type Bitfinex struct {
	exchangedata.GenericExchange
}

func New() *Bitfinex {
	return NewWithRequestSender(exchangedata.NewRequestSender())
}

func NewWithRequestSender(requestSender exchangedata.RequestSender) *Bitfinex {
	var URL map[string]string = map[string]string{
		"OrderBookURLTemplate":  ApiBase + OrderBookEndpoint,
		"TickerURLTemplate":     ApiBase + TickerEndpoint,
		"LastTradesURLTemplate": ApiBase + LastTradesEndpoint,
	}

	genericExchange := exchangedata.NewGenericExchange("bitfinex", 80, NewSymbolMap(), URL, requestSender)

	return &Bitfinex{*genericExchange}
}
