package regression

import (
	"math"
	"strings"
)

func hasFileFormat(filepath string, format string) bool {

	subStrings := strings.Split(filepath, ".")
	ending := subStrings[len(subStrings)-1]

	if ending != format {
		return false
	}

	return true

}

func avg(values []float64) float64 {

	var sum float64

	for _, value := range values {
		sum += value
	}

	lenght := float64(len(values))
	average := sum / lenght

	return average
}

func distanceToNumber(origin float64, target float64) float64 {
	return math.Abs(target - origin)
}

func biggestValueSorted(array []float64) float64 {
	return array[len(array)-1]
}
