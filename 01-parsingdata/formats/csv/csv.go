package csv

import (
	"encoding/csv"
	"io"
)

// Getter ... more on interfaces: https://github.com/golang/go/wiki/CodeReviewComments#interfaces
type Getter interface {
	Get() (io.ReadCloser, error)
}

// Read consuma l'interfaccia Getter, nel senso che la usa come parametro
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
