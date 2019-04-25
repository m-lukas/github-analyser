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

	fmt.Println(dbRoot.ScoreConfig)

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
	followersK := GetScoreParam(redisClient, "followers", "k")
	gistsK := GetScoreParam(redisClient, "gists", "k")
	issuesK := GetScoreParam(redisClient, "issues", "k")
	organizationsK := GetScoreParam(redisClient, "organizations", "k")
	projectsK := GetScoreParam(redisClient, "projects", "k")
	pullRequestsK := GetScoreParam(redisClient, "pullr_equests", "k")
	contributionsK := GetScoreParam(redisClient, "contributions", "k")
	starredK := GetScoreParam(redisClient, "starred", "k")
	watchingk := GetScoreParam(redisClient, "watching", "k")
	commitCommentsK := GetScoreParam(redisClient, "commit_comments", "k")
	gistCommentsK := GetScoreParam(redisClient, "gist_comments", "k")
	issueCommentsK := GetScoreParam(redisClient, "issue_comments", "k")
	reposK := GetScoreParam(redisClient, "repos", "k")
	commitFrequenzK := GetScoreParam(redisClient, "commit_frequenz", "k")
	stargazersK := GetScoreParam(redisClient, "stargazers", "k")
	forksK := GetScoreParam(redisClient, "forks", "k")

	scoreConfig := &ScoreParams{
		FollowingK:      followingK,
		FollowersK:      followersK,
		GistsK:          gistsK,
		IssuesK:         issuesK,
		OrganizationsK:  organizationsK,
		ProjectsK:       projectsK,
		PullRequestsK:   pullRequestsK,
		ContributionsK:  contributionsK,
		StarredK:        starredK,
		Watchingk:       watchingk,
		CommitCommentsK: commitCommentsK,
		GistCommentsK:   gistCommentsK,
		IssueCommentsK:  issueCommentsK,
		ReposK:          reposK,
		CommitFrequenzK: commitFrequenzK,
		StargazersK:     stargazersK,
		ForksK:          forksK,
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
