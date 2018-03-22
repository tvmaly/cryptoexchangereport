package requestcounter

import (
	"testing"
	"time"
)

var RCounter AutoRefreshedRequestCounter

func init() {
	RCounter = New(3, NewSimpleRefresher(1))
}

func TestCounter(t *testing.T) {
	err := RCounter.SetCounter(5)

	if err == nil {
		t.Fatal("Should not allow to set counter above the limit! Counter: %d, Limit: %d", RCounter.GetCounter(), RCounter.GetLimit())
	}

	err = RCounter.Hit(1)

	if err != nil {
		t.Fatal(err)
	}

	err = RCounter.Hit(2)

	if err == nil {
		t.Fatal("Cannot hit 2, as we will reach limit")
	}

	if RCounter.GetCounter() != 1 {
		t.Fatal("We didn't hit 2, hence our counter stays the same")
	}

	RCounter.Hit(1)

	if RCounter.GetCounter() != 2 {
		t.Fatalf("We hitted 2 times by one, but counter is at %d", RCounter.GetCounter())
	}
}

func TestLimit(t *testing.T) {
	RCounter.SetLimit(10)

	if RCounter.GetLimit() != 10 {
		t.Fatalf("SetLimit does not set limit, limit is at %d", RCounter.GetCounter())
	}

}

func TestExpiration(t *testing.T) {
	RCounter.SetLimit(4)
	RCounter.SetCounter(0)

	RCounter.Hit(3)

	time.Sleep(4000 * time.Millisecond)

	if RCounter.GetCounter() != 0 {
		t.Fatalf("Refresh does not work! After 3 seconds counter is %d", RCounter.GetCounter())
	}
}
