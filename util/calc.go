package util

func Avg(values []float64) float64 {

	var sum float64

	for _, value := range values {
		sum += value
	}

	lenght := float64(len(values))
	average := sum / lenght

	return average
}
