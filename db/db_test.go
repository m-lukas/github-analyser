package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DB(t *testing.T) {

	dbRoot = nil

	t.Run("db root check can't initialize db", func(t *testing.T) {
		err := checkDbRoot()
		assert.Nil(t, err)
	})
	t.Run("cannot initialize db root", func(t *testing.T) {
		root, err := getRoot()
		assert.Nil(t, err)
		assert.NotNil(t, root)
	})
}
