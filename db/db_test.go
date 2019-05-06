package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_DB(t *testing.T) {

	var err error

	dbRoot = &DatabaseRoot{}

	t.Run("root check reinitializes existing root", func(t *testing.T) {
		err = checkDbRoot()
		require.Nil(t, err)
	})

	dbRoot = nil

	t.Run("root check doesn't initialize non-existing root", func(t *testing.T) {
		err = checkDbRoot()
		require.Error(t, err)
	})

	dbRoot = &DatabaseRoot{}

	t.Run("doesn't retrieve root properly", func(t *testing.T) {
		root, err := getDefaultRoot()
		require.Nil(t, err)
		assert.Equal(t, dbRoot, root)
	})

	dbRoot = nil

}
