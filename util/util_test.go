package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Util(t *testing.T) {
	t.Run("duration format is wrong", func(t *testing.T) {
		startTime := time.Now().Add(-1*time.Hour - 5*time.Minute - 20*time.Second)
		duration := time.Since(startTime)
		durationString := FormatDuration(duration)
		assert.Equal(t, "01h:05m:20s", durationString)
	})

	t.Run("poping n items from slice", func(t *testing.T) {

		popSlice := []string{
			"test1",
			"test2",
			"test3",
			"test4",
			"test5",
			"test6",
			"test7",
			"test8",
		}
		popLength := len(popSlice)
		n := 5

		popSlice, result1 := PopN(popSlice, n)
		assert.Equal(t, n, len(result1), "number of items doesn't match n")
		assert.Equal(t, popLength-n, len(popSlice), "items weren't removed from slice")

		popSlice, result2 := PopN(popSlice, n)
		assert.Equal(t, 3, len(result2), "deviation to n not fullfilled")
		assert.Equal(t, 0, len(popSlice), "items weren't removed from slice")
	})

	t.Run("remove duplicated", func(t *testing.T) {

		duplicateSlice := []string{
			"hallo",
			"hallö",
			"lukas",
			"berlin",
			"code",
			"hallo",
			"code",
			"hans",
		}
		expected := []string{"hallo", "hallö", "lukas", "berlin", "code", "hans"}

		output := RemoveDuplicates(duplicateSlice)
		assert.Equal(t, expected, output, "didn't remove duplicates properly")
	})
}
