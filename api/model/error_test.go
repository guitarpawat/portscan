package model

import (
	"testing"
)

func makeError() interface{} {
	return Error{"test"}
}

func TestError_Error(t *testing.T) {
	err := makeError()
	switch e := err.(type) {
	case error:
		// Expected
		if e.Error() != "test" {
			t.Errorf("expected error msg: %s, but get: %s", "test", e.Error())
		}
	default:
		t.Errorf("expected interface: error, but get: %T", err)
	}
}
