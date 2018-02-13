package quotes

type ExchangeQuote interface {
	GenericQuote() *GenericQuote
}

type Quote interface {
	MidPoint() float64
	Weighted() bool
}
