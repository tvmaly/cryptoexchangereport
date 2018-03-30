package gemini

import (
	"testing"
)

func TestNewSymbolMap(t *testing.T) {
	symbolMap := NewSymbolMap()

	symbol, err := symbolMap.GetSymbol("BTC", "USD")

	if err != nil {
		t.Fatal(err)
	}

	symbol, err = symbolMap.GetSymbol("usd", "btc")

	if err != nil {
		t.Fatal(err)
	}

	symbol, err = symbolMap.GetSymbol("btc", "ltc")

	if err == nil {
		t.Fatalf("%s does not exists on gemini", symbol)
	}

	assets := symbolMap.GetSymbolsForAsset("BTC")

	if len(assets) == 0 {
		t.Fatalf("there are should be symbols for gemini")
	}
}
