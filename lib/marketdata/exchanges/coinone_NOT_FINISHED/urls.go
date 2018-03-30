package coinone

const ApiBase = "https://api.coinone.co.kr" // for further refactoring from constants to a package

const OrderBookEndpoint = "/orderbook?currency=%s&format=json"
const LastTradesEndpoint = "/trades?currency=%s&format=json"
const TickerEndpoint = "/ticker?currency=%s&format=json"

// In future, will be produced by parsing api documentation, if no better source
// will be found
const Symbols []string = []string{
	"btc",
	"bch",
	"eth",
	"etc",
	"xrp",
	"qtum",
	"ltc",
	"iota",
	"btg",
}
