package binance

import (
	"testing"
)

func TestNewSymbolMap(t *testing.T) {
	assetenum := NewExchangeAssetsEnum()

	symbol, err := assetenum.Pair("BTC", "LTC")

	if err != nil {
		t.Fatal(err)
	}

	symbol, err = assetenum.Pair("ltc", "btc")

	if err != nil {
		t.Fatal(err)
	}

	symbol, err = assetenum.Pair("this", "nott")

	if err == nil {
		t.Fatalf("%s does not exists on binance", symbol)
	}
}
