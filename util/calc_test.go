package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Calc(t *testing.T) {
	slice := []float64{15.4, 24.1234, 16.001, 200.51423, 12.0, 50.999}

	t.Run("average calc failed", func(t *testing.T) {
		expected := 53.17293833333334
		assert.Equal(t, expected, Avg(slice))
	})
}
