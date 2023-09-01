package json

import (
	"github.com/remoting/frame/util"
)

type Object map[string]interface{}
type Array []interface{}

func (array Array) Size() int {
	return len(array)
}
func (array Array) GetArray(i int) Array {
	obj := array[i]
	if ret, ok := obj.([]interface{}); ok {
		return ret
	}
	return nil
}

func (array Array) GetObject(i int) Object {
	obj := array[i]
	if ret, ok := obj.(map[string]interface{}); ok {
		return ret
	}
	return nil
}
func NewObject() Object {
	return Object{}
}
func ToObject(obj map[string]interface{}) Object {
	return obj
}
func ToArray(obj []interface{}) Array {
	return obj
}
func (json Object) GetObject(name string) Object {
	if obj, ok1 := json[name]; ok1 {
		if ret, ok := obj.(map[string]interface{}); ok {
			return ret
		}
		if ret, ok := obj.(Object); ok {
			return ret
		}
	}
	return nil
}
func (json Object) GetArray(name string) Array {
	if obj, ok1 := json[name]; ok1 {
		if ret, ok := obj.([]interface{}); ok {
			return ret
		}
		if ret, ok := obj.(Array); ok {
			return ret
		}
	}
	return nil
}
func (json Object) GetString(name string) string {
	value, ok := json[name]
	if ok {
		return util.String(value)
	}
	return ""
}
func (json Object) GetInt64(name string) int64 {
	var t2 int64
	t1, ok := json[name]
	if ok {
		return util.Int64(t1)
	}
	return t2
}
func (json Object) Contains(name string) bool {
	_, ok := json[name]
	return ok
}
func (json Object) GetInt(name string) int {
	var t2 int
	t1, ok := json[name]
	if ok {
		t2 = util.Int(t1)
	}
	return t2
}
func (json Object) Keys() []string {
	keys := make([]string, 0, len(json))
	for k := range json {
		keys = append(keys, k)
	}
	return keys
}
func (json Object) Put(name string, val any) {
	json[name] = val
}
