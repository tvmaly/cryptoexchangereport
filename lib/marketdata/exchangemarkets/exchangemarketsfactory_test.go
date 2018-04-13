package exchangemarkets

import (
	"testing"
)

var exchangeMarketsFactory ExchangeMarketsFactory = ExchangeMarketsFactory{}

func TestSuccessfulFactory(t *testing.T) {
	exchangeMarketsFactory.Register("binance", GetExchangeMarkets)
	binanceMarkets := exchangeMarketsFactory.GetExchangeMarkets("binance")

	if binanceMarkets == nil {
		t.Fatal("Binance is listed, so should be created!")
	}
}

func TestFailingFactory(t *testing.T) {
	kavabangaMarkets := exchangeMarketsFactory.GetExchangeMarkets("kavabanga")

	if kavabangaMarkets != nil {
		t.Fatalf("There are no such exchange as Kavabanga registered in our system")
	}
}

func GetExchangeMarkets() ExchangeMarkets {
	return ExchangeMarketsRepository{}
}
