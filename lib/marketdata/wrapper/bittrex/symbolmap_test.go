package bittrex

import (
    "testing"
)

func TestGetSymbolMap (t *testing.T) {

    symbolMap := GetSymbolMap()

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
        t.Fatalf("%s does not exists on bittrex", symbol)
    }

    assets := symbolMap.GetSymbolsForAsset("BTC")

    if len(assets) == 0 {
        t.Fatalf("there are should be symbols for bittrex")
    }
}
