package book

import (
	"github.com/tvmaly/cryptoexchangereport/lib/parsing/quotes"
)

type DepthOfBook interface {
	BBO() *quotes.GenericQuote
}

func (b *CoinoneBook) BBO() *GenericQuote {

	var lowest_ask float64
	var lowest_ask_quantity float64

	var highest_bid float64
	var highest_bid_quantity float64

	for _, ask := range b.Ask {

		if lowest_ask == 0 {
			lowest_ast = ask.Price
			lowest_ask_quantity = ask.Qty
		} else {

			if ask.Price < lowest_ask {
				lowest_ast = ask.Price
				lowest_ask_quantity = ask.Qty
			}
		}

	}

	for _, bid := range b.Bid {

		if highest_bid == 0 {
			highest_bid = bid.Price
			highest_bid_quantity = bid.Qty
		} else {

			if bid.Price > lowest_ask {
				highest_bid = bid.Price
				highest_bid_quantity = bid.Qty
			}
		}

	}

	return &GenericQuote{
		Exchange:    constants.COINONE,
		TimeStamp:   b.Timestamp,
		BidPrice:    highest_bid,
		BidQuantity: highest_bid_quantity,
		AskPrice:    lowest_ask,
		AskQuantity: lowest_ask_quantity,
	}
}
