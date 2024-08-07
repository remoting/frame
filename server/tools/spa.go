package tools

import (
	"errors"
	"net/http"
	"os"
	"strings"
)

func ServeFile(fileSystem http.FileSystem, filepath, index string) (string, http.File, error) {
	f, err := fileSystem.Open(filepath)
	if err != nil {
		//有错误
		if errors.Is(err, os.ErrNotExist) {
			filepath = index
			f, err = fileSystem.Open(index)
		} else {
			return "", nil, err
		}
	}
	if err != nil {
		return "", nil, err
	}
	d, err := f.Stat()
	if d.IsDir() {
		if !strings.HasSuffix(filepath, "/") {
			filepath = filepath + "/"
		}
		filepath = filepath + "index.html"
		return ServeFile(fileSystem, filepath, index)
	} else {
		return filepath, f, nil
	}
}
