package bitstamp

const ApiBase = "https://www.bitstamp.net/api/v2" // for further refactoring from constants to a package

const OrderBookEndpoint = "/transactions/%s/"
const LastTradesEndpoint = "/order_book/%s/"
const TickerEndpoint = "/ticker/%s/"
const SymbolsEndpoint = "/trading-pairs-info/"
