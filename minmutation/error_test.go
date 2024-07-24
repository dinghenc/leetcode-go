package minmutation

import (
	"fmt"
	"testing"
)

type ProcessorError struct {
	Cause     error
	Retryable bool
}

func (err ProcessorError) Error() string {
	return fmt.Sprintf("ProcessorError (Retryable: %t): %v", err.Retryable, err.Cause)
}

func TestMainHHH(t *testing.T) {
	e := getError()
	if e != nil {
		t.Logf("%T, %v", e, e)
		e.Error()
	}
}

func getError() error {
	e := &ProcessorError{}
	e = nil
	return e
}
