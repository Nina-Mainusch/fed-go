package main

import (
	"fmt"
)

func main() {

	// Client 1
	data := ReadData("../data/data.csv")
	coordinates := ExtractCoordinates(data)
	// Do linear regression
	aHat, bHat := getCoefficients(coordinates)
	fmt.Printf("%f, %f \n", aHat, bHat)

	// CLient 2
	data2 := ReadData("../data/data2.csv")
	coordinates2 := ExtractCoordinates(data2)
	aHat2, bHat2 := getCoefficients(coordinates2)
	fmt.Printf("%f, %f \n", aHat2, bHat2)

	// Do linear regression together - exchange estimates
	// Do homomorphic encryption
	// send updates
	// Compute fed avg
	aHatGlobal := Avg(aHat, aHat2)
	bHatGlobal := Avg(bHat, bHat2)
	fmt.Printf("%f, %f \n", aHatGlobal, bHatGlobal)

}
