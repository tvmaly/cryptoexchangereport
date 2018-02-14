package wrapper

type ExchangeAdapter interface {
    Ticker() []byte
    OrderBook() []byte
    LastTrades() []byte
}
