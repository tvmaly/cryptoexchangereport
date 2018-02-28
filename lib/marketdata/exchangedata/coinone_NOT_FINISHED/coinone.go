package coinone

import (
	"cryptoexchangereport/marketdata/exchangedata"
)

type Coinone struct {
	exchangedata.GenericExchange
}

func New() *Coinone {
	return NewWithRequestSender(exchangedata.NewRequestSender())
}

func NewWithRequestSender(requestSender exchangedata.RequestSender) *Coinone {
	var URL map[string]string = map[string]string{
		"OrderBookURLTemplate":  ApiBase + OrderBookEndpoint,
		"TickerURLTemplate":     ApiBase + TickerEndpoint,
		"LastTradesURLTemplate": ApiBase + LastTradesEndpoint,
	}

	genericExchange := exchangedata.NewGenericExchange("coinone", 500, NewSymbolMap(), URL, requestSender)

	return &Coinone{*genericExchange}
}
