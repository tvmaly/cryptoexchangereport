package poloniex

import (
	"cryptoexchangereport/marketdata/exchangedata"
)

type Poloniex struct {
	exchangedata.GenericExchange
}

func New() *Poloniex {
	return NewWithRequestSender(exchangedata.NewRequestSender())
}

func NewWithRequestSender(requestSender exchangedata.RequestSender) *Poloniex {
	var URL map[string]string = map[string]string{
		"OrderBookURLTemplate":  ApiBase + OrderBookEndpoint,
		"TickerURLTemplate":     ApiBase + TickerEndpoint,
		"LastTradesURLTemplate": ApiBase + LastTradesEndpoint,
	}

	genericExchange := exchangedata.NewGenericExchange("poloniex", 120, NewSymbolMap(), URL, requestSender)

	return &Poloniex{*genericExchange}
}
