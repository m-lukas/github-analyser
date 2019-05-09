package app

import (
	"context"
	"testing"

	"github.com/m-lukas/github-analyser/db"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
)

func setupMongoTest(t *testing.T, root *db.DatabaseRoot, collectionName string, ctx context.Context) (*db.MongoClient, *mongo.Collection) {

	err := root.InitMongoClient()
	require.Nil(t, err)

	mongoClient := root.MongoClient

	require.Equal(t, mongoClient.Config.Enviroment, db.ENV_TEST) //check for right db config
	collection := mongoClient.Database.Collection(collectionName)

	clearMongoTestCollection(t, collection, ctx)

	return mongoClient, collection
}

func clearMongoTestCollection(t *testing.T, collection *mongo.Collection, ctx context.Context) {
	err := collection.Drop(ctx)
	require.Nil(t, err, "droping of collection failed")
}

func setTestScoreConfig(root *db.DatabaseRoot) {

	root.ScoreConfig = &db.ScoreParams{
		FollowingK:      1.0,
		FollowingW:      1.0,
		FollowersK:      1.0,
		FollowersW:      1.0,
		GistsK:          1.0,
		GistsW:          1.0,
		IssuesK:         1.0,
		IssuesW:         1.0,
		OrganizationsK:  1.0,
		OrganizationsW:  1.0,
		ProjectsK:       1.0,
		ProjectsW:       1.0,
		PullRequestsK:   1.0,
		PullRequestsW:   1.0,
		ContributionsK:  1.0,
		ContributionsW:  1.0,
		StarredK:        1.0,
		StarredW:        1.0,
		Watchingk:       1.0,
		WatchingW:       1.0,
		CommitCommentsK: 1.0,
		CommitCommentsW: 1.0,
		GistCommentsK:   1.0,
		GistCommentsW:   1.0,
		IssueCommentsK:  1.0,
		IssueCommentsW:  1.0,
		ReposK:          1.0,
		ReposW:          1.0,
		CommitFrequenzK: 1.0,
		CommitFrequenzW: 1.0,
		StargazersK:     1.0,
		StargazersW:     1.0,
		ForksK:          1.0,
		ForksW:          1.0,
	}

}
