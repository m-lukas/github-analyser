package controller

import "github.com/m-lukas/github-analyser/db"

func calcActivityScore(scores *db.Scores, config *db.ScoreParams) float64 {
	var score = 0.0

	score += config.FollowingW * scores.FollowingScore
	score += config.GistsW * scores.GistsScore
	score += config.IssuesW * scores.IssuesScore
	score += config.OrganizationsW * scores.OrganizationsScore
	score += config.ProjectsW * scores.ProjectsScore
	score += config.PullRequestsW * scores.PullRequestsScore
	score += config.ReposW * scores.ReposScore
	score += config.ContributionsW * scores.ContributionsScore
	score += config.StarredW * scores.StarredScore
	score += config.WatchingW * scores.WatchingScore
	score += config.CommitCommentsW * scores.CommitCommentsScore
	score += config.GistCommentsW * scores.GistCommentsScore
	score += config.IssueCommentsW * scores.IssueCommentsScore
	score += config.CommitFrequenzW * scores.CommitFrequenzScore

	score /= 14

	return score
}

func calcPopularityScore(scores *db.Scores, config *db.ScoreParams) float64 {
	var score = 0.0

	score += config.FollowersW * scores.FollowersScore
	score += config.StargazersW * scores.StargazersScore
	score += config.ForksW * scores.ForksScore

	score /= 3

	return score
}

func calcScores(user *db.User, config *db.ScoreParams) *db.Scores {

	scores := &db.Scores{
		FollowingScore:      ScoreFunc(float64(user.Following), config.FollowingK),
		FollowersScore:      ScoreFunc(float64(user.Followers), config.FollowersK),
		GistsScore:          ScoreFunc(float64(user.Gists), config.GistsK),
		IssuesScore:         ScoreFunc(float64(user.Issues), config.IssuesK),
		OrganizationsScore:  ScoreFunc(float64(user.Organizations), config.OrganizationsK),
		ProjectsScore:       ScoreFunc(float64(user.Projects), config.ProjectsK),
		PullRequestsScore:   ScoreFunc(float64(user.PullRequests), config.PullRequestsK),
		ContributionsScore:  ScoreFunc(float64(user.RepositoriesContributedTo), config.ContributionsK),
		StarredScore:        ScoreFunc(float64(user.StarredRepositories), config.StarredK),
		WatchingScore:       ScoreFunc(float64(user.Watching), config.Watchingk),
		CommitCommentsScore: ScoreFunc(float64(user.CommitComments), config.CommitCommentsK),
		GistCommentsScore:   ScoreFunc(float64(user.GistComments), config.GistCommentsK),
		IssueCommentsScore:  ScoreFunc(float64(user.IssueComments), config.IssueCommentsK),
		ReposScore:          ScoreFunc(float64(user.Repositories), config.ReposK),
		CommitFrequenzScore: ScoreFunc(user.CommitFrequenz, config.CommitFrequenzK),
		StargazersScore:     ScoreFunc(float64(user.Stargazers), config.StargazersK),
		ForksScore:          ScoreFunc(float64(user.Forks), config.ForksK),
	}

	return scores
}

func ScoreFunc(x float64, k float64) float64 {
	return x / (0.01*x + k)
}
