package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type Coordinates struct {
	X float64
	Y float64
}

func ReadData(path string) [][]string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func ExtractCoordinates(data [][]string) []Coordinates {
	var coordinates []Coordinates
	for i, line := range data {
		if i > 0 {
			var coord Coordinates
			for j, field := range line {
				// convert field to float64
				s, err := strconv.ParseFloat(field, 32)
				if err != nil {
					log.Fatal(err)
				}
				if j == 0 {
					coord.X = s
				}
				if j == 1 {
					coord.Y = s
				}
			}
			coordinates = append(coordinates, coord)
		}
	}
	return coordinates
}
