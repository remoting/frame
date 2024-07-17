package requests

import (
	"bytes"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func NewHttpClient(params ...any) *HttpClient {
	timeout := 30
	if len(params) > 0 {
		_params, ok := params[0].(map[string]any)
		if ok {
			_timeout, _ok := _params["timeout"].(int)
			if _ok {
				timeout = _timeout
			}
		} else {
			_timeout, ok1 := params[0].(int)
			if ok1 {
				timeout = _timeout
			}
		}
	}
	return &HttpClient{
		client: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
	}
}
func NewCustomClient(_client *http.Client) *HttpClient {
	return &HttpClient{
		client: _client,
	}
}
func NewDisableClient(timeout int) *HttpClient {
	return &HttpClient{
		client: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
			Transport: &http.Transport{
				// 禁用连接池
				DisableKeepAlives: true,
				// 确保没有空闲连接
				MaxIdleConnsPerHost: 0,
				// 设置连接超时
				DialContext: (&net.Dialer{
					Timeout:   5 * time.Second,
					KeepAlive: 0,
				}).DialContext,
				// 设置TLS连接超时
				TLSHandshakeTimeout: 5 * time.Second,
			},
		},
	}
}
func (cli *HttpClient) getClient() *http.Client {
	return cli.client
}

func (cli *HttpClient) DoHttpRequest(method, url string, header map[string]string, _body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, _body)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	resp, err := cli.getClient().Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (cli *HttpClient) httpRequest(method, url string, header map[string]string, _body io.Reader) (*HttpResponse, error) {
	resp, err := cli.DoHttpRequest(method, url, header, _body)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	_resp := new(HttpResponse)
	_resp.Status = resp.StatusCode
	_resp.Data = body
	_resp.Header = make(map[string][]string, 0)
	for k := range resp.Header {
		_resp.Header[k] = resp.Header.Values(k)
	}
	return _resp, nil
}

func (cli *HttpClient) Get(url string, params, header map[string]string) (*HttpResponse, error) {
	_header := cli.initHeader(header)
	_, ok := _header["Content-Type"]
	if !ok {
		_header["Content-Type"] = "application/json"
	}
	return cli.httpRequest("GET", cli.getUrlParams(url, params), _header, nil)
}
func (cli *HttpClient) PostFormData(url string, params, header map[string]string) (*HttpResponse, error) {
	_header := cli.initHeader(header)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for name, value := range params {
		writer.WriteField(name, value)
	}
	writer.Close()
	_, ok := _header["Content-Type"]
	if !ok {
		_header["Content-Type"] = writer.FormDataContentType()
	}
	return cli.httpRequest("POST", url, _header, body)
}
func (cli *HttpClient) PostUrlData(url string, params, header map[string]string) (*HttpResponse, error) {
	_header := cli.initHeader(header)
	_, ok := _header["Content-Type"]
	if !ok {
		_header["Content-Type"] = "application/x-www-form-urlencoded"
	}
	return cli.httpRequest("POST", url, _header, strings.NewReader(cli.getParamString(params)))
}
func (cli *HttpClient) PostJson(url string, data map[string]any, header map[string]string) (*HttpResponse, error) {
	_header := cli.initHeader(header)
	_, ok := _header["Content-Type"]
	if !ok {
		_header["Content-Type"] = "application/json"
	}
	_body, err := ObjToStr(data)
	if err != nil {
		return nil, err
	}
	return cli.httpRequest("POST", url, _header, strings.NewReader(_body))
}
func (cli *HttpClient) PostFile(url string, files []*FormFile, params, header map[string]string) (*HttpResponse, error) {
	_header := cli.initHeader(header)
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	for _, file := range files {
		//关键的一步操作
		fileWriter, err := bodyWriter.CreateFormFile(file.FileKey, file.FilePath)
		if err != nil {
			return nil, err
		}
		//打开文件句柄操作
		fh, err := os.Open(file.FilePath)
		if err != nil {
			return nil, err
		}
		defer fh.Close()
		//iocopy
		_, err = io.Copy(fileWriter, fh)
		if err != nil {
			return nil, err
		}
	}

	for name, value := range params {
		bodyWriter.WriteField(name, value)
	}
	bodyWriter.Close()
	_, ok := _header["Content-Type"]
	if !ok {
		_header["Content-Type"] = bodyWriter.FormDataContentType()
	}
	return cli.httpRequest("POST", url, _header, bodyBuf)
}
func (cli *HttpClient) PostFileBody(url string, filePath string, header map[string]string) (*HttpResponse, error) {
	bodyBuf := &bytes.Buffer{}
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	_, err = bodyBuf.Write(data)
	if err != nil {
		return nil, err
	}
	return cli.httpRequest("POST", url, header, bodyBuf)
}

// getParamString
func (cli *HttpClient) getParamString(params map[string]string) string {
	q, _ := url.ParseQuery("")
	for name, value := range params {
		q.Add(name, value)
	}
	return q.Encode()
}
func (cli *HttpClient) getUrlParams(_url string, params map[string]string) string {
	if strings.Contains(_url, "?") {
		return _url + "&" + cli.getParamString(params)
	} else {
		return _url + "?" + cli.getParamString(params)
	}
}
func (cli *HttpClient) GetUrlParams(_url string, params map[string]string) string {
	return cli.getUrlParams(_url, params)
}
func (cli *HttpClient) initHeader(header map[string]string) map[string]string {
	if header == nil {
		return make(map[string]string, 0)
	}
	return header
}
func (cli *HttpClient) GetJsonHeader() map[string]string {
	header := make(map[string]string, 0)
	header["Content-Type"] = "application/json"
	return header
}
