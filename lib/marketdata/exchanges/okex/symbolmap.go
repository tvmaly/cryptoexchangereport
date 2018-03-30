package okex

import (
	"cryptoexchangereport/marketdata/assets"
	"log"
	"strings"
)

func NewSymbolMap() *assets.SymbolMap {
	symbols, err := getSymbols()

	if err != nil {
		log.Fatalf("Error while getting symbols from okex: %v", err)
	}

	return assets.NewSymbolMap("okex", "_", false, symbols)
}

func getSymbols() (map[string][2]string, error) {
	var symbolsMap map[string][2]string = map[string][2]string{}

	symbols := []string{
		"ltc_btc",
		"eth_btc",
		"etc_btc",
		"bch_btc",
		"btc_usdt",
		"eth_usdt",
		"ltc_usdt",
		"etc_usdt",
		"bch_usdt",
		"etc_eth",
		"bt1_btc",
		"bt2_btc",
		"btg_btc",
		"qtum_btc",
		"hsr_btc",
		"neo_btc",
		"gas_btc",
		"qtum_usdt",
		"hsr_usdt",
		"neo_usdt",
		"gas_usdt",
	}

	for _, symbol := range symbols {

		coins := strings.Split(symbol, "_")

		symbolsMap[symbol] = [2]string{
			coins[0],
			coins[1],
		}
	}

	return symbolsMap, nil
}
