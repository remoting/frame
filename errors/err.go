package errors

import (
	"errors"
	pkgErrors "github.com/pkg/errors"
)

type RestError struct {
	error
	Code int
}

func NewRestError(code int, msg string) RestError {
	return RestError{
		error: New(msg),
		Code:  code,
	}
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}
func As(err, target error) bool {
	return errors.As(err, target)
}
func Unwrap(err error) error {
	return errors.Unwrap(err)
}

// ////////
func New(message string) error {
	return pkgErrors.New(message)
}

func Errorf(format string, args ...interface{}) error {
	return pkgErrors.Errorf(format, args)
}
func WithStack(err error) error {
	return pkgErrors.WithStack(err)
}
func Wrap(err error, message string) error {
	return pkgErrors.Wrap(err, message)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return pkgErrors.Wrapf(err, format, args)
}

func WithMessage(err error, message string) error {
	return pkgErrors.WithMessage(err, message)
}

func WithMessagef(err error, format string, args ...interface{}) error {
	return pkgErrors.WithMessagef(err, format, args)
}
func Cause(err error) error {
	return pkgErrors.Cause(err)
}
