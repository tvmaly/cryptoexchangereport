package assetsenumfactory

import (
	"testing"
)

func TestSuccessfulFactory(t *testing.T) {
	binanceAssetsEnum := NewExchangeAssetsEnum("binance")

	if binanceAssetsEnum == nil {
		t.Fatal("Binance is listed, so should be created!")
	}
}

func TestFailingFactory(t *testing.T) {
	kavabangaAssetsEnum := NewExchangeAssetsEnum("kavabanga")

	if kavabangaAssetsEnum != nil {
		t.Fatalf("There are no such exchange as Kavabanga registered in our system")
	}
}
