package exchangedata

import (
	"encoding/json"
	"fmt"
	"log"

	"cryptoexchangereport/marketdata/assets"
)

type GenericExchange struct {
	*assets.SymbolMap
	RequestSender
	name  string
	limit int
	URL   map[string]string
}

func NewGenericExchange(name string, limit int, symbolMap *assets.SymbolMap, url map[string]string, requestSender RequestSender) *GenericExchange {
	return &GenericExchange{
		symbolMap,
		requestSender,
		name,
		limit,
		url,
	}
}

func (ge *GenericExchange) Name() string {
	return ge.name
}

func (ge *GenericExchange) OrderBook(s1, s2 string) ([]byte, error) {
	response := ge.GETRequest(ge.OrderBookURLTemplate(), s1, s2)
	return response, nil
}

func (ge *GenericExchange) OrderBookURLTemplate() string {
	return ge.URL["OrderBookURLTemplate"]
}

func (ge *GenericExchange) Ticker(s1, s2 string) ([]byte, error) {
	response := ge.GETRequest(ge.TickerURLTemplate(), s1, s2)
	return response, nil
}

func (ge *GenericExchange) TickerURLTemplate() string {
	return ge.URL["TickerURLTemplate"]
}

func (ge *GenericExchange) LastTrades(s1, s2 string) ([]byte, error) {
	response := ge.GETRequest(ge.LastTradesURLTemplate(), s1, s2)
	return response, nil
}

func (ge *GenericExchange) LastTradesURLTemplate() string {
	return ge.URL["LastTradesURLTemplate"]
}

func (ge *GenericExchange) GETRequest(urlTemplate, s1, s2 string) []byte {
	symbol, err := ge.GetSymbol(s1, s2)
	url := fmt.Sprintf(urlTemplate, symbol)

	response, err := ge.SendRequest(url, symbol)

	if err != nil {
		log.Fatalf("Error occured during %s GET request to %s: %v\n", ge.Name(), url, err)
	}

	marshalledResp, err := json.Marshal(response)

	if err != nil {
		log.Fatalf("Error while marshalling %v for %s\n", response, ge.Name())
	}

	return marshalledResp
}

func (ge *GenericExchange) Limit() int {
	return ge.limit
}

func (ge *GenericExchange) IsLimit(count int) bool {
	return ge.Limit()-count <= 5
}
