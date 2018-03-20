package exchangedata

import (
	"errors"
	"time"
)

type LimitCounter interface {
	Refresh(t time.Time)
	Limit() int
	Counter() int
	Hit(int) error
	SetLimit(int)
	SetCounter(int) error
}

type Limit struct {
	limit            int
	counter          int
	getExpiredCounts func(time.Time, time.Time) int
	lastReset        time.Time
}

func New(limit int, getExpiredCounts func(time.Time, time.Time) int) LimitCounter {
	lim := &Limit{limit, 0, getExpiredCounts, time.Now()}

	ticker := time.NewTicker(350 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			lim.Refresh(t)
		}
	}()

	return lim
}

func (l *Limit) Limit() int { return l.limit }

func (l *Limit) Counter() int { return l.counter }

func (l *Limit) SetLimit(newLimit int) { l.limit = newLimit }

func (l *Limit) SetCounter(newCounter int) error {
	if l.Limit() < newCounter {
		err := errors.New("Counter cannot be bigger then a limit")
		return err

	} else {
		l.counter = newCounter
		return nil
	}
}

func (l *Limit) Refresh(t time.Time) {
	expiredCounts := l.getExpiredCounts(l.lastReset, t)

	if expiredCounts != 0 {
		l.Expired(expiredCounts)
		l.lastReset = time.Now()
	}
}

func (l *Limit) Hit(pointsToIncrement int) error {
	if l.Counter()+pointsToIncrement < l.Limit() {
		l.counter += pointsToIncrement
		return nil

	} else {

		err := errors.New("Hit the limit!")
		return err
	}
}

func (l *Limit) Expired(pointsToDecrement int) {
	if pointsToDecrement > l.Counter() {
		l.SetCounter(0)
		return

	} else {
		l.SetCounter(l.Counter() - pointsToDecrement)
		return
	}
}

func MockgetExpiredCounts() func(time.Time, time.Time) int {
	return func(t1, t2 time.Time) int {
		var seconds int = int(t2.Sub(t1).Seconds())
		return seconds
	}
}
