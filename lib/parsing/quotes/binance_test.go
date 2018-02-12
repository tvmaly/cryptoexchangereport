package quotes

import (
	"fmt"
	"testing"
)

func TestNewBinanceQuotes(t *testing.T) {

	quotes, err := NewBinanceQuotes(BinanceTicker())

	if err != nil {
		t.Errorf("NewBinanceQuotes Failed: %s", err)
	} else {

		for _, v := range quotes {

			fmt.Printf("Symbol: %s MidPoint: %f\n", v.Symbol, v.MidPoint())

		}
	}

}

func BinanceTicker() []byte {

	return []byte(`[{"symbol":"ETHBTC","bidPrice":"0.10104800","bidQty":"18.00700000","askPrice":"0.10105000","askQty":"0.00800000"},{"symbol":"LTCBTC","bidPrice":"0.01806000","bidQty":"2.21000000","askPrice":"0.01811000","askQty":"7.87000000"},{"symbol":"BNBBTC","bidPrice":"0.00097100","bidQty":"34.88000000","askPrice":"0.00097510","askQty":"25.82000000"}]`)

}
