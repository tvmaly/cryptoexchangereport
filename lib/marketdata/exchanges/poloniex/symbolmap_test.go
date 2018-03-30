package poloniex

import (
	"testing"
)

func TestNewSymbolMap(t *testing.T) {

	symbolMap := NewSymbolMap()

	symbol, err := symbolMap.GetSymbol("BTC", "LTC")

	if err != nil {
		t.Fatal(err)
	}

	symbol, err = symbolMap.GetSymbol("ltc", "btc")

	if err != nil {
		t.Fatal(err)
	}

	symbol, err = symbolMap.GetSymbol("this", "nott")

	if err == nil {
		t.Fatalf("%s does not exists on poloniex", symbol)
	}

	assets := symbolMap.GetSymbolsForAsset("BTC")

	if len(assets) == 0 {
		t.Fatalf("there are should be symbols for poloniex")
	}
}
