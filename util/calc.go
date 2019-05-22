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

func NearestDistance(target float64, first float64, second float64, third float64) float64 {

	distanceFromFirst := DistanceToNumber(first, target)
	distanceFromSecond := DistanceToNumber(second, target)
	distanceFromThird := DistanceToNumber(third, target)

	if distanceFromFirst < distanceFromSecond && distanceFromSecond < distanceFromThird {
		return first
	}

	if distanceFromSecond < distanceFromFirst && distanceFromSecond < distanceFromThird {
		return second
	}

	return third

}

func DistanceToNumber(origin float64, target float64) float64 {
	return math.Abs(target - origin)
}

func BiggestValueSorted(array []float64) float64 {
	return array[len(array)-1]
}

func CalcKFromY(y float64, x float64) float64 {
	return x/y - 0.01*x
}
