package bitz

const ApiBase = "https://www.bit-z.com/api_v1" // for further refactoring from constants to a package

const OrderBookEndpoint = "/depth?coin=%s"
const LastTradesEndpoint = "/orders?coin=%s"
const TickerEndpoint = "/ticker?coin=%s"
const SymbolsEndpoint = "/tickerall"
