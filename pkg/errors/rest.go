package errors

type RestError interface {
	Error() string
	Code() int
}
type restError struct {
	msg  string
	code int
}

func (f *restError) Error() string {
	return f.msg
}

func (f *restError) Code() int {
	return f.code
}
func NewRestError(code int, msg string) RestError {
	return &restError{
		msg:  msg,
		code: code,
	}
}
