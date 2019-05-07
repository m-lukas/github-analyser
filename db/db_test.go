package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_DB(t *testing.T) {

	TestRoot = &DatabaseRoot{}

	t.Run("root check reinitializes existing root", func(t *testing.T) {
		root, err := checkDbRoot()
		require.Nil(t, root)
		require.NotNil(t, err)
	})

	TestRoot = nil

	t.Run("root check doesn't initialize non-existing root", func(t *testing.T) {
		root, err := checkDbRoot()
		require.NotNil(t, root)
		require.Nil(t, err)
	})

	TestRoot = &DatabaseRoot{}

	t.Run("doesn't retrieve root properly", func(t *testing.T) {
		root, err := getDefaultRoot()
		require.Nil(t, err)
		assert.Equal(t, dbRoot, root)
	})

	TestRoot = nil

}
