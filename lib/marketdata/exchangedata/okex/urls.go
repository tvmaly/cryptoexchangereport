package okex

const ApiBase = "https://www.okex.com/api/v1" // for further refactoring from constants to a package

const OrderBookEndpoint = "/depth.do?symbol=%s"
const LastTradesEndpoint = "/trades.do?symbol=%s"
const TickerEndpoint = "/ticker.do?symbol=%s"
