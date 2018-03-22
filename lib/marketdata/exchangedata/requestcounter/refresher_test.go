package requestcounter

import (
	"fmt"
	"testing"
	"time"
)

func TestRefresher(t *testing.T) {
	var refresher Refresher
	var counter int = 4

	fmt.Printf("Initial counter value is %d\n", counter)

	refresher = NewSimpleRefresher(1)

	counter = refresher.Refresh(counter)
	VerifyCounter(t, 0, 4, counter)

	time.Sleep(3 * time.Second)
	counter = refresher.Refresh(counter)
	VerifyCounter(t, 3, 1, counter)

	time.Sleep(2 * time.Second)
	counter = refresher.Refresh(counter)
	VerifyCounter(t, 2, 0, counter)
}

func VerifyCounter(t *testing.T, waited, expected, obtained int) {
	if obtained != expected {
		t.Fatal(fmt.Errorf("We waited %d seconds. Counter should be at %d, but we have %d", waited, expected, obtained))
	}
}
