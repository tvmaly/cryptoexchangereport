package bitz

import (
	"cryptoexchangereport/marketdata/exchangedata"
)

type BitZ struct {
	exchangedata.GenericExchange
}

func New() *BitZ {
	return NewWithRequestSender(exchangedata.NewRequestSender())
}

func NewWithRequestSender(requestSender exchangedata.RequestSender) *BitZ {
	var URL map[string]string = map[string]string{
		"OrderBookURLTemplate":  ApiBase + OrderBookEndpoint,
		"TickerURLTemplate":     ApiBase + TickerEndpoint,
		"LastTradesURLTemplate": ApiBase + LastTradesEndpoint,
	}

	genericExchange := exchangedata.NewGenericExchange("BitZ", 500, NewSymbolMap(), URL, requestSender)

	return &BitZ{*genericExchange}
}
