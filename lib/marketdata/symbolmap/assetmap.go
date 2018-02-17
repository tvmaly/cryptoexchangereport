package assets

import (
    "fmt"
)

type AssetMap struct {
    exchange string
    assets map[string][]string
}

func NewAssetMap(exchange string, assets map[string][]string) *AssetMap {
    var am AssetMap = AssetMap{
        exchange: exchange,
        assets: assets,
    }

    return &am
}

func (am AssetMap) AddSymbol (asset string, newSymbol string) error {

    if am.Exist(asset) {

        symbols, err := am.Get(asset)

        if err != nil {
            return err
        }

        for _, currentSymbol := range symbols {

            if newSymbol == currentSymbol {
                err := fmt.Errorf("%s already assigned for %s", newSymbol, asset)
                return err
            }
        }

        am.assets[asset] = append(symbols, newSymbol)

    } else {
        am.assets[asset] = []string{newSymbol}

    }

    return nil
}

func (am AssetMap) Get (asset string) ([]string, error) {

    var symbols []string

    if !am.Exist(asset) {
        err := fmt.Errorf("There are no %s coin on %s", asset, am.exchange)
        return symbols, err

    } else {
        symbols = am.assets[asset]

        if len(symbols) == 0 {
            err := fmt.Errorf("Zero symbols for existing coin %s\n", asset)
            return []string{}, err
        }

        return symbols, nil
    }
}

func (am AssetMap) Exist (asset string) bool {
    if _, ok := am.assets[asset]; ok {
        return true
    } else {
        return false
    }
}
