package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMean(t *testing.T) {
	table := []struct {
		input            [][]string
		out1, out2, out3 float64
	}{
		{
			[][]string{
				[]string{"", "1", "1", "", "", "", "", "1"},
				[]string{"", "3", "3", "", "", "", "", "6"},
				[]string{"", "7", "5", "", "", "", "", "9"},
			},
			5.0, 4.0, 7.5,
		},
	}
	/*
		for _, el := range table {
			expected1, expected2, expected3 := mean(el.input)
			if expected1 != el.out1 {
				t.Errorf("expected1 got %f, expected %f", el.out1, 5.000000)
			}
			if expected2 != el.out2 {
				t.Errorf("expected2 got %f, expected %f", el.out2, 4.000000)
			}
			if expected3 != el.out3 {
				t.Errorf("expected3 got %f, expected %f", el.out3, 7.500000)
			}
		}
	*/

	// 2. usare github.com/stretchr/testify/assert al post di t.Errorf
	for _, el := range table {
		res1, res2, res3 := mean(el.input)
		assert.Equal(t, res1, el.out1, "Mean Temperature is wrong")
		assert.Equal(t, res2, el.out2, "Mean Pressure is wrong")
		assert.Equal(t, res3, el.out3, "Mean Wind is wrong")
	}

	// 3. aggiungere piu' test
}
