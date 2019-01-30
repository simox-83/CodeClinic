package csv

import (
	"encoding/csv"
	"io"
)

// more on interfaces: https://github.com/golang/go/wiki/CodeReviewComments#interfaces
type Getter interface {
	Get() (io.ReadCloser, error)
}

func Read(r io.Reader) ([][]string, error) {
	f := csv.NewReader(r)
	f.Comma = '\t'
	records, err := f.ReadAll()

	return records, err
}
