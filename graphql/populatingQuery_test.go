package graphql

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func Test_PopulatingQuery(t *testing.T) {

	rawData, err := populatingQueryTestResult()
	require.Nil(t, err, "internal: query parse error")

	t.Run("login list", func(t *testing.T) {

		expected := []string{"LBeul", "fluidsonic", "m-lukas", "florianstahr", "toorusr", "Nickramas"}
		assert.Equal(t, expected, GetLoginList(rawData))

	})

}
