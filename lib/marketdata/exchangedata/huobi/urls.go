package huobi

const ApiBase = "https://api.huobi.pro"

const OrderBookEndpoint = "/market/depth?symbol=%s&type=step0"
const LastTradesEndpoint = "/market/trade?symbol=%s"
const TickerEndpoint = "/market/detail/merged?symbol=%s"
const SymbolsEndpoint = "/v1/common/symbols"
const CoinsEndpoint = "/v1/common/currencys"
