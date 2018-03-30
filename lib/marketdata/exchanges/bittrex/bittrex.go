package bittrex

import (
	"cryptoexchangereport/marketdata/exchangedata"
)

type Bittrex struct {
	exchangedata.GenericExchange
}

func New() *Bittrex {
	return NewWithRequestSender(exchangedata.NewRequestSender())
}

func NewWithRequestSender(requestSender exchangedata.RequestSender) *Bittrex {
	var URL map[string]string = map[string]string{
		"OrderBookURLTemplate":  ApiBase + OrderBookEndpoint,
		"TickerURLTemplate":     ApiBase + TickerEndpoint,
		"LastTradesURLTemplate": ApiBase + LastTradesEndpoint,
	}

	genericExchange := exchangedata.NewGenericExchange("bittrex", 500, NewSymbolMap(), URL, requestSender)

	return &Bittrex{*genericExchange}
}
