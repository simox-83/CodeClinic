package resources

import (
	"errors"
	"io"
	"os"
)

type File struct {
	Name string
}

type HTTP struct {
}

type DB struct {
}

func (f *File) Get() (io.ReadCloser, error) {

	return os.Open(f.Name)

}

func (h *HTTP) Get() (io.ReadCloser, error) {
	return nil, errors.New("not implemented")
}

func (d *DB) Get() (io.ReadCloser, error) {
	return nil, errors.New("not implemented")
}
