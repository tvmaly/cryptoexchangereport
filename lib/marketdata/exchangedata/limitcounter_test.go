package exchangedata

import (
	"testing"
	"time"
)

var TLimit LimitCounter

func init() {
	TLimit = New(3, MockgetExpiredCounts())
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

func TestExpiration(t *testing.T) {
	TLimit.SetLimit(4)
	TLimit.SetCounter(0)

	TLimit.Hit(3)

	time.Sleep(4000 * time.Millisecond)

	if TLimit.Counter() != 0 {
		t.Fatalf("Refresh does not work! After 3 seconds counter is %d", TLimit.Counter())
	}
}
