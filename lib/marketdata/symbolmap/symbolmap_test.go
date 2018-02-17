package assets

import (
    "testing"
)

var symbolsTestMap map[string][2]string = map[string][2]string {
    "LTC-BTC": [2]string{"LTC", "BTC"},
    "ETH-BTC": [2]string{"ETH", "BTC"},
    "ENG-ETH": [2]string{"ENG", "ETH"},
}

func TestGetSymbol (t *testing.T) {

    symbolMap := NewSymbolMap("bittrex", "-", symbolsTestMap)

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

func TestGetSymbolsForAsset (t *testing.T) {

    symbolMap := NewSymbolMap("bittrex", "-", symbolsTestMap)

    symbols := symbolMap.GetSymbolsForAsset("BTC")

    for _, symbol := range symbols {
        if symbol != "LTC-BTC" && symbol != "ETH-BTC" {
            t.Fatalf("symbol should be either LTC-BTC or ETH-BTC, but instead we got %s\n", symbol)
        }
    }

    symbols = symbolMap.GetSymbolsForAsset("ENG")

    if len(symbols) > 1 {
        t.Fatal("There are only one symbol for ENG coin")
    }

    if symbols[0] != "ENG-ETH" {
        t.Fatalf("should be ENG-ETH for ETH")
    }
}

func TestGetAssetMap (t *testing.T) {

    symbolMap := NewSymbolMap("bittrex", "-", symbolsTestMap)

    assetMap := symbolMap.BuildAssetMap()

    if !assetMap.Exist("BTC") {
        t.Fatal("BTC should exist")
    }
}
