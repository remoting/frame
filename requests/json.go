package requests

import (
	"encoding/json"

	json2 "github.com/remoting/frame/json"
)

func ObjToStr(data map[string]any) (string, error) {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonStr), nil
}

func StrToObject(data string) (json2.Object, error) {
	obj := json2.Object{}
	err := json.Unmarshal([]byte(data), &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
