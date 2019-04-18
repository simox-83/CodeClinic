package resources

import (
	"errors"
	"io"
	"net/http"
	"os"
)

type File struct {
	Name string
}

type HTTP struct {
	URL string
}

type DB struct {
}

func (f *File) Get() (io.ReadCloser, error) {

	return os.Open(f.Name)

}

func (h *HTTP) Get() (io.ReadCloser, error) {
	resp, err := http.Get(h.URL)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil

}

func (d *DB) Get() (io.ReadCloser, error) {
	return nil, errors.New("not implemented")
}
