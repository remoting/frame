package requests

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
	Timeout int
}
