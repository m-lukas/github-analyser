package metrix

import (
	"fmt"
	"time"

	"github.com/m-lukas/github-analyser/controller"
	"github.com/m-lukas/github-analyser/db"
	"github.com/m-lukas/github-analyser/logger"
	"github.com/m-lukas/github-analyser/util"
)

const (
	prefix = "METRIX |"

	TYPE_FOLLOWING      = "TYPE_FOLLOWING"
	TYPE_FOLLOWERS      = "TYPE_FOLLOWERS"
	TYPE_GISTS          = "TYPE_GISTS"
	TYPE_ISSUES         = "TYPE_ISSUES"
	TYPE_ORGANIZATIONS  = "TYPE_ORGANIZATIONS"
	TYPE_PROJECTS       = "TYPE_PROJECTS"
	TYPE_PULLREQUESTS   = "TYPE_PULLREQUESTS"
	TYPE_CONTRIBUTIONS  = "TYPE_CONTRIBUTIONS"
	TYPE_STARRED        = "TYPE_STARRED"
	TYPE_WATCHING       = "TYPE_WATCHING"
	TYPE_COMMITCOMMENTS = "TYPE_COMMITCOMMENTS"
	TYPE_GISTCOMMENTS   = "TYPE_GISTCOMMENTS"
	TYPE_ISSUECOMMENTS  = "TYPE_ISSUECOMMENTS"
	TYPE_REPOS          = "TYPE_REPOS"
	TYPE_COMMITFREQUENZ = "TYPE_COMMITFREQUENZ"
	TYPE_STARGAZERS     = "TYPE_STARGAZERS"
	TYPE_FORKS          = "TYPE_FORKS"
)

func CalcScoreParams() error {

	startTime := time.Now()
	logger.Info(fmt.Sprintf("%s Start time: %s\n", prefix, util.FormatDuration(time.Since(startTime))))

	inputFiles := []string{"./metrix/input/sindresorhus.txt"}
	userArray, err := populateData(inputFiles)
	if err != nil {
		return err
	}

	var dbPairs = make(map[string]interface{}, 0)

	for fieldType, redisKey := range fieldTypes() {

		k := calcK(userArray, fieldType)
		if k <= 0 {
			continue //ignore flawed values
		}

		dbPairs[redisKey] = k

	}

	err = controller.SaveConfigValues(dbPairs, "k")
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("%s Successfully saved user data!\n", prefix))

	err = db.ReinitializeScoreConfig()
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("%s Reinitialized score config!\n", prefix))

	err = controller.UpdateAllScores()
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("%s Finished metrix initialization!\n", prefix))

	logger.Info(fmt.Sprintf("%s End time: %s\n", prefix, util.FormatDuration(time.Since(startTime))))

	return nil
}

func fieldTypes() map[string]string {
	return map[string]string{
		TYPE_FOLLOWING:      "following",
		TYPE_FOLLOWERS:      "followers",
		TYPE_GISTS:          "gists",
		TYPE_ISSUES:         "issues",
		TYPE_ORGANIZATIONS:  "organizations",
		TYPE_PROJECTS:       "projects",
		TYPE_PULLREQUESTS:   "pull_requests",
		TYPE_CONTRIBUTIONS:  "contributions",
		TYPE_STARRED:        "starred",
		TYPE_WATCHING:       "watching",
		TYPE_COMMITCOMMENTS: "commit_comments",
		TYPE_GISTCOMMENTS:   "gist_comments",
		TYPE_ISSUECOMMENTS:  "issue_comments",
		TYPE_REPOS:          "repos",
		TYPE_COMMITFREQUENZ: "commit_frequenz",
		TYPE_STARGAZERS:     "stargazers",
		TYPE_FORKS:          "forks",
	}
}
