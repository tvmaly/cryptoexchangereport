package quotes

// ExchangeQuote this interface is for exchange quotes that can
// represent themselves in a generic normalized form
type ExchangeQuote interface {
	GenericQuote() *GenericQuote
}

// Quote interface
type Quote interface {
	MidPoint() float64
	SpreadPercentage() float64
	Weighted() bool
}
