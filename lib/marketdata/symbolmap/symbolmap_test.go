package symbolmap

import (
    "testing"
)

func TestGetSymbol (t *testing.T) {

    var symbolsTestMap map[string]bool = map[string]bool{
        "LTC-BTC": true,
        "ETH-BTC": true,
        "ENG-ETH": true,
    }

    symbolMap := New("bittrex", "-", symbolsTestMap)

    mySymbol, err := symbolMap.GetSymbol("BTC", "LTC")

    if err != nil {
        t.Fatal(err)
    }

    mySymbol, err = symbolMap.GetSymbol("btc", "ltc")

    if err != nil {
        t.Fatal(err)
    }

    mySymbol, err = symbolMap.GetSymbol("BTC", "WAKA")

    if err == nil {
        t.Fatalf("BTC-WAKA does not exists, should have to output error, but instead gave %s", mySymbol)
    }
}
