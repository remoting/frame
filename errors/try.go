package errors

import "fmt"

func TryError(tryFunc func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	return tryFunc()
}
func TryCatch[T any](tryFunc func() (T, error)) (t T, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	return tryFunc()
}
func TryThrow(err error) {
	if err != nil {
		panic(err)
	}
}
