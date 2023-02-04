package client

import "math"

func MeanCoordinate(data []Coordinates, c string) float64 {
	var sum float64 = 0
	for _, entry := range data {
		if c == "X" {
			sum += entry.X
		} else if c == "Y" {
			sum += entry.Y
		}
	}
	return sum / float64(len(data))
}

func GetBHat(data []Coordinates) float64 {
	yMean := MeanCoordinate(data, "Y")
	xMean := MeanCoordinate(data, "X")
	var cov float64 = 0
	var xVariance float64 = 0
	for _, entry := range data {
		cov += (entry.X - xMean) * (entry.Y - yMean)
		xVariance += math.Pow(entry.X-xMean, 2)
	}
	return cov / xVariance
}

func GetCoefficients(data []Coordinates) (float64, float64) {
	// Formula: y = bx + a
	var aHat float64
	var bHat float64
	yMean := MeanCoordinate(data, "Y")
	xMean := MeanCoordinate(data, "X")
	bHat = GetBHat(data)
	aHat = yMean - bHat*xMean
	return aHat, bHat

}
