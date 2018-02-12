package quotes

import (
	"fmt"
	"testing"
)

func TestNewPoloniexQuotes(t *testing.T) {

	quotes, err := NewPoloniexQuotes(PoloniexTicker())

	if err != nil {
		t.Errorf("NewPoloniexQuotes Failed: %s", err)
	} else {

		for _, v := range quotes {

			fmt.Printf("Symbol: %s MidPoint: %f\n", v.Symbol, v.MidPoint())

		}
	}

}

func PoloniexTicker() []byte {

	return []byte(`{"BTC_BCN":{"id":7,"last":"0.00000056","lowestAsk":"0.00000056","highestBid":"0.00000055","percentChange":"0.05660377","baseVolume":"70.97455071","quoteVolume":"131725384.93695521","isFrozen":"0","high24hr":"0.00000057","low24hr":"0.00000052"},"BTC_BELA":{"id":8,"last":"0.00001762","lowestAsk":"0.00001762","highestBid":"0.00001761","percentChange":"-0.06076759","baseVolume":"7.20126434","quoteVolume":"405523.25254952","isFrozen":"0","high24hr":"0.00001906","low24hr":"0.00001700"}}`)

}
