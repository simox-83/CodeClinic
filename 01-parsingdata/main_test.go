package main

import "testing"

func TestMean(t *testing.T) {
	input := [][]string{
		[]string{"", "3", "3", "", "", "", "", "6"},
		[]string{"", "3", "3", "", "", "", "", "6"},
	}
	exp1, exp2, exp3 := mean(input)

	if exp1 != 3.0 {
		t.Errorf("exp1 got %f, expected %f", exp1, 3.0)
	}
	if exp2 != 3.0 {
		t.Errorf("exp2 got %f, expected %f", exp1, 3.0)
	}
	if exp3 != 6.0 {
		t.Errorf("exp3 got %f, expected %f", exp1, 6.0)
	}
}
