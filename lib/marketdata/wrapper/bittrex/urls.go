package bittrex

const ApiBase = "https://bittrex.com/api/v1.1/" // for further refactoring from constants to a package

const OrderBookUrl = "https://bittrex.com/api/v1.1/public/getorderbook?market=%s&type=both"
const LastTradesUrl = "https://bittrex.com/api/v1.1/public/getmarkethistory?market=%s"
const TickerUrl = "https://bittrex.com/api/v1.1/public/getticker?market=%s"

const SymbolsUrl = "https://bittrex.com/api/v1.1/public/getmarkets"
