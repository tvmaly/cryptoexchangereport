package bitfinex

import (
    "cryptoexchangereport/marketdata/assets"
)

func GetAssetMap() *assets.AssetMap {

    symbolMap := GetSymbolMap()

    return symbolMap.BuildAssetMap()
}
