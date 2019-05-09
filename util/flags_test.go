package util_test //black-box testing

import (
	"testing"

	"github.com/m-lukas/github-analyser/util"

	"github.com/stretchr/testify/require"
)

func Test_Flags(t *testing.T) {
	t.Run("IsTesting(): doesn't realize testing enviroment", func(t *testing.T) {
		testing := util.IsTesting()
		require.True(t, testing)
	})
}
