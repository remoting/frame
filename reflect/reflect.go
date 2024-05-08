package reflect

import "reflect"

func GetAnyValue(v any) any {
	x := ReflectGetPtrValue(reflect.ValueOf(v))
	if x.IsValid() {
		return x.Interface()
	} else {
		return nil
	}
}
func IsNil(v any) bool {
	return GetAnyValue(v) == nil
}
func ReflectGetPtrValue(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Ptr {
		return ReflectGetPtrValue(v.Elem())
	}
	return v
}
