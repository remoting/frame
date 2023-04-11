package errors

import "errors"

type RestError struct {
	error
	Code int
}

func New(code int, msg string) RestError {
	return RestError{
		error: errors.New(msg),
		Code:  code,
	}
}
