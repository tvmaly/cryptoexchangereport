package bitstamp

import (
	"cryptoexchangereport/marketdata/exchangedata"
)

type Bitstamp struct {
	exchangedata.GenericExchange
}

func New() *Bitstamp {
	return NewWithRequestSender(exchangedata.NewRequestSender())
}

func NewWithRequestSender(requestSender exchangedata.RequestSender) *Bitstamp {
	var URL map[string]string = map[string]string{
		"OrderBookURLTemplate":  ApiBase + OrderBookEndpoint,
		"TickerURLTemplate":     ApiBase + TickerEndpoint,
		"LastTradesURLTemplate": ApiBase + LastTradesEndpoint,
	}

	genericExchange := exchangedata.NewGenericExchange("bitstamp", 60, NewSymbolMap(), URL, requestSender)

	return &Bitstamp{*genericExchange}
}
