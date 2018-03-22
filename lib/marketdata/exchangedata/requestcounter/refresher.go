//Exchanges API counter will decrement over time
//every exchange has its own logic for decrementing api requests counter
//create struct with your exchange-specific api requests counter decrementing logic
//that implements Refresher interface and pass it to the constructor of RequestCounter struct
//Refresh method takes in value of your current counter and returns decremented counter value
package requestcounter

import (
	"time"
)

type Refresher interface {
	Refresh(int) int
}

//SimpleRefresher struct
//pass in an interval - amount of seconds after which counter will be decremented by 1
//if interval > 100 - will be treated as milliseconds value and convert it to seconds, so 200 will be passed as 0.2
type SimpleRefresher struct {
	interval float64
	lastHit  time.Time
}

func NewSimpleRefresher(interval float64) Refresher {
	if interval > 100 {
		interval = interval / 1000
	}

	return &SimpleRefresher{interval, time.Now()}
}

//Core logic.
//Takes in current counter, Calculates amount of points to decrement by calculating
//amount of intervals passed since last request and returns refreshed counter
func (sr *SimpleRefresher) Refresh(counter int) int {
	t2 := time.Now()

	var countsToDecrement int = int(t2.Sub(sr.lastHit).Seconds() / sr.interval)

	sr.lastHit = t2

	if countsToDecrement > counter {
		return 0
	} else {
		return counter - countsToDecrement
	}
}
