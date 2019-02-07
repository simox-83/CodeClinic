package csv

import (
	"encoding/csv"
	"io"
)

// more on interfaces: https://github.com/golang/go/wiki/CodeReviewComments#interfaces
type Getter interface {
	Get() (io.ReadCloser, error)
}

func Read(g Getter) ([][]string, error) {
	f, err := g.Get()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = '\t'
	records, err := r.ReadAll()

	return records, err
}
