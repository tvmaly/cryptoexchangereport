package requestcounter

import (
	"time"
)

type Refresher interface {
	Refresh(int) int
}

type SimpleRefresher struct {
	interval int
	lastHit  time.Time
}

func NewSimpleRefresher(interval int) Refresher {
	return &SimpleRefresher{interval, time.Now()}
}

func (sr *SimpleRefresher) Refresh(counter int) int {
	t2 := time.Now()

	var countsToDecrement int = int(t2.Sub(sr.lastHit).Seconds() / float64(sr.interval))

	sr.lastHit = t2

	if countsToDecrement > counter {
		return 0
	} else {
		return counter - countsToDecrement
	}
}
