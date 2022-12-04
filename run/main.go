package main

import (
	"fmt"
)

func main() {

	data := ReadData("../data/data.csv")
	coordinates := ExtractCoordinates(data)
	fmt.Printf("%+v\n", coordinates)
}
