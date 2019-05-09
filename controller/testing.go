package controller

import (
	"context"
	"testing"

	"github.com/m-lukas/github-analyser/db"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
)

//setupMongoTest initializes the mongo client and provides a cleared collection
func setupMongoTest(t *testing.T, root *db.DatabaseRoot, collectionName string, ctx context.Context) (*db.MongoClient, *mongo.Collection) {

	err := root.InitMongoClient()
	require.Nil(t, err)

	mongoClient := root.MongoClient

	require.Equal(t, mongoClient.Config.Enviroment, db.ENV_TEST) //check for right db config
	collection := mongoClient.Database.Collection(collectionName)

	clearMongoTestCollection(t, collection, ctx)

	return mongoClient, collection
}

//removes all documents from the specific collection
func clearMongoTestCollection(t *testing.T, collection *mongo.Collection, ctx context.Context) {
	err := collection.Drop(ctx)
	require.Nil(t, err, "droping of collection failed")
}

func getTestScoreConfig() *db.ScoreParams {

	config := &db.ScoreParams{
		FollowingK:      0.4,
		FollowingW:      1.0,
		FollowersK:      0.1,
		FollowersW:      1.0,
		GistsK:          0.00002,
		GistsW:          1.0,
		IssuesK:         0.0043,
		IssuesW:         1.0,
		OrganizationsK:  0.02,
		OrganizationsW:  1.0,
		ProjectsK:       0.00000167,
		ProjectsW:       1.0,
		PullRequestsK:   0.7,
		PullRequestsW:   1.0,
		ContributionsK:  0.01,
		ContributionsW:  1.0,
		StarredK:        1.4,
		StarredW:        1.0,
		Watchingk:       1.1,
		WatchingW:       1.0,
		CommitCommentsK: 0.1,
		CommitCommentsW: 1.0,
		GistCommentsK:   0.00024,
		GistCommentsW:   1.0,
		IssueCommentsK:  0.04,
		IssueCommentsW:  1.0,
		ReposK:          0.6,
		ReposW:          1.0,
		CommitFrequenzK: 1.2,
		CommitFrequenzW: 1.0,
		StargazersK:     0.5,
		StargazersW:     1.0,
		ForksK:          0.3,
		ForksW:          1.0,
	}

	return config
}

func getTestUser() *db.User {

	user := &db.User{
		Login:                     "testuser",
		Following:                 55,
		Followers:                 23,
		Gists:                     2,
		Issues:                    42,
		Organizations:             10,
		Projects:                  1,
		PullRequests:              67,
		RepositoriesContributedTo: 3,
		StarredRepositories:       102,
		Watching:                  57,
		CommitComments:            23,
		GistComments:              1,
		IssueComments:             98,
		Repositories:              44,
		CommitFrequenz:            44.56,
		Stargazers:                32,
		Forks:                     15,
	}

	return user
}
