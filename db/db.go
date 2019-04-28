package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var dbRoot *DatabaseRoot

type DatabaseRoot struct {
	MongoClient *mongo.Client
	RedisClient *redis.Client
	Config      *DatabaseConfig
	ScoreConfig *ScoreParams
}

type DatabaseConfig struct {
	MongoDatabaseName string
	MongoURI          string
}

func Init() error {
	var err error

	dbRoot = &DatabaseRoot{
		Config: defaultDatabaseConfig(),
	}

	err = dbRoot.initMongoClient()
	if err != nil {
		return err
	}

	err = dbRoot.initRedisClient()
	if err != nil {
		return err
	}

	err = dbRoot.initScoreConfig()
	if err != nil {
		return err
	}

	return nil
}

func GetMongo() (*mongo.Database, error) {

	root, err := getRoot()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = root.MongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		err := root.initMongoClient()
		if err != nil {
			return nil, err
		}
	}

	client := root.MongoClient
	config := root.Config
	db := client.Database(config.MongoDatabaseName)
	return db, nil
}

func GetRedis() (*redis.Client, error) {

	root, err := getRoot()
	if err != nil {
		return nil, err
	}

	client := root.RedisClient
	_, err = client.Ping().Result()
	if err != nil {

		err = root.initRedisClient()
		if err != nil {
			return nil, err
		}

	}

	return client, nil
}

func GetScoreConfig() (*ScoreParams, error) {

	root, err := getRoot()
	if err != nil {
		return nil, err
	}

	if root.ScoreConfig == nil {
		err := root.initScoreConfig()
		if err != nil {
			return nil, err
		}
	}

	return root.ScoreConfig, nil

}

//TODO:
func ReinitializeScoreConfig() error {
	root, err := getRoot()
	if err != nil {
		return err
	}

	err = root.initScoreConfig()
	if err != nil {
		return err
	}

	return nil
}

/*
	Initializes the mongoDB Client to access databases and collections.
*/
func (root *DatabaseRoot) initMongoClient() error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(root.Config.MongoURI))
	if err != nil {
		return err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	root.MongoClient = client
	log.Println("Initialized mongo client!")

	return nil
}

func (root *DatabaseRoot) initRedisClient() error {

	client := redis.NewClient(&redis.Options{
		Addr:     getRedisURI(),
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return err
	}

	root.RedisClient = client
	log.Println("Initialized redis client!")

	return nil

}

func (root *DatabaseRoot) initScoreConfig() error {

	if root.RedisClient == nil {
		return errors.New("redis client not initialized!")
	}

	redisClient, err := GetRedis()
	if err != nil {
		return err
	}

	followingK := GetScoreParam(redisClient, "following", "k")
	followingW := GetScoreParam(redisClient, "following", "w")
	followersK := GetScoreParam(redisClient, "followers", "k")
	followersW := GetScoreParam(redisClient, "followers", "w")
	gistsK := GetScoreParam(redisClient, "gists", "k")
	gistsW := GetScoreParam(redisClient, "gists", "w")
	issuesK := GetScoreParam(redisClient, "issues", "k")
	issuesW := GetScoreParam(redisClient, "issues", "w")
	organizationsK := GetScoreParam(redisClient, "organizations", "k")
	organizationsW := GetScoreParam(redisClient, "organizations", "w")
	projectsK := GetScoreParam(redisClient, "projects", "k")
	projectsW := GetScoreParam(redisClient, "projects", "w")
	pullRequestsK := GetScoreParam(redisClient, "pull_requests", "k")
	pullRequestsW := GetScoreParam(redisClient, "pull_requests", "w")
	contributionsK := GetScoreParam(redisClient, "contributions", "k")
	contributionsW := GetScoreParam(redisClient, "contributions", "k")
	starredK := GetScoreParam(redisClient, "starred", "k")
	starredW := GetScoreParam(redisClient, "starred", "w")
	watchingK := GetScoreParam(redisClient, "watching", "k")
	watchingW := GetScoreParam(redisClient, "watching", "w")
	commitCommentsK := GetScoreParam(redisClient, "commit_comments", "k")
	commitCommentsW := GetScoreParam(redisClient, "commit_comments", "w")
	gistCommentsK := GetScoreParam(redisClient, "gist_comments", "k")
	gistCommentsW := GetScoreParam(redisClient, "gist_comments", "w")
	issueCommentsK := GetScoreParam(redisClient, "issue_comments", "k")
	issueCommentsW := GetScoreParam(redisClient, "issue_comments", "w")
	reposK := GetScoreParam(redisClient, "repos", "k")
	reposW := GetScoreParam(redisClient, "repos", "w")
	commitFrequenzK := GetScoreParam(redisClient, "commit_frequenz", "k")
	commitFrequenzW := GetScoreParam(redisClient, "commit_frequenz", "w")
	stargazersK := GetScoreParam(redisClient, "stargazers", "k")
	stargazersW := GetScoreParam(redisClient, "stargazers", "w")
	forksK := GetScoreParam(redisClient, "forks", "k")
	forksW := GetScoreParam(redisClient, "forks", "w")

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

func defaultDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		MongoDatabaseName: "core",
		MongoURI:          getMongoURI(),
	}
}

func getRoot() (*DatabaseRoot, error) {

	var err error

	err = checkDbRoot()
	if err != nil {
		return nil, err
	}

	return dbRoot, nil

}

func checkDbRoot() error {

	if dbRoot == nil {
		err := Init()
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}

/*
	Returns configurated URI string for MongoDB.
*/
func getMongoURI() (uri string) {
	dbHost := os.Getenv("MONGO_HOST")
	dbName := os.Getenv("MONGO_NAME")
	dbUser := os.Getenv("MONGO_USER")
	dbPass := os.Getenv("MONGO_PASS")

	return fmt.Sprintf(`mongodb://%s:%s@%s/%s`, dbUser, dbPass, dbHost, dbName)
}

func getRedisURI() (uri string) {
	dbHost := os.Getenv("REDIS_HOST")
	return dbHost
}
