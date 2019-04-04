package main

import (
	"fmt"
	"log"

	"github.com/simox-83/CodeClinic/01-parsingdata/calculations"
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

	MeanTemp, MeanPressure, MeanWind := calculations.Mean(records)
	fmt.Println("La temperatura media e' stata di", MeanTemp)
	fmt.Println("La pressione media e' stata di", MeanPressure)
	fmt.Println("Il vento medio e' stato di", MeanWind)

	//now we need to sort
	sortedTemp, sortedPressure, sortedWind := calculations.SortMatrix(records)
	MedianTemp := calculations.Median(sortedTemp)
	MedianPressure := calculations.Median(sortedPressure)
	MedianWind := calculations.Median(sortedWind)

	fmt.Println("La mediana delle temperature e'", MedianTemp)
	fmt.Println("La mediana delle pressioni e'", MedianPressure)
	fmt.Println("La mediana del vento e'", MedianWind)

}
