package calculations

import (
	"sort"
	"strconv"
)

// Mean calcola la media di temperatura, pressione e vento
func Mean(matrix [][]string) (float64, float64, float64) {
	var tottemp, totpres, totvent, conta float64

	for i, row := range matrix {
		if i == 0 {
			continue
		}
		temperatura, _ := strconv.ParseFloat(row[1], 64)
		pressione, _ := strconv.ParseFloat(row[2], 64)
		vento, _ := strconv.ParseFloat(row[7], 64)
		tottemp = temperatura + tottemp
		totpres = pressione + totpres
		totvent = vento + totvent
		conta++
	}
	return tottemp / conta, totpres / conta, totvent / conta
}

func SortMatrix(matrix [][]string) ([]float64, []float64, []float64) {
	var sortedTemp, sortedPressure, sortedWind []float64
	for i, row := range matrix {
		if i == 0 {
			continue
		}

		temperature, _ := strconv.ParseFloat(row[1], 64)
		pressure, _ := strconv.ParseFloat(row[2], 64)
		wind, _ := strconv.ParseFloat(row[7], 64)

		sortedTemp = append(sortedTemp, temperature)
		sortedPressure = append(sortedPressure, pressure)
		sortedWind = append(sortedWind, wind)
	}
	sort.Float64s(sortedTemp)
	sort.Float64s(sortedPressure)
	sort.Float64s(sortedWind)
	return sortedTemp, sortedPressure, sortedWind
}

func Median(f []float64) float64 {

	var mdn float64
	if len(f)%2 != 0 {
		mdn = f[len(f)/2]
	} else {
		middle := len(f) / 2
		higher := f[middle]
		lower := f[middle-1]
		mdn = (higher + lower) / 2

	}
	return mdn
}
