package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
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

	file, err := os.Open("Environmental_Data_Deep_Moor_2015.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//fmt.Printf("File opened, %v", file)

	r := csv.NewReader(file)
	r.Comma = '\t'
	records, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("primo valore di temperatura e' %v\n", records[1][0])
	MeanTemp, MeanPressure, MeanWind := mean(records)
	fmt.Println("La temperatura media e' stata di", MeanTemp)
	fmt.Println("La pressione media e' stata di", MeanPressure)
	fmt.Println("Il vento medio e' stato di", MeanWind)

	//fmt.Print(records)

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
		if i != 0 {
			temperatura, _ := strconv.ParseFloat(row[1], 64)
			pressione, _ := strconv.ParseFloat(row[2], 64)
			vento, _ := strconv.ParseFloat(row[7], 64)
			tottemp = temperatura + tottemp
			totpres = pressione + totpres
			totvent = vento + totvent
			conta++
		}
	}
	return tottemp / conta, totpres / conta, totvent / conta
}

func sortMatrix(matrix [][]string) ([]float64, []float64, []float64) {
	var sortedTemp, sortedPressure, sortedWind []float64
	for i, row := range matrix {
		if i != 0 {
			temperature, _ := strconv.ParseFloat(row[1], 64)
			pressure, _ := strconv.ParseFloat(row[2], 64)
			wind, _ := strconv.ParseFloat(row[7], 64)
			sortedTemp = append(sortedTemp, temperature)
			sortedPressure = append(sortedPressure, pressure)
			sortedWind = append(sortedWind, wind)
		}
	}
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
		mdn = higher + lower/2

	}
	return mdn
}
