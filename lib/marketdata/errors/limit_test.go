package errors

import (
	"testing"
)

func TestLimitError(t *testing.T) {
	err := FunctionThatReturnsError()

	if err.ExchangeName() != "Bittrex" {
		t.Fatalf("Exchange name should be Bittrex, we have: %s", err.ExchangeName())
	}

	if err.Error() != "Bittrex: Limit Hit" {
		t.Fatalf("error message should be \"Bittrex: Limit Hit\", we have: %s", err.Error())
	}

	if _, ok := err.(LimitError); !ok {
		t.Fatalf("incorrect error type, must be of LimitError type")
	}
}

func FunctionThatReturnsError() LimitError {
	return NewLimitError("Bittrex")
}
