package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/simox-83/CodeClinic/01-parsingdata/formats/csv"
	"github.com/simox-83/CodeClinic/01-parsingdata/resources"
)

func main() {

	/*
	      The idea was to get the data from the URL below, but it doesn't exist anymore. We use a local file instead.

	   	url, err := http.Get("http://lpo.dt.navy.mil/data/DM/Environmental_Data_Deep_Moor_2015.txt")
	   	if err != nil {
	   		log.Fatal(err)
	   	}
	       r := csv.NewReader(url.Body)
	*/

	records, err := csv.Read(&resources.File{
		Name: "Environmental_Data_Deep_Moor_2015.txt",
	})

	if err != nil {
		log.Fatal(err)
	}

	MeanTemp, MeanPressure, MeanWind := mean(records)
	fmt.Println("La temperatura media e' stata di", MeanTemp)
	fmt.Println("La pressione media e' stata di", MeanPressure)
	fmt.Println("Il vento medio e' stato di", MeanWind)

	//now we need to sort
	sortedTemp, sortedPressure, sortedWind := sortMatrix(records)
	MedianTemp := median(sortedTemp)
	MedianPressure := median(sortedPressure)
	MedianWind := median(sortedWind)

	fmt.Println("La mediana delle temperature e'", MedianTemp)
	fmt.Println("La mediana delle pressioni e'", MedianPressure)
	fmt.Println("La mediana del vento e'", MedianWind)

}

func mean(matrix [][]string) (float64, float64, float64) {
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

func sortMatrix(matrix [][]string) ([]float64, []float64, []float64) {
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

func median(f []float64) float64 {

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
