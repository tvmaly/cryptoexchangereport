package liqui

import (
	"cryptoexchangereport/marketdata/exchangedata"
)

type Liqui struct {
	exchangedata.GenericExchange
}

func New() *Liqui {
	return NewWithRequestSender(exchangedata.NewRequestSender())
}

func NewWithRequestSender(requestSender exchangedata.RequestSender) *Liqui {
	var URL map[string]string = map[string]string{
		"OrderBookURLTemplate":  ApiBase + OrderBookEndpoint,
		"TickerURLTemplate":     ApiBase + TickerEndpoint,
		"LastTradesURLTemplate": ApiBase + LastTradesEndpoint,
	}

	genericExchange := exchangedata.NewGenericExchange("liqui", 120, NewSymbolMap(), URL, requestSender)

	return &Liqui{*genericExchange}
}
