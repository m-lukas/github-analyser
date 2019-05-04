package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Flags(t *testing.T) {
	t.Run("doesn't recognise testing enviroment", func(t *testing.T) {
		testing := IsTesting()
		require.True(t, testing)
	})
}
