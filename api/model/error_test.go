package model

import (
	"testing"
)

func makeError() interface{} {
	return Error{"test"}
}

func TestError_Error(t *testing.T) {
	err := makeError()
	switch err.(type) {
	case error:
		// Expected
	default:
		t.Errorf("expected interface: error, but get: %T", err)
	}
}
