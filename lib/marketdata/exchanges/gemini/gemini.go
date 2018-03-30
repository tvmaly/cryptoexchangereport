package gemini

import (
	"cryptoexchangereport/marketdata/exchangedata"
)

type Gemini struct {
	exchangedata.GenericExchange
}

func New() *Gemini {
	return NewWithRequestSender(exchangedata.NewRequestSender())
}

func NewWithRequestSender(requestSender exchangedata.RequestSender) *Gemini {
	var URL map[string]string = map[string]string{
		"OrderBookURLTemplate":  ApiBase + OrderBookEndpoint,
		"TickerURLTemplate":     ApiBase + TickerEndpoint,
		"LastTradesURLTemplate": ApiBase + LastTradesEndpoint,
	}

	genericExchange := exchangedata.NewGenericExchange("gemini", 120, NewSymbolMap(), URL, requestSender)

	return &Gemini{*genericExchange}
}
