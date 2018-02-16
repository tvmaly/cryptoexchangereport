package symbolmap

import (
    "strings"
    "fmt"
)

type SymbolMap struct {
    exchange string
    symbols map[string]bool
    delimiter string
}

func New(exchange, delimiter string, symbols map[string]bool) *SymbolMap {

    var cp SymbolMap = SymbolMap{
        exchange: exchange,
        symbols: symbols,
        delimiter: delimiter,
    }

    return &cp
}

func (cp SymbolMap) GetSymbol (s1, s2 string) (string, error) {

    s1 = strings.ToUpper(s1)
    s2 = strings.ToUpper(s2)

    market := strings.Join([]string{s1, s2}, cp.delimiter)

    if cp.symbols[market] {
        return market, nil
    }

    market = strings.Join([]string{s2, s1}, cp.delimiter)

    if cp.symbols[market] {
        return market, nil
    }

    err := fmt.Errorf("%s does not support %s - %s market\n", cp.exchange, s1, s2)

    return "", err
}
