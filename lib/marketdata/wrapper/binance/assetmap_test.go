package binance

import (
    "cryptoexchangereport/marketdata/assets"
    "testing"
)

func TestGetAssetMap (t *testing.T) {

    var assetMap *assets.AssetMap = GetAssetMap()

    if !assetMap.Exist("BTC") {
        t.Fatal("Binance has BTC")
    }

    symbols, err := assetMap.Get("BTC")

    if err != nil {
        t.Fatal(err)
    }

    oldLen := len(symbols)

    assetMap.AddSymbol("BTC", "BTC-KOKOKO")

    symbols, err = assetMap.Get("BTC")

    if len(symbols) - oldLen != 1 {
        t.Fatal("We added 1 symbol, so in new set there should be one more then in old one")
    }

    if assetMap.Exist("KOKOKO") {
        t.Fatal("No such coin as KOKOKO. At least for now")
    }
}
