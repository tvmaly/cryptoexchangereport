// module to control both cryptocurrency coins and pairs
package assets

import (
	"fmt"
	"strings"
)

type CryptoCoinChecker interface {
	Coins() []string
	GetPairsForCoin(string) ([]string, error)
	CoinExist(string) bool
}

type CryptoPairChecker interface {
	Pairs() []string
	Pair(string, string) (string, error)
	PairExist(string) bool
}

type ExchangeAssetsChecker interface {
	CryptoCoinChecker
	CryptoPairChecker
}

// Underlying struct
type ExchangeAssetsLookup struct {
	coinToPairs map[string][]string
	pairToCoins map[string][2]string
	delimiter   string
	adjustCase  func(...*string)
}

// Constructor.
// Arguments:
// assets <map[string][2]string>. Each key is a coins pair and value is a coins that makes up this pair. Will build underlying data structure from it
// delimiter <string>. To Join coins for a Pair() method.
// caseFlag <bool>: true - uppercase, false - lowercase. To build caseAdjuster closure
func New(assets map[string][2]string, delimiter string, caseFlag bool) ExchangeAssetsChecker {
	coinToPairs := map[string][]string{}
	pairToCoins := map[string][2]string{}
	var caseAdjuster func(...*string)

	if caseFlag {
		caseAdjuster = WrapWithVariadicArgs(strings.ToUpper)
	} else {
		caseAdjuster = WrapWithVariadicArgs(strings.ToLower)
	}

	for pair, coins := range assets {
		caseAdjuster(&pair, &coins[0], &coins[1])

		pairToCoins[pair] = coins

		for _, coin := range coins {

			if val, ok := coinToPairs[coin]; ok {
				coinToPairs[coin] = append(val, pair)
			} else {
				coinToPairs[coin] = []string{pair}
			}
		}
	}

	return &ExchangeAssetsLookup{
		coinToPairs,
		pairToCoins,
		delimiter,
		caseAdjuster,
	}
}

// Wrapper with Variadic *string arguments around strings.ToUpper or strings.ToLower
// Some crypto exchanges restrict coins and pairs only to upper or lower case, in this case we will normalize case for all entries and output
func WrapWithVariadicArgs(adjustCase func(string) string) func(...*string) {
	return func(args ...*string) {
		for _, arg := range args {
			*arg = adjustCase(*arg)
		}
	}
}

// Returns all the current pairs
func (eal ExchangeAssetsLookup) Pairs() []string {
	pairs := make([]string, 0, len(eal.pairToCoins))
	for pair := range eal.pairToCoins {
		pairs = append(pairs, pair)
	}
	return pairs
}

// Given two coins will return a pair if it exists, otherwise - empty string and error
func (eal ExchangeAssetsLookup) Pair(coin1, coin2 string) (string, error) {
	eal.adjustCase(&coin1, &coin2)

	pair := strings.Join([]string{coin1, coin2}, eal.delimiter)

	if eal.PairExist(pair) {
		return pair, nil
	}

	pair = strings.Join([]string{coin2, coin1}, eal.delimiter)

	if eal.PairExist(pair) {
		return pair, nil
	}

	err := fmt.Errorf("Could not find pair for %s and %s\n", coin1, coin2)

	return "", err
}

// Returns true if pair exists, otherwise false
func (eal ExchangeAssetsLookup) PairExist(pair string) bool {
	eal.adjustCase(&pair)
	if _, ok := eal.pairToCoins[pair]; ok {
		return true
	} else {
		return false
	}
}

// Returns all the current coins
func (eal ExchangeAssetsLookup) Coins() []string {
	coins := make([]string, 0, len(eal.coinToPairs))
	for coin := range eal.coinToPairs {
		coins = append(coins, coin)
	}
	return coins
}

// Returns all Pairs for a given coin
func (eal ExchangeAssetsLookup) GetPairsForCoin(coin string) ([]string, error) {
	eal.adjustCase(&coin)

	var pairs []string

	if !eal.CoinExist(coin) {
		err := fmt.Errorf("%s is not found", coin)
		return pairs, err

	} else {
		pairs = eal.coinToPairs[coin]

		if len(pairs) == 0 {
			err := fmt.Errorf("Zero symbols for existing coin %s\n", coin)
			return pairs, err
		}

		return pairs, nil
	}
}

// Returns true if coin exists, otherwise false
func (eal ExchangeAssetsLookup) CoinExist(coin string) bool {
	eal.adjustCase(&coin)
	if _, ok := eal.coinToPairs[coin]; ok {
		return true
	} else {
		return false
	}
}
