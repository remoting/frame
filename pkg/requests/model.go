package requests

import "net/http"

type HttpResponse struct {
	Data   []byte
	Status int
	Header map[string][]string
}
type FormFile struct {
	FileKey  string
	FilePath string
}
type HttpClient struct {
	client *http.Client
}
