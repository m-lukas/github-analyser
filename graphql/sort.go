package graphql

import (
	"sort"
	"time"
)

type dateSlice []time.Time //custom time slice

//sorting helper functions
func (slice dateSlice) Less(i, j int) bool {
	return slice[i].Before(slice[j])
}
func (slice dateSlice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
func (slice dateSlice) Len() int {
	return len(slice)
}

func sortDatesAsc(slice dateSlice) dateSlice {

	sort.Sort(slice)
	return slice
}
