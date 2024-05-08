package requests

var _httpclient = NewHttpClient(30)

func Get(url string, params, header map[string]string) (*HttpResponse, error) {
	return _httpclient.Get(url, params, header)
}
func PostFormData(url string, params, header map[string]string) (*HttpResponse, error) {
	return _httpclient.PostFormData(url, params, header)
}
func PostUrlData(url string, params, header map[string]string) (*HttpResponse, error) {
	return _httpclient.PostUrlData(url, params, header)
}
func PostJson(url string, data map[string]any, header map[string]string) (*HttpResponse, error) {
	return _httpclient.PostJson(url, data, header)
}
func PostFile(url string, files []*FormFile, params, header map[string]string) (*HttpResponse, error) {
	return _httpclient.PostFile(url, files, params, header)
}
func PostFileBody(url string, filePath string, header map[string]string) (*HttpResponse, error) {
	return _httpclient.PostFileBody(url, filePath, header)
}
