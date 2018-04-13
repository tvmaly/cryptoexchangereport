package exchangemarkets

import (
	"testing"
)

var TestMarkets map[string][2]string = map[string][2]string{
	"LTC-BTC": [2]string{"LTC", "BTC"},
	"ETH-BTC": [2]string{"ETH", "BTC"},
	"ENG-ETH": [2]string{"ENG", "ETH"},
}

var em ExchangeMarkets = New(TestMarkets, "-", true)

func TestPair(t *testing.T) {
	myPair, err := em.Pair("BTC", "LTC")

	if err != nil {
		t.Fatal(err)
	}

	myPair, err = em.Pair("btc", "ltc")

	if err != nil {
		t.Fatal(err)
	}

	myPair, err = em.Pair("BTC", "WAKA")

	if err == nil {
		t.Fatalf("BTC-WAKA does not exists, should have to output error, but instead gave %s", myPair)
	}
}

func TestGetPairsForCoin(t *testing.T) {
	pairs, err := em.GetPairsForCoin("btc")

	if err != nil {
		t.Fatal(err)
	}

	if len(pairs) != 2 {
		t.Fatalf("GetPairsForCoin() Error, %v", pairs)
	}

	pairs, err = em.GetPairsForCoin("ENG")

	if err != nil {
		t.Fatal(err)
	}

	if len(pairs) != 1 {
		t.Fatalf("GetPairsForCoin() Error, %v", pairs)
	}
}

func TestPairs(t *testing.T) {
	pairs := em.Pairs()
	if len(pairs) != 3 {
		t.Fatalf("Pairs() Error, %v", pairs)
	}
}

func TestCoins(t *testing.T) {
	coins := em.Coins()
	if len(coins) != 4 {
		t.Fatalf("Coins() Error, %v", coins)
	}
}

func TestPairExist(t *testing.T) {
	invalidPair := "BTC-WAKA"
	validPair := "LTC-BTC"

	if em.PairExist(invalidPair) {
		t.Fatalf("PairExist() Error, %s should not exist", invalidPair)
	}

	if !em.PairExist(validPair) {
		t.Fatalf("PairExist() Error, %s shoul exist", validPair)
	}
}

func TestCoinExist(t *testing.T) {
	invalidCoin := "WAKA"
	validCoin := "ENG"

	if em.CoinExist(invalidCoin) {
		t.Fatalf("CoinExist() Error, %s should not exist", invalidCoin)
	}

	if !em.CoinExist(validCoin) {
		t.Fatalf("CoinExist() Error, %s should exist", validCoin)
	}
}
