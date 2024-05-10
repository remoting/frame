package errors

import (
	"fmt"
	"github.com/remoting/frame/reflect"
)

func TryCatch(tryFunc func() error, catchFunc func(err error)) {
	catchFunc(TryError(tryFunc))
}
func TryError(tryFunc func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	return tryFunc()
}
func TryAny[T any](tryFunc func() (T, error)) (t T, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	return tryFunc()
}
func TryThrow(err error) {
	if !reflect.IsNil(err) {
		panic(err)
	}
}
