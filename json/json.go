package json

import "encoding/json"

func Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}
func NewObject() Object {
	return Object{}
}
func NewArray() Array {
	return Array{}
}
func ToObject(obj map[string]interface{}) Object {
	return obj
}
func ToArray(obj []interface{}) Array {
	return obj
}
func ObjectToString(data map[string]interface{}) (string, error) {
	jsonStr, err := Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonStr), nil
}

func StringToObject(data string) (Object, error) {
	obj := Object{}
	err := Unmarshal([]byte(data), &obj)
	if err != nil {
		return obj, err
	}
	return obj, nil
}
func ArrayToString(data []interface{}) (string, error) {
	jsonStr, err := Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonStr), nil
}

func StringToArray(data string) (Array, error) {
	obj := Array{}
	err := Unmarshal([]byte(data), &obj)
	if err != nil {
		return obj, err
	}
	return obj, nil
}
