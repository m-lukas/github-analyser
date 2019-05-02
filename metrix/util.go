package metrix

import (
	"fmt"
	"math"
	"time"
)

//TODO: DUPLICATE! -> move to utils

func distanceToNumber(origin float64, target float64) float64 {
	return math.Abs(target - origin)
}

func biggestValueSorted(array []float64) float64 {
	return array[len(array)-1]
}

func formatDuration(duration time.Duration) string {
	duration = duration.Round(time.Second)
	h := duration / time.Hour
	duration -= h * time.Hour
	m := duration / time.Minute
	duration -= m * time.Minute
	s := duration / time.Second

	return fmt.Sprintf("%02dh:%02dm:%02ds", h, m, s)
}

func popN(slice []string, n int) ([]string, []string) {

	if len(slice) == 0 {
		return nil, []string{}
	}

	if len(slice) < n {
		n = len(slice)
	}

	poped, origin := slice[0:n], slice[n:]

	return origin, poped

}
