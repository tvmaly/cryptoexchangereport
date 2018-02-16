package quotes

// GenericQuote this represents the basic set of fields that
// we wish to assume a quote has
// the one exception is the quantity fields
// these are not always present and this will affect calculations
// in respect to them being weighted by quantity values
// this is also used by Depth of Book to return a BBO value as a GenericQuote
// note, not all exchanges have a proper quote from their ticker api
// coinone is one such exchange, for this particular exchange use the BBO()
// method on its depth of book
type GenericQuote struct {
	Exchange          string  `json:exchange`
	Symbol            string  `json:symbol`
	Timestamp         float64 `json:timestamp,string`
	BidPrice          float64 `json:"bidprice,string"`
	BidQuantity       float64 `json:"bidquantity,string"`
	AskPrice          float64 `json:"askprice,string"`
	AskQuantity       float64 `json:"askquantity,string"`
	SentTimestamp     float64 `json:"senttimestamp,string"`
	ReceivedTimestamp float64 `json:"recievedtimestamp,string"`
}

// MidPoint calculate the mid point price of a quote
// if quantity is available, it will be weighted
// otherwise it will be a simple midpoint
func (q *GenericQuote) MidPoint() float64 {

	if q.Weighted() {
		return (q.BidPrice*q.BidQuantity + q.AskPrice*q.AskQuantity) / (q.BidQuantity + q.AskQuantity)
	} else {
		return (q.BidPrice + q.AskPrice) / 2
	}
}

// Weighted return true if quantity is set otherwise false
func (q *GenericQuote) Weighted() bool {
	return q.BidQuantity != 0 && q.AskQuantity != 0
}

// SpreadPercentage return the spread percentage of the quote
// to measure liquidity and true cost to trade
// formula ( ask - bid ) / ask * 100
func (q *GenericQuote) SpreadPercentage() float64 {

	if q.AskPrice == 0 {

		if q.BidPrice == 0 {
			return 0
		}

		return 1.0
	}

	return ((q.AskPrice - q.BidPrice) / q.AskPrice) * 100.0
}

// Latency is a rough estimation of round trip request time to exchange divided by 2
// returned value is in seconds
func (q *GenericQuote) Latency() float64 {

	if q.SentTimestamp == 0 || q.ReceivedTimestamp == 0 {
		return 0
	}

	return (q.ReceivedTimestamp - q.SentTimestamp) / 2.0

}

func (q *GenericQuote) EstimatedTimestamp() float64 {

	if q.Latency() == 0 {
		return 0
	}

	return q.SentTimestamp + q.Latency()

}
