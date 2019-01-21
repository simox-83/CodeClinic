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

	// 2. usare github.com/stretchr/testify/assert al post di t.Errorf
	for _, el := range table {
		res1, res2, res3 := mean(el.input)
		assert.Equal(t, res1, el.out1, "Mean Temperature is wrong")
		assert.Equal(t, res2, el.out2, "Mean Pressure is wrong")
		assert.Equal(t, res3, el.out3, "Mean Wind is wrong")
	}

	// 3. aggiungere piu' test
}

func TestSortMatrix(t *testing.T) {
	table := []struct {
		input            [][]string
		out1, out2, out3 []float64
	}{
		{
			[][]string{
				[]string{"", "1", "1", "", "", "", "", "1"},
				[]string{"", "9", "7", "", "", "", "", "4"},
				[]string{"", "5", "3", "", "", "", "", "2"},
			},
			[]float64{5, 9},
			[]float64{3, 7},
			[]float64{2, 4},
		},
	}

	for _, el := range table {
		res1, res2, res3 := sortMatrix(el.input)
		assert.Equal(t, res1, el.out1, "Sorting Temperature is wrong")
		assert.Equal(t, res2, el.out2, "Sorting Pressure is wrong")
		assert.Equal(t, res3, el.out3, "Sorting Wind is wrong")
	}

}

func TestMedian(t *testing.T) {
	type ss struct {
		desc  string
		input []float64
		exp   float64
	}
	table := []ss{
		{
			"Median with odd elements is wrong",
			[]float64{
				1, 2, 3, 4, 5,
			},
			3,
		},
		{
			"Median with even elements is wrong",
			[]float64{
				1, 2, 3, 4,
			},
			2.5,
		},
	}

	for _, el := range table {
		res := median(el.input)
		assert.Equal(t, el.exp, res, el.desc)

	}
}
