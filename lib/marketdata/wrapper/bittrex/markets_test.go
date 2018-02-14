package bittrex

import (
    "testing"
    "fmt"
)

func TestGetMarkets(t *testing.T) {

    markets, err := getMarkets(MarketsUrl)

    if err != nil {
        t.Error(fmt.Sprintf("getMarkets method failed: %v", err))
    }

    if !markets["BTC-ETH"] {
        t.Error("BTC-ETH should be a valid market")
    }
}
