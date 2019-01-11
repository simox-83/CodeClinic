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
	var tottemp = 0.0
	var totpres = 0.0
	var totvent = 0.0
	var conta = 0.0
	for i, row := range records {
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
	fmt.Println("La temperatura media e' stata di", (tottemp / conta))
	fmt.Println("La pressione media e' stata di", totpres/conta)
	fmt.Println("Il vento medio e' stato di", totvent/conta)

	//fmt.Print(records)

	//now we need to sort
	var sortedTemp, sortedPressure, sortedWind []float64
	for i, row := range records {
		if i != 0 {
			temperature, _ := strconv.ParseFloat(row[1], 64)
			pressure, _ := strconv.ParseFloat(row[2], 64)
			wind, _ := strconv.ParseFloat(row[7], 64)
			sortedTemp = append(sortedTemp, temperature)
			sortedPressure = append(sortedPressure, pressure)
			sortedWind = append(sortedWind, wind)
		}
	}
	var medianTemp float64
	if len(sortedTemp)%2 != 0 {
		medianTemp = sortedTemp[len(sortedTemp)/2]
	} else {
		middle := len(sortedTemp) / 2
		higher := sortedTemp[middle]
		lower := sortedTemp[middle-1]
		medianTemp = higher + lower/2

	}
	fmt.Println("Median temp is", medianTemp)
}
