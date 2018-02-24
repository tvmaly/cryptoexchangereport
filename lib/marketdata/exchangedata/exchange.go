package exchangedata

type Exchange interface {
	Name() string
	OrderBook(string, string) ([]byte, error)
	Ticker(string, string) ([]byte, error)
	LastTrades(string, string) ([]byte, error)
	GetSymbol(string, string) (string, error)
	SendRequest(string, string) (Response, error)
}
