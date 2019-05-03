package util

import "math"

func Avg(values []float64) float64 {

	var sum float64

	for _, value := range values {
		sum += value
	}

	lenght := float64(len(values))
	average := sum / lenght

	return average
}

func DistanceToNumber(origin float64, target float64) float64 {
	return math.Abs(target - origin)
}

func BiggestValueSorted(array []float64) float64 {
	return array[len(array)-1]
}
