package bittrex

import (
    "testing"
    "fmt"
)


func Test (t *testing.T) {
    fmt.Println("Dummy")
}

/*
func TestOrderBook (t *testing.T) {
    adapter := New()
    fmt.Printf("%v\n",adapter.OrderBook("BTC-LTC"))
}
func TestTicker (t *testing.T) {
    adapter := New()
    adapter.Ticker("BTC-LTC")
}

func TestLastTrades (t *testing.T) {
    adapter := New()
    adapter.LastTrades("BTC-LTC")
}
*/

/*
func TestLimit (t *testing.T) {
    adapter := New()
    adapter.counter = 500

    adapter.LastTrades("BTC-LTC")
}
*/
