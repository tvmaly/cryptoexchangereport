package assets

import (
	"fmt"
	"strings"
)

type SymbolMap struct {
	exchange  string
	symbols   map[string][2]string
	delimiter string
}

func NewSymbolMap(exchange, delimiter string, symbols map[string][2]string) *SymbolMap {

	for symbol, coin := range symbols {
		delete(symbols, symbol)
		symbols[strings.ToUpper(symbol)] = [2]string{
			strings.ToUpper(coin[0]),
			strings.ToUpper(coin[1]),
		}
	}

	var cp SymbolMap = SymbolMap{
		exchange:  exchange,
		symbols:   symbols,
		delimiter: delimiter,
	}

	return &cp
}

func (cp SymbolMap) GetSymbol(s1, s2 string) (string, error) {

	s1 = strings.ToUpper(s1)
	s2 = strings.ToUpper(s2)

	market := strings.Join([]string{s1, s2}, cp.delimiter)

	if _, ok := cp.symbols[market]; ok {
		return market, nil
	}

	market = strings.Join([]string{s2, s1}, cp.delimiter)

	if _, ok := cp.symbols[market]; ok {
		return market, nil
	}

	err := fmt.Errorf("%s does not support %s - %s market\n", cp.exchange, s1, s2)

	return "", err
}

func (cp SymbolMap) GetSymbolsForAsset(a1 string) []string {
	var matchedSymbols []string
	for symbol, assets := range cp.symbols {

		if assets[0] == a1 || assets[1] == a1 {

			matchedSymbols = append(matchedSymbols, symbol)
		}
	}

	return matchedSymbols
}

func (cp SymbolMap) BuildAssetMap() *AssetMap {

	var assets map[string][]string = map[string][]string{}

	for market, coins := range cp.symbols {
		for _, coin := range coins {
			if val, ok := assets[coin]; ok {
				assets[coin] = append(val, market)
				continue
			}
			assets[coin] = []string{market}
		}
	}

	return NewAssetMap(cp.exchange, assets)
}
