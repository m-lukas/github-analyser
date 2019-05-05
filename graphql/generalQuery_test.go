package graphql

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func Test_GeneralQuery(t *testing.T) {

	rawData, err := generalQueryTestResult()
	require.Nil(t, err, "internal: query parse error")

	t.Run("stargazer and fork sum", func(t *testing.T) {

		var stargazers int
		var forks int
		var expectedStargazers int
		var expectedForks int

		stargazers, forks = calcStargazersAndForks(rawData, "m-lukas") //shouldn't count repo of "somebody" in
		expectedStargazers = 104
		expectedForks = 41
		assert.Equal(t, expectedStargazers, stargazers, "stargazers of m-lukas")
		assert.Equal(t, expectedForks, forks, "forks of m-lukas")

		stargazers, forks = calcStargazersAndForks(rawData, "somebody") //shouldn't count repo of "m-lukas" in
		expectedStargazers = 5
		expectedForks = 2
		assert.Equal(t, expectedStargazers, stargazers, "stargazers of somebody")
		assert.Equal(t, expectedForks, forks, "forks of somebody")

		stargazers, forks = calcStargazersAndForks(rawData, "") //shouldn't count any repo in
		expectedStargazers = 0
		expectedForks = 0
		assert.Equal(t, expectedStargazers, stargazers, "no repositories")
		assert.Equal(t, expectedForks, forks, "no repositories")
	})

	t.Run("raw data conversion", func(t *testing.T) {

		data := convertGeneralData(rawData)
		rawLogin := rawData.RepositoryOwner.Login
		rawMail := rawData.RepositoryOwner.Email
		rawName := rawData.RepositoryOwner.Name
		rawRepos := rawData.RepositoryOwner.Repositories.TotalCount

		assert.Equal(t, rawLogin, data.Login)
		assert.Equal(t, rawMail, data.Email)
		assert.Equal(t, rawName, data.Name)
		assert.Equal(t, rawRepos, data.Repositories)
	})

}
