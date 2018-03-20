package exchangedata

import (
	"errors"
)

type LimitCounter interface {
	Limit() int
	Counter() int
	Hit(int) error
	SetLimit(int)
	SetCounter(int) error
}

type Limit struct {
	limit   int
	counter int
}

func New(limit int) LimitCounter {
	return &Limit{limit, 0}
}

func (l *Limit) Limit() int {
	return l.limit
}

func (l *Limit) Counter() int {
	return l.counter
}

func (l *Limit) SetLimit(newLimit int) {
	l.limit = newLimit
}

func (l *Limit) SetCounter(newCounter int) error {

	if l.Limit() < newCounter {
		err := errors.New("Counter cannot be bigger then a limit")
		return err

	} else {
		l.counter = newCounter
		return nil
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
