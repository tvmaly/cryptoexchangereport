package exchangemarkets

type ExchangeMarketsFactory map[string]func() ExchangeMarkets

func (emf *ExchangeMarketsFactory) Register(exchange string, getExchangeMarkets func() ExchangeMarkets) {
	(*emf)[exchange] = getExchangeMarkets
}

func (emf *ExchangeMarketsFactory) Delete(exchange string) {
	delete(*emf, exchange)
}

func (emf *ExchangeMarketsFactory) GetExchangeMarkets(exchange string) ExchangeMarkets {
	if getExchangeMarkets, ok := (*emf)[exchange]; ok {
		return getExchangeMarkets()
	} else {
		return nil
	}
}
