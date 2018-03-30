package errors

import "fmt"

type LimitError interface {
	Error() string
	ExchangeName() string
}

func NewLimitError(name string) LimitError {
	return &limitError{name}
}

type limitError struct {
	name string
}

func (e *limitError) Error() string {
	return fmt.Sprintf("%s: Limit Hit", e.name)
}

func (e *limitError) ExchangeName() string {
	return e.name
}
