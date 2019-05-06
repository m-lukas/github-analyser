package controller

import (
	"errors"
	"fmt"
	"time"

	"github.com/m-lukas/github-analyser/db"
	"github.com/m-lukas/github-analyser/graphql"
)

func fetchUser(userName string) (*db.User, error) {

	generalChannel := make(chan graphql.GeneralDataResponse)
	commitChannel := make(chan graphql.CommitDataResponse)

	var startTime = time.Now()

	go func(userName string) {
		generalChannel <- graphql.GetGeneralData(userName)
	}(userName)

	go func(userName string) {
		commitChannel <- graphql.GetCommitData(userName)
	}(userName)

	var generalData *graphql.GeneralData
	var commitData *graphql.CommitData

	for {
		select {
		case resp := <-generalChannel:

			if resp.Error != nil {
				fmt.Printf("General: %v", resp.Error)
				return nil, resp.Error
			} else {
				generalData = resp.Data
			}

		case resp := <-commitChannel:

			if resp.Error != nil {
				fmt.Printf("Commit: %v", resp.Error)
				return nil, resp.Error
			} else {
				commitData = resp.Data
			}

		case <-time.After(50 * time.Millisecond):
			break
		}

		if generalData != nil && commitData != nil {
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
		Stargazers:                generalData.Stargazers,
		Forks:                     generalData.Forks,
		UpdatedAt:                 time.Now(),
	}

	return user, nil

}
