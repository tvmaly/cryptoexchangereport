package bittrex

const ApiBase = "https://bittrex.com/api/v1.1" // for further refactoring from constants to a package

const OrderBookEndpoint = "/public/getorderbook?market=%s&type=both"
const LastTradesEndpoint = "/public/getmarkethistory?market=%s"
const TickerEndpoint = "/public/getticker?market=%s"
const SymbolsEndpoint = "/public/getmarkets"
