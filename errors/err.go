package errors

import "errors"

type RestError struct {
	error
	Code int
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}
func NewRestError(code int, msg string) RestError {
	return RestError{
		error: errors.New(msg),
		Code:  code,
	}
}
