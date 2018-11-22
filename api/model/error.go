package model

import "encoding/json"

// Error struct that can marshall to json and implements with error interface
type Error struct {
	ErrorMsg string `json:"msg"`
}

// Error returns the error string to satisfy with error interface
func (e *Error) Error() string {
	return e.ErrorMsg
}

// Marshal returns the json version of the error
func (e *Error) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

// MakeError makes error with json support from error
func MakeError(err error) *Error {
	return &Error{err.Error()}
}

// MakeErrorString makes error with json support from string
func MakeErrorString(err string) *Error {
	return &Error{err}
}
