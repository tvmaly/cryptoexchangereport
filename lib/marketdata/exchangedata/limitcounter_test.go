package exchangedata

import (
	"testing"
)

var TLimit LimitCounter

func init() {
	TLimit = New(3)
}

func TestCounter(t *testing.T) {
	err := TLimit.SetCounter(5)

	if err == nil {
		t.Fatal("Should not allow to set counter above the limit! Counter: %d, Limit: %d", TLimit.Counter(), TLimit.Limit())
	}

	err = TLimit.Hit(1)

	if err != nil {
		t.Fatal(err)
	}

	err = TLimit.Hit(2)

	if err == nil {
		t.Fatal("Cannot hit 2, as we will reach limit")
	}

	if TLimit.Counter() != 1 {
		t.Fatal("We didn't hit 2, hence our counter stays the same")
	}

	TLimit.Hit(1)

	if TLimit.Counter() != 2 {
		t.Fatalf("We hitted 2 times by one, but counter is at %d", TLimit.Counter())
	}
}

func TestLimit(t *testing.T) {
	TLimit.SetLimit(10)

	if TLimit.Limit() != 10 {
		t.Fatalf("SetLimit does not set limit, limit is at %d", TLimit.Counter())
	}
}

/*
type LimitCounter interface {
	UpperLimit() int
	Counter() int
	Hit(int) error
	SetLimit(int)
	SetCounter(int)
}
*/
