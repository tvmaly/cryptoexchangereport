//requestcounter
//Exchanges impose limits on api request rate. Exceeding this limits could result in //a temporary ban from api requests, which is non-acceptable situation for our service.
//In order to make sure that we do not exceed api requests rate limit, we need a smart counter.
package requestcounter

import (
	"errors"
)

//Simple RequestCounter interace
type RequestCounter interface {
	IsLimit(int) bool
	Hit(int) error

	SetLimit(int)
	SetCounter(int) error
	GetLimit() int
	GetCounter() int
}

//RequestCounter with Refresher logic
type AutoRefreshedRequestCounter interface {
	RequestCounter
	Refresh()
}

//Simple struct that implements AutoRefreshedRequestCounter interface
type ExchangeAPIRequestCounter struct {
	refresher Refresher
	limit     int
	counter   int
}

func New(limit int, refresher Refresher) AutoRefreshedRequestCounter {
	return &ExchangeAPIRequestCounter{refresher, limit, 0}
}

//Limit Getter
func (earc *ExchangeAPIRequestCounter) GetLimit() int { return earc.limit }

//Limit Setter
func (earc *ExchangeAPIRequestCounter) SetLimit(newLimit int) { earc.limit = newLimit }

//Counter Getter
//Always refreshes counter before returning its value
func (earc *ExchangeAPIRequestCounter) GetCounter() int {
	earc.Refresh()
	return earc.counter
}

//Counter Setter
//Will return an error if specified counter > current limit
func (earc *ExchangeAPIRequestCounter) SetCounter(newCounter int) error {
	if earc.GetLimit() < newCounter {
		err := errors.New("Counter cannot be bigger then a limit")
		return err

	} else {
		earc.counter = newCounter
		return nil
	}
}

//Refreshes our counter
//passes counter to our refresher strategy and assings received value to our counter
func (earc *ExchangeAPIRequestCounter) Refresh() {
	earc.SetCounter(earc.refresher.Refresh(earc.counter))
}

//Increments counter on a specified amount
//Will return an error if counter value + specified amount > limit
func (earc *ExchangeAPIRequestCounter) Hit(pointsToHit int) error {
	if earc.GetCounter()+pointsToHit < earc.GetLimit() {
		earc.SetCounter(earc.counter + pointsToHit)
		return nil

	} else {

		err := errors.New("Hit the limit!")
		return err
	}
}

//Checks if request will hit the limit
//Returns true if yes, false otherwise
func (earc *ExchangeAPIRequestCounter) IsLimit(pointsToHit int) bool {
	if pointsToHit+earc.GetCounter() < earc.GetLimit() {
		return false
	} else {
		return true
	}
}
