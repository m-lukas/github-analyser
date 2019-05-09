package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_DB(t *testing.T) {

	TestRoot = &DatabaseRoot{}

	t.Run("checkDbRoot(): root check doesn't initialize non-existing root", func(t *testing.T) {
		root, err := checkDbRoot()
		require.Error(t, err)
		require.Nil(t, root)
	})

	t.Run("getTestRoot(): doesn't return TestRoot", func(t *testing.T) {
		root := getTestRoot()
		require.Equal(t, TestRoot, root)
	})

	t.Run("getRoot(): got nil, expected TestRoot", func(t *testing.T) {
		root, err := getRoot()
		require.Nil(t, err)
		assert.Equal(t, TestRoot, root)
	})

	dbRoot = &DatabaseRoot{}

	t.Run("getDefaultRoot(): doesn't retrieve root properly", func(t *testing.T) {
		root, err := getDefaultRoot()
		require.Nil(t, err)
		assert.Equal(t, dbRoot, root)
	})

	dbRoot = nil
}
