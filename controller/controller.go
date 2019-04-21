package controller

import (
	"fmt"
	"time"

	"github.com/m-lukas/github-analyser/db"
	"github.com/m-lukas/github-analyser/graphql"
)

func QueryUser(userName string) (*db.User, error) {

	generalChannel := make(chan graphql.GeneralDataResponse)
	commitChannel := make(chan graphql.CommitDataResponse)
	popularityChannel := make(chan graphql.PopularityDataResponse)

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
			fmt.Println(".")
			break
		}

		if generalData != nil && commitData != nil && popularityData != nil {
			break
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
