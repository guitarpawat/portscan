package model

type Error struct {
	ErrorMsg string `json:"msg"`
}

func (e Error) Error() string {
	return e.ErrorMsg
}
