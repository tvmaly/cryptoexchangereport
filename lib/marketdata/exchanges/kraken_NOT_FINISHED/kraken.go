// Needs authentication for higher API requests limit!
package kraken

import (
	"cryptoexchangereport/marketdata/exchangedata"
)

type Kraken struct {
	exchangedata.GenericExchange
}

func New() *Kraken {
	return NewWithRequestSender(exchangedata.NewRequestSender())
}

func NewWithRequestSender(requestSender exchangedata.RequestSender) *Kraken {
	var URL map[string]string = map[string]string{
		"OrderBookURLTemplate":  ApiBase + OrderBookEndpoint,
		"TickerURLTemplate":     ApiBase + TickerEndpoint,
		"LastTradesURLTemplate": ApiBase + LastTradesEndpoint,
	}

	genericExchange := exchangedata.NewGenericExchange("kraken", 10, NewSymbolMap(), URL, requestSender)

	return &Kraken{*genericExchange}
}
