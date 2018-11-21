package model

// Error struct that can marshall to json and implements with error interface
type Error struct {
	ErrorMsg string `json:"msg"`
}

// Error returns the error string to satisfy with error interface
func (e Error) Error() string {
	return e.ErrorMsg
}
