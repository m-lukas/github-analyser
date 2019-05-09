package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Score(t *testing.T) {

	config := getTestScoreConfig()
	testUser := getTestUser()

	t.Run("SetScore(): scores were calculated wrong", func(t *testing.T) {

		//calculate score and assign them to user
		SetScore(testUser, config)

		//check datailed scores
		scores := testUser.Scores
		assert.Equal(t, 69.69696969696969, scores.FollowersScore, "FollowersScore")
		assert.Equal(t, 57.89473684210526, scores.FollowingScore, "FollowingScore")
		assert.Equal(t, 99.9000999000999, scores.GistsScore, "GistsScore")
		assert.Equal(t, 98.98656610888521, scores.IssuesScore, "IssuesScore")
		assert.Equal(t, 83.33333333333333, scores.OrganizationsScore, "OrganizationsScore")
		assert.Equal(t, 99.98330278843433, scores.ProjectsScore, "ProjectsScore")
		assert.Equal(t, 48.90510948905109, scores.PullRequestsScore, "PullRequestsScore")
		assert.Equal(t, 75.0, scores.ContributionsScore, "ContributionsScore")
		assert.Equal(t, 42.14876033057851, scores.StarredScore, "StarredScore")
		assert.Equal(t, 34.131736526946106, scores.WatchingScore, "WatchingScore")
		assert.Equal(t, 69.69696969696969, scores.CommitCommentsScore, "CommitCommentsScore")
		assert.Equal(t, 97.65624999999999, scores.GistCommentsScore, "GistCommentsScore")
		assert.Equal(t, 96.07843137254902, scores.IssueCommentsScore, "IssueCommentsScore")
		assert.Equal(t, 42.30769230769231, scores.ReposScore, "ReposScore")
		assert.Equal(t, 27.07826932425863, scores.CommitFrequenzScore, "CommitFrequenzScore")
		assert.Equal(t, 39.02439024390244, scores.StargazersScore, "StargazersScore")
		assert.Equal(t, 33.333333333333336, scores.ForksScore, "ForksScore")

		//check general scores
		assert.Equal(t, 69.50723271577881, testUser.ActivityScore, "ActivityScore")
		assert.Equal(t, 47.35156442473516, testUser.PopularityScore, "PopularityScore")
	})
}
