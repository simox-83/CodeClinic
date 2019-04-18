package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/simox-83/CodeClinic/01-parsingdata/calculations"
	"github.com/simox-83/CodeClinic/01-parsingdata/formats/csv"
	"github.com/simox-83/CodeClinic/01-parsingdata/resources"
)

var resourceType string

func init() {
	flag.StringVar(&resourceType, "r", "", "specify the resource type (file, or http)")
}

func main() {
	flag.Parse()
	records := read(resourceType)
	viewMean(records)
	viewMedian(records)
}

func read(rt string) [][]string {
	var resource csv.Getter
	switch rt {
	case "http":
		resource = &resources.HTTP{
			URL: "http://localhost:8080/Environmental_Data_Deep_Moor_2015.csv",
		}
	case "file":
		resource = &resources.File{
			Name: "Environmental_Data_Deep_Moor_2015.txt",
		}
	default:
		log.Fatalf("unsupported resource '%s'", rt)
	}
	records, err := csv.Read(resource)
	if err != nil {
		log.Fatal(err)
	}
	return records
}

func viewMean(rec [][]string) {
	MeanTemp, MeanPressure, MeanWind := calculations.Mean(rec)
	fmt.Println("La temperatura media e' stata di", MeanTemp)
	fmt.Println("La pressione media e' stata di", MeanPressure)
	fmt.Println("Il vento medio e' stato di", MeanWind)
}

func viewMedian(rec [][]string) {
	//now we need to sort
	sortedTemp, sortedPressure, sortedWind := calculations.SortMatrix(rec)
	MedianTemp := calculations.Median(sortedTemp)
	MedianPressure := calculations.Median(sortedPressure)
	MedianWind := calculations.Median(sortedWind)

	fmt.Println("La mediana delle temperature e'", MedianTemp)
	fmt.Println("La mediana delle pressioni e'", MedianPressure)
	fmt.Println("La mediana del vento e'", MedianWind)
}
