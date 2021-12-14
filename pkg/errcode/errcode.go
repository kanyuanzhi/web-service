package errcode

import "github.com/kanyuanzhi/web-service/global"

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var codes = map[int]string{}

func NewError(code int, message string) *Error {
	if _, ok := codes[code]; ok {
		global.Log.Warnf("code=%d repeats, please use another one", code)
		return nil
	}
	codes[code] = message
	return &Error{
		Code:    code,
		Message: message,
	}
}
