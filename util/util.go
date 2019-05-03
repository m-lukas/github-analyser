package util

import (
	"fmt"
	"time"
)

func FormatDuration(duration time.Duration) string {
	duration = duration.Round(time.Second)
	h := duration / time.Hour
	duration -= h * time.Hour
	m := duration / time.Minute
	duration -= m * time.Minute
	s := duration / time.Second

	return fmt.Sprintf("%02dh:%02dm:%02ds", h, m, s)
}

func PopN(slice []string, n int) ([]string, []string) {

	if len(slice) == 0 {
		return nil, []string{}
	}
	if len(slice) < n {
		n = len(slice)
	}

	poped, origin := slice[0:n], slice[n:]
	return origin, poped
}
