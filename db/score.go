package db

import (
	"errors"
	"log"
)

func (root *DatabaseRoot) InitScoreConfig() error {

	if root.RedisClient == nil {
		err := root.InitRedisClient()
		if err != nil {
			return errors.New("redis client couldn't be reinitialized!")
		}
	}

	redisClient, err := GetRedis()
	if err != nil {
		return err
	}

	followingK := redisClient.GetScoreParam("following", "k")
	followingW := redisClient.GetScoreParam("following", "w")
	followersK := redisClient.GetScoreParam("followers", "k")
	followersW := redisClient.GetScoreParam("followers", "w")
	gistsK := redisClient.GetScoreParam("gists", "k")
	gistsW := redisClient.GetScoreParam("gists", "w")
	issuesK := redisClient.GetScoreParam("issues", "k")
	issuesW := redisClient.GetScoreParam("issues", "w")
	organizationsK := redisClient.GetScoreParam("organizations", "k")
	organizationsW := redisClient.GetScoreParam("organizations", "w")
	projectsK := redisClient.GetScoreParam("projects", "k")
	projectsW := redisClient.GetScoreParam("projects", "w")
	pullRequestsK := redisClient.GetScoreParam("pull_requests", "k")
	pullRequestsW := redisClient.GetScoreParam("pull_requests", "w")
	contributionsK := redisClient.GetScoreParam("contributions", "k")
	contributionsW := redisClient.GetScoreParam("contributions", "k")
	starredK := redisClient.GetScoreParam("starred", "k")
	starredW := redisClient.GetScoreParam("starred", "w")
	watchingK := redisClient.GetScoreParam("watching", "k")
	watchingW := redisClient.GetScoreParam("watching", "w")
	commitCommentsK := redisClient.GetScoreParam("commit_comments", "k")
	commitCommentsW := redisClient.GetScoreParam("commit_comments", "w")
	gistCommentsK := redisClient.GetScoreParam("gist_comments", "k")
	gistCommentsW := redisClient.GetScoreParam("gist_comments", "w")
	issueCommentsK := redisClient.GetScoreParam("issue_comments", "k")
	issueCommentsW := redisClient.GetScoreParam("issue_comments", "w")
	reposK := redisClient.GetScoreParam("repos", "k")
	reposW := redisClient.GetScoreParam("repos", "w")
	commitFrequenzK := redisClient.GetScoreParam("commit_frequenz", "k")
	commitFrequenzW := redisClient.GetScoreParam("commit_frequenz", "w")
	stargazersK := redisClient.GetScoreParam("stargazers", "k")
	stargazersW := redisClient.GetScoreParam("stargazers", "w")
	forksK := redisClient.GetScoreParam("forks", "k")
	forksW := redisClient.GetScoreParam("forks", "w")

	scoreConfig := &ScoreParams{
		FollowingK:      followingK,
		FollowingW:      followingW,
		FollowersK:      followersK,
		FollowersW:      followersW,
		GistsK:          gistsK,
		GistsW:          gistsW,
		IssuesK:         issuesK,
		IssuesW:         issuesW,
		OrganizationsK:  organizationsK,
		OrganizationsW:  organizationsW,
		ProjectsK:       projectsK,
		ProjectsW:       projectsW,
		PullRequestsK:   pullRequestsK,
		PullRequestsW:   pullRequestsW,
		ContributionsK:  contributionsK,
		ContributionsW:  contributionsW,
		StarredK:        starredK,
		StarredW:        starredW,
		Watchingk:       watchingK,
		WatchingW:       watchingW,
		CommitCommentsK: commitCommentsK,
		CommitCommentsW: commitCommentsW,
		GistCommentsK:   gistCommentsK,
		GistCommentsW:   gistCommentsW,
		IssueCommentsK:  issueCommentsK,
		IssueCommentsW:  issueCommentsW,
		ReposK:          reposK,
		ReposW:          reposW,
		CommitFrequenzK: commitFrequenzK,
		CommitFrequenzW: commitFrequenzW,
		StargazersK:     stargazersK,
		StargazersW:     stargazersW,
		ForksK:          forksK,
		ForksW:          forksW,
	}

	root.ScoreConfig = scoreConfig
	log.Println("Initialized score config!")

	return nil

}
