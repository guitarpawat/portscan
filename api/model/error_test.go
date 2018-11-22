package model

import (
	"encoding/json"
	"errors"
	"testing"
)

func makeError() Json {
	return &Error{"test"}
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

func TestError_Marshal(t *testing.T) {
	err := makeError()
	b, _ := err.Marshal()
	expected, _ := json.Marshal(err)
	if string(expected) != string(b) {
		t.Errorf("expected json: %s, but get: %s", string(expected), string(b))
	}
}

func TestMakeError(t *testing.T) {
	expected := "maketest"
	err := errors.New(expected)
	jsonErr := MakeError(err)

	if err.Error() != jsonErr.Error() {
		t.Errorf("expected error: %s, but get: %s", err, jsonErr)
	}
}

func TestMakeErrorString(t *testing.T) {
	expected := "maketest"
	jsonErr := MakeErrorString(expected)

	if expected != jsonErr.Error() {
		t.Errorf("expected error: %s, but get: %s", expected, jsonErr)
	}
}
