package assets

import (
    "testing"
)

var assetsTestMap map[string][]string = map[string][]string{
    "BTC": []string{"ETH-BTC", "LTC-BTC", "BCH-BTC"},
    "ETH": []string{"BTC-ETH", "OMG-ETH"},
}

func TestAssetMap (t *testing.T) {

    assetMap := NewAssetMap("bittrex", assetsTestMap)

    assetMap.AddSymbol("BTC", "EOS-BTC")

    if !assetMap.Exist("BTC") {
        t.Fatal("BTC should exist")
    }

    symbols, err := assetMap.Get("BTC")

    if err != nil {
        t.Fatal(err)
    }

    if len(symbols) != 4 {
        t.Fatalf("Should be 4 symbols already, %v", symbols)
    }


}
