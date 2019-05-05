package util_test //black-box testing

import (
	"testing"

	"github.com/m-lukas/github-analyser/util"

	"github.com/stretchr/testify/assert"
)

func Test_Calc(t *testing.T) {

	t.Run("average calc failed", func(t *testing.T) {
		slice := []float64{15.4, 24.1234, 16.001, 200.51423, 12.0, 50.999}
		expected := 53.17293833333334
		assert.Equal(t, expected, util.Avg(slice))
	})

	t.Run("distance between numbers", func(t *testing.T) {
		var origin float64
		var target float64
		var expected float64
		var distance float64

		origin = 34.543
		target = 22.3
		expected = 12.242999999999999
		distance = util.DistanceToNumber(origin, target)
		assert.Equal(t, expected, distance, "failed to get distance from negative")

		origin = 53.25
		target = 75.113
		expected = 21.863
		distance = util.DistanceToNumber(origin, target)
		assert.Equal(t, expected, distance, "failed to get distance from positive")
	})

	t.Run("biggest value in sorted array", func(t *testing.T) {
		expected := 22.98

		sliceSorted := []float64{1.12, 3.14, 5.0032, 10.5367, 16.001, 22.98}
		assert.Equal(t, expected, util.BiggestValueSorted(sliceSorted))

		sliceMixed := []float64{3.14, 22.98, 1.12, 16.001, 10.5367, 5.0032}
		assert.NotEqual(t, expected, util.BiggestValueSorted(sliceMixed))
	})
}
