package hitbtc

import (
	"cryptoexchangereport/marketdata/exchangedata"
)

type Hitbtc struct {
	exchangedata.GenericExchange
}

func New() *Hitbtc {
	return NewWithRequestSender(exchangedata.NewRequestSender())
}

func NewWithRequestSender(requestSender exchangedata.RequestSender) *Hitbtc {
	var URL map[string]string = map[string]string{
		"OrderBookURLTemplate":  ApiBase + OrderBookEndpoint,
		"TickerURLTemplate":     ApiBase + TickerEndpoint,
		"LastTradesURLTemplate": ApiBase + LastTradesEndpoint,
	}

	genericExchange := exchangedata.NewGenericExchange("hitbtc", 120, NewSymbolMap(), URL, requestSender)

	return &Hitbtc{*genericExchange}
}
