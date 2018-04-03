package assets

import (
	"testing"
)

var TestAssets map[string][2]string = map[string][2]string{
	"LTC-BTC": [2]string{"LTC", "BTC"},
	"ETH-BTC": [2]string{"ETH", "BTC"},
	"ENG-ETH": [2]string{"ENG", "ETH"},
}

var eae ExchangeAssetsChecker = New(TestAssets, "-", true)

func TestPair(t *testing.T) {
	myPair, err := eae.Pair("BTC", "LTC")

	if err != nil {
		t.Fatal(err)
	}

	myPair, err = eae.Pair("btc", "ltc")

	if err != nil {
		t.Fatal(err)
	}

	myPair, err = eae.Pair("BTC", "WAKA")

	if err == nil {
		t.Fatalf("BTC-WAKA does not exists, should have to output error, but instead gave %s", myPair)
	}
}

func TestGetPairsForCoin(t *testing.T) {
	pairs, err := eae.GetPairsForCoin("btc")

	if err != nil {
		t.Fatal(err)
	}

	if len(pairs) != 2 {
		t.Fatalf("GetPairsForCoin() Error, %v", pairs)
	}

	pairs, err = eae.GetPairsForCoin("ENG")

	if err != nil {
		t.Fatal(err)
	}

	if len(pairs) != 1 {
		t.Fatalf("GetPairsForCoin() Error, %v", pairs)
	}
}

func TestPairs(t *testing.T) {
	pairs := eae.Pairs()
	if len(pairs) != 3 {
		t.Fatalf("Pairs() Error, %v", pairs)
	}
}

func TestCoins(t *testing.T) {
	coins := eae.Coins()
	if len(coins) != 4 {
		t.Fatalf("Coins() Error, %v", coins)
	}
}

func TestPairExist(t *testing.T) {
	invalidPair := "BTC-WAKA"
	validPair := "LTC-BTC"

	if eae.PairExist(invalidPair) {
		t.Fatalf("PairExist() Error, %s should not exist", invalidPair)
	}

	if !eae.PairExist(validPair) {
		t.Fatalf("PairExist() Error, %s shoul exist", validPair)
	}
}

func TestCoinExist(t *testing.T) {
	invalidCoin := "WAKA"
	validCoin := "ENG"

	if eae.CoinExist(invalidCoin) {
		t.Fatalf("CoinExist() Error, %s should not exist", invalidCoin)
	}

	if !eae.CoinExist(validCoin) {
		t.Fatalf("CoinExist() Error, %s should exist", validCoin)
	}
}
