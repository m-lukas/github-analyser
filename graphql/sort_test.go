package graphql

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Sort(t *testing.T) {

	//parse test-time from string and layout
	timeLayout := "2006-01-02T15:04:05.000Z"
	baseTime := "2019-05-05T14:10:00.000Z"
	dateTime, err := time.Parse(timeLayout, baseTime)
	require.Nil(t, err, "internal: time parse error")

	//create sample dates
	date1 := dateTime.Add(-4*24*time.Hour - 14*time.Hour - 15*time.Minute)
	date2 := dateTime.Add(-1*24*time.Hour - 3*time.Hour - 59*time.Minute)
	date3 := dateTime.Add(+2*24*time.Hour + 20*time.Hour + 42*time.Minute)

	//testTable with different orderings
	testTable := []dateSlice{
		{date2, date3, date1},
		{date3, date1, date2},
		{date3, date2, date1},
		{date1, date2, date3},
	}

	t.Run("sortDatesAsc(): sort date slice ascending", func(t *testing.T) {

		expected := dateSlice{date1, date2, date3}

		//sort every slice and check ordering
		for _, slice := range testTable {
			sortedSlice := sortDatesAsc(slice)
			assert.Equal(t, expected, sortedSlice)
		}

	})

}
