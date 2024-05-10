package errors

import "errors"

type RestError struct {
	error
	Code int
}

func NewRestError(code int, msg string) RestError {
	return RestError{
		error: errors.New(msg),
		Code:  code,
	}
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}
func New(text string) error {
	return errors.New(text)
}
func As(err, target error) bool {
	return errors.As(err, target)
}
func Unwrap(err error) error {
	return errors.Unwrap(err)
}
