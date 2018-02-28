package gdax

const ApiBase = "https://api.gdax.com" // for further refactoring from constants to a package

const OrderBookEndpoint = "/products/%s/book?level=3"
const LastTradesEndpoint = "/products/%s/trades"
const TickerEndpoint = "/products/%s/ticker"
const SymbolsEndpoint = "/products"
const CoinsEndpoint = "/currencies"
