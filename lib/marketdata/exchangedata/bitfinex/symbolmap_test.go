package bitfinex

import (
	"testing"
)

func TestNewSymbolMap(t *testing.T) {

	symbolmap := NewSymbolMap()

	symbol, err := symbolmap.GetSymbol("BTC", "LTC")

	if err != nil {
		t.Fatal(err)
	}

	symbol, err = symbolmap.GetSymbol("ltc", "btc")

	if err != nil {
		t.Fatal(err)
	}

	symbol, err = symbolmap.GetSymbol("this", "nott")

	if err == nil {
		t.Fatalf("%s does not exists on bitfinex", symbol)
	}
}
