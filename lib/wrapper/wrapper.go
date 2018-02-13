package wrapper

type ExchangeAdapter interface {
    AllMarkets() []byte
    Ticker() []byte
    OrderBook() []byte
    LastTrades() []byte
}
