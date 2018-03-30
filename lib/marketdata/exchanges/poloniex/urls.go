package poloniex

const ApiBase = "https://poloniex.com" // for further refactoring from constants to a package

const OrderBookEndpoint = "/public?command=returnOrderBook&currencyPair=%s"
const LastTradesEndpoint = "/public?command=returnTradeHistory&currencyPair=%s"
const TickerEndpoint = "/public?command=returnTicker"
const SymbolsEndpoint = "/public?command=returnTicker"
const CoinEndpoint = "/public?command=returnCurrencies"
