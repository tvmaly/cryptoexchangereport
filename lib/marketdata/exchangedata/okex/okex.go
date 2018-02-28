package okex

import (
	"cryptoexchangereport/marketdata/exchangedata"
)

type Okex struct {
	exchangedata.GenericExchange
}

func New() *Okex {
	return NewWithRequestSender(exchangedata.NewRequestSender())
}

func NewWithRequestSender(requestSender exchangedata.RequestSender) *Okex {
	var URL map[string]string = map[string]string{
		"OrderBookURLTemplate":  ApiBase + OrderBookEndpoint,
		"TickerURLTemplate":     ApiBase + TickerEndpoint,
		"LastTradesURLTemplate": ApiBase + LastTradesEndpoint,
	}

	genericExchange := exchangedata.NewGenericExchange("okex", 120, NewSymbolMap(), URL, requestSender)

	return &Okex{*genericExchange}
}
