package errors

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
