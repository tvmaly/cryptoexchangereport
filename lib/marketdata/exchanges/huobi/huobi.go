package huobi

import (
	"cryptoexchangereport/marketdata/exchangedata"
)

type Huobi struct {
	exchangedata.GenericExchange
}

func New() *Huobi {
	return NewWithRequestSender(exchangedata.NewRequestSender())
}

func NewWithRequestSender(requestSender exchangedata.RequestSender) *Huobi {
	var URL map[string]string = map[string]string{
		"OrderBookURLTemplate":  ApiBase + OrderBookEndpoint,
		"TickerURLTemplate":     ApiBase + TickerEndpoint,
		"LastTradesURLTemplate": ApiBase + LastTradesEndpoint,
	}

	genericExchange := exchangedata.NewGenericExchange("huobi", 100, NewSymbolMap(), URL, requestSender)

	return &Huobi{*genericExchange}
}
