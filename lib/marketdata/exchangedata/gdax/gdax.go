package gdax

import (
	"cryptoexchangereport/marketdata/exchangedata"
)

type Gdax struct {
	exchangedata.GenericExchange
}

func New() *Gdax {
	return NewWithRequestSender(exchangedata.NewRequestSender())
}

func NewWithRequestSender(requestSender exchangedata.RequestSender) *Gdax {
	var URL map[string]string = map[string]string{
		"OrderBookURLTemplate":  ApiBase + OrderBookEndpoint,
		"TickerURLTemplate":     ApiBase + TickerEndpoint,
		"LastTradesURLTemplate": ApiBase + LastTradesEndpoint,
	}

	genericExchange := exchangedata.NewGenericExchange("gdax", 120, NewSymbolMap(), URL, requestSender)

	return &Gdax{*genericExchange}
}
