package requestcounter

import (
	"errors"
)

type RequestCounter interface {
	IsLimit(int) bool
	Hit(int) error

	SetLimit(int)
	SetCounter(int) error
	GetLimit() int
	GetCounter() int
}

type AutoRefreshedRequestCounter interface {
	RequestCounter
	Refresh()
}

type ExchangeAPIRequestCounter struct {
	refresher Refresher
	limit     int
	counter   int
}

func New(limit int, refresher Refresher) AutoRefreshedRequestCounter {
	return &ExchangeAPIRequestCounter{refresher, limit, 0}
}

func (earc *ExchangeAPIRequestCounter) GetLimit() int { return earc.limit }

func (earc *ExchangeAPIRequestCounter) GetCounter() int {
	earc.Refresh()
	return earc.counter
}

func (earc *ExchangeAPIRequestCounter) SetLimit(newLimit int) { earc.limit = newLimit }

func (earc *ExchangeAPIRequestCounter) SetCounter(newCounter int) error {
	if earc.GetLimit() < newCounter {
		err := errors.New("Counter cannot be bigger then a limit")
		return err

	} else {
		earc.counter = newCounter
		return nil
	}
}

func (earc *ExchangeAPIRequestCounter) Refresh() {
	earc.SetCounter(earc.refresher.Refresh(earc.counter))
}

func (earc *ExchangeAPIRequestCounter) Hit(pointsToHit int) error {
	if earc.GetCounter()+pointsToHit < earc.GetLimit() {
		earc.SetCounter(earc.counter + pointsToHit)
		return nil

	} else {

		err := errors.New("Hit the limit!")
		return err
	}
}

func (earc *ExchangeAPIRequestCounter) IsLimit(pointsToHit int) bool {
	if pointsToHit+earc.GetCounter() < earc.GetLimit() {
		return true
	} else {
		return false
	}
}
