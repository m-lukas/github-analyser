package controller

import (
	"errors"
	"log"
	"time"

	"github.com/m-lukas/github-analyser/db"
	"github.com/m-lukas/github-analyser/graphql"
)

const timoutSeconds = 15

func GetUser(userName string) (*db.User, error) {

	user, err := queryUser(userName)
	if err != nil {
		return nil, err
	}

	if user.Login == "" {
		return nil, errors.New("User does not exist!")
	}

	/*
		config, err := db.GetScoreConfig()
		if err != nil {
			return nil, err
		}

		user.Scores = CalcScores(user, config)

		user.ActivityScore = CalcActivityScore(user.Scores, config)
		user.PopularityScore = CalcPopularityScore(user.Scores, config)
	*/

	go func(user *db.User) {
		err = db.CacheUser(user)
		if err != nil {
			log.Println(err)
		}
	}(user)

	return user, nil

}

func queryUser(userName string) (*db.User, error) {

	generalChannel := make(chan graphql.GeneralDataResponse)
	commitChannel := make(chan graphql.CommitDataResponse)
	popularityChannel := make(chan graphql.PopularityDataResponse)

	var startTime = time.Now()

	go func(userName string) {
		generalChannel <- graphql.GetGeneralData(userName)
	}(userName)

	go func(userName string) {
		commitChannel <- graphql.GetCommitData(userName)
	}(userName)

	go func(userName string) {
		popularityChannel <- graphql.GetPopularity(userName)
	}(userName)

	var generalData *graphql.GeneralData
	var commitData *graphql.CommitData
	var popularityData *graphql.Popularity

	for {
		select {
		case resp := <-generalChannel:

			if resp.Error != nil {
				return nil, resp.Error
			} else {
				generalData = resp.Data
			}

		case resp := <-commitChannel:

			if resp.Error != nil {
				return nil, resp.Error
			} else {
				commitData = resp.Data
			}

		case resp := <-popularityChannel:

			if resp.Error != nil {
				return nil, resp.Error
			} else {
				popularityData = resp.Data
			}

		case <-time.After(50 * time.Millisecond):
			break
		}

		if generalData != nil && commitData != nil && popularityData != nil {
			break
		}

		if time.Since(startTime).Seconds() >= float64(timoutSeconds) {
			return nil, errors.New("Timeout while trying to receive data!")
		}

	}

	user := &db.User{
		Login:                     generalData.Login,
		Name:                      generalData.Name,
		Email:                     generalData.Email,
		Bio:                       generalData.Bio,
		Company:                   generalData.Company,
		Location:                  generalData.Location,
		AvatarURL:                 generalData.AvatarURL,
		WebsiteURL:                generalData.WebsiteURL,
		GithubJoinDate:            generalData.CreatedAt,
		IsCampusExpert:            generalData.IsCampusExpert,
		IsDeveloperProgramMember:  generalData.IsDeveloperProgramMember,
		IsEmployee:                generalData.IsEmployee,
		Following:                 generalData.Following,
		Followers:                 generalData.Followers,
		Gists:                     generalData.Gists,
		Issues:                    generalData.Issues,
		Organizations:             generalData.Organizations,
		Projects:                  generalData.Projects,
		PullRequests:              generalData.PullRequests,
		RepositoriesContributedTo: generalData.RepositoriesContributedTo,
		StarredRepositories:       generalData.StarredRepositories,
		Watching:                  generalData.Watching,
		CommitComments:            generalData.CommitComments,
		GistComments:              generalData.GistComments,
		IssueComments:             generalData.IssueComments,
		Repositories:              generalData.Repositories,
		CommitFrequenz:            commitData.CommitFrequenz,
		Stargazers:                popularityData.Stargazers,
		Forks:                     popularityData.Forks,
		UpdatedAt:                 time.Now(),
	}

	return user, nil

}
