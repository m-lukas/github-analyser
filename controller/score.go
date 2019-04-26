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
		FollowingScore:      scoreFunc(float64(user.Following), config.FollowingK),
		FollowersScore:      scoreFunc(float64(user.Followers), config.FollowersK),
		GistsScore:          scoreFunc(float64(user.Gists), config.GistsK),
		IssuesScore:         scoreFunc(float64(user.Issues), config.IssuesK),
		OrganizationsScore:  scoreFunc(float64(user.Organizations), config.OrganizationsK),
		ProjectsScore:       scoreFunc(float64(user.Projects), config.ProjectsK),
		PullRequestsScore:   scoreFunc(float64(user.PullRequests), config.PullRequestsK),
		ContributionsScore:  scoreFunc(float64(user.RepositoriesContributedTo), config.ContributionsK),
		StarredScore:        scoreFunc(float64(user.StarredRepositories), config.StarredK),
		WatchingScore:       scoreFunc(float64(user.Watching), config.Watchingk),
		CommitCommentsScore: scoreFunc(float64(user.CommitComments), config.CommitCommentsK),
		GistCommentsScore:   scoreFunc(float64(user.GistComments), config.GistCommentsK),
		IssueCommentsScore:  scoreFunc(float64(user.IssueComments), config.IssueCommentsK),
		ReposScore:          scoreFunc(float64(user.Repositories), config.ReposK),
		CommitFrequenzScore: scoreFunc(user.CommitFrequenz, config.CommitFrequenzK),
		StargazersScore:     scoreFunc(float64(user.Stargazers), config.StargazersK),
		ForksScore:          scoreFunc(float64(user.Forks), config.ForksK),
	}

	return scores
}

func scoreFunc(x float64, k float64) float64 {
	return x / (0.01*x + k)
}
