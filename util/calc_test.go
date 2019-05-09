package util_test //black-box testing

import (
	"testing"

	"github.com/m-lukas/github-analyser/util"

	"github.com/stretchr/testify/assert"
)

func Test_Calc(t *testing.T) {

	t.Run("Avg(): average calc failed", func(t *testing.T) {
		slice := []float64{15.4, 24.1234, 16.001, 200.51423, 12.0, 50.999}
		expected := 53.17293833333334
		assert.Equal(t, expected, util.Avg(slice))
	})

	t.Run("DistanceToNumber(): distance between numbers", func(t *testing.T) {
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

	t.Run("BiggestValueSorted(): biggest value in sorted array", func(t *testing.T) {
		expected := 22.98

		sliceSorted := []float64{1.12, 3.14, 5.0032, 10.5367, 16.001, 22.98}
		assert.Equal(t, expected, util.BiggestValueSorted(sliceSorted))

		sliceMixed := []float64{3.14, 22.98, 1.12, 16.001, 10.5367, 5.0032}
		assert.NotEqual(t, expected, util.BiggestValueSorted(sliceMixed))
	})

	t.Run("NearestDistance(): nearest distance", func(t *testing.T) {
		testTable := []struct {
			Target   float64
			First    float64
			Second   float64
			Third    float64
			Expected float64
		}{
			{Target: 15.0, First: 14.25, Second: 18.33, Third: -15.25, Expected: 14.25},
			{Target: 2.335, First: 10.1333, Second: 10.1332, Third: -12.02, Expected: 10.1332},
			{Target: -22.99, First: -25.5, Second: -1.764, Third: -22.98, Expected: -22.98},
		}

		for _, set := range testTable {
			output := util.NearestDistance(set.Target, set.First, set.Second, set.Third)
			assert.Equal(t, set.Expected, output)
		}
	})

	t.Run("calc k from y", func(t *testing.T) {
		y := 50.0
		x := 125.25
		expected := 1.2525

		k := util.CalcKFromY(y, x)
		assert.Equal(t, expected, k)
	})

}
