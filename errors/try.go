package errors

import (
	"fmt"
	"github.com/remoting/frame/reflect"
)

func Recover(err *error) {
	if p := recover(); p != nil {
		*err = fmt.Errorf("%v", p)
	} else {
		*err = nil
	}
}
func TryCatch(tryFunc func() error, catchFunc func(err error)) {
	catchFunc(TryError(tryFunc))
}
func TryError(tryFunc func() error) (err error) {
	defer Recover(&err)
	return tryFunc()
}
func TryAny[T any](tryFunc func() (T, error)) (t T, err error) {
	defer Recover(&err)
	return tryFunc()
}
func TryThrow(err error) {
	if !reflect.IsNil(err) {
		panic(err)
	}
}
