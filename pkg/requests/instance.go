package requests

import (
	"context"
	"io"
)

var _httpclient = NewHttpClient(30)

func ConfigHttpClient(_client *HttpClient) {
	_httpclient = _client
}

func Get(url string, params, header map[string]string) (*HttpResponse, error) {
	return _httpclient.Get(context.Background(), url, params, header)
}
func PostFormData(url string, params, header map[string]string) (*HttpResponse, error) {
	return _httpclient.PostFormData(context.Background(), url, params, header)
}
func PostUrlData(url string, params, header map[string]string) (*HttpResponse, error) {
	return _httpclient.PostUrlData(context.Background(), url, params, header)
}
func PostJson(url string, data map[string]any, header map[string]string) (*HttpResponse, error) {
	return _httpclient.PostJson(context.Background(), url, data, header)
}
func PostBody(url string, body io.Reader, header map[string]string) (*HttpResponse, error) {
	return _httpclient.PostBody(context.Background(), url, body, header)
}
func PostFile(url string, files []*FormFile, params, header map[string]string) (*HttpResponse, error) {
	return _httpclient.PostFile(context.Background(), url, files, params, header)
}
func PostFileBody(url string, filePath string, header map[string]string) (*HttpResponse, error) {
	return _httpclient.PostFileBody(context.Background(), url, filePath, header)
}
