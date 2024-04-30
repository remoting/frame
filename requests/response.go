package requests

import (
	"errors"
	"strconv"

	"github.com/remoting/frame/json"
)

func (resp *HttpResponse) GetText() (string, error) {
	if resp.Status == 200 {
		return string(resp.Data), nil
	} else {
		return string(resp.Data), errors.New("http error " + strconv.Itoa(resp.Status))
	}
}
func (resp *HttpResponse) GetHeader(key string) string {
	if resp.Header == nil {
		return ""
	}
	v := resp.Header[key]
	if len(v) == 0 {
		return ""
	}
	return v[0]
}
func (resp *HttpResponse) GetJson() (json.Object, error) {
	if resp.Status == 200 {
		obj, err := StrToObject(string(resp.Data))
		if err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, errors.New("http error " + strconv.Itoa(resp.Status))
	}
}
