package graphql

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_CommitQuery(t *testing.T) {

	timeLayout := "2006-01-02T15:04:05.000Z"
	baseTime := "2019-05-05T14:10:00.000Z"
	dateTime, err := time.Parse(timeLayout, baseTime)
	require.Nil(t, err, "internal: time parse error")

	rawData, err := commitQueryTestResult()
	require.Nil(t, err, "internal: query parse error")

	t.Run("hours difference", func(t *testing.T) {

		startDate := dateTime.Add(-5*time.Hour - 30*time.Minute)
		endDate := dateTime

		require.Equal(t, 5.5, getHoursDifference(endDate, startDate))

		startDate = dateTime.Add(-20*time.Hour - 12*time.Minute - 15*time.Second)
		require.Equal(t, 20.204166666666666, getHoursDifference(endDate, startDate))
	})

	t.Run("commit frequenz calculation", func(t *testing.T) {
		expected := 194.21968055555556
		frequenz := GetCommitFrequenz(rawData, dateTime)
		require.Equal(t, expected, frequenz)

		rawDataEmpty := &CommitDataRaw{} //no commits
		expected = -1.0                  //negative value =~ error
		frequenz = GetCommitFrequenz(rawDataEmpty, dateTime)
		require.Equal(t, expected, frequenz)
	})
}
