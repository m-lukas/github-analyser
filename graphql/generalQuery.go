package graphql

import (
	"errors"
	"strings"
	"time"
)

type GeneralDataRaw struct {
	RateLimit       rateLimit
	RepositoryOwner struct {
		Login                    string
		Name                     string
		Email                    string
		Bio                      string
		Company                  string
		Location                 string
		AvatarURL                string
		WebsiteURL               string
		CreatedAt                time.Time
		IsCampusExpert           bool
		IsDeveloperProgramMember bool
		IsEmployee               bool
		Following                struct {
			TotalCount int
		}
		Followers struct {
			TotalCount int
		}
		Gists struct {
			TotalCount int
		}
		Issues struct {
			TotalCount int
		}
		Organizations struct {
			TotalCount int
		}
		Projects struct {
			TotalCount int
		}
		PullRequests struct {
			TotalCount int
		}
		RepositoriesContributedTo struct {
			TotalCount int
		}
		StarredRepositories struct {
			TotalCount int
		}
		Watching struct {
			TotalCount int
		}
		CommitComments struct {
			TotalCount int
		}
		GistComments struct {
			TotalCount int
		}
		IssueComments struct {
			TotalCount int
		}
		Repositories struct {
			TotalCount int
			Edges      []struct {
				Node struct {
					NameWithOwner string
					Stargazers    struct {
						TotalCount int
					}
					Forks struct {
						TotalCount int
					}
				}
			}
		}
	}
}

type GeneralData struct {
	Login                     string
	Name                      string
	Email                     string
	Bio                       string
	Company                   string
	Location                  string
	AvatarURL                 string
	WebsiteURL                string
	CreatedAt                 time.Time
	IsCampusExpert            bool
	IsDeveloperProgramMember  bool
	IsEmployee                bool
	Following                 int
	Followers                 int
	Gists                     int
	Issues                    int
	Organizations             int
	Projects                  int
	PullRequests              int
	RepositoriesContributedTo int
	StarredRepositories       int
	Watching                  int
	CommitComments            int
	GistComments              int
	IssueComments             int
	Repositories              int
	Stargazers                int
	Forks                     int
}

type GeneralDataResponse struct {
	Data  *GeneralData
	Error error
}

//GetCommitData fetches and processes the general data of the user
func GetGeneralData(userName string) GeneralDataResponse {

	if userName == "" {
		return GeneralDataResponse{Data: nil, Error: errors.New("username must not be empty!")}
	}

	var rawData GeneralDataRaw

	err := query(userName, "./graphql/queries/general.gql", &rawData)
	if err != nil {
		return GeneralDataResponse{Data: nil, Error: err}
	}

	convertedData := convertGeneralData(&rawData)

	return GeneralDataResponse{Data: convertedData, Error: nil}

}

//convertGeneralData converts raw to normalised data
func convertGeneralData(rawData *GeneralDataRaw) *GeneralData {

	data := rawData.RepositoryOwner
	stargazers, forks := calcStargazersAndForks(rawData, data.Login)

	convertedData := &GeneralData{
		Login:                     data.Login,
		Name:                      data.Name,
		Email:                     data.Email,
		Bio:                       data.Bio,
		Company:                   data.Company,
		Location:                  data.Location,
		AvatarURL:                 data.AvatarURL,
		WebsiteURL:                data.WebsiteURL,
		CreatedAt:                 data.CreatedAt,
		IsCampusExpert:            data.IsCampusExpert,
		IsDeveloperProgramMember:  data.IsDeveloperProgramMember,
		IsEmployee:                data.IsEmployee,
		Following:                 data.Following.TotalCount,
		Followers:                 data.Followers.TotalCount,
		Gists:                     data.Gists.TotalCount,
		Issues:                    data.Issues.TotalCount,
		Organizations:             data.Organizations.TotalCount,
		Projects:                  data.Projects.TotalCount,
		PullRequests:              data.PullRequests.TotalCount,
		RepositoriesContributedTo: data.RepositoriesContributedTo.TotalCount,
		StarredRepositories:       data.StarredRepositories.TotalCount,
		Watching:                  data.Watching.TotalCount,
		CommitComments:            data.CommitComments.TotalCount,
		GistComments:              data.GistComments.TotalCount,
		IssueComments:             data.IssueComments.TotalCount,
		Repositories:              data.Repositories.TotalCount,
		Forks:                     forks,
		Stargazers:                stargazers,
	}

	return convertedData

}

//calcStargazersAndForks loops through the nested data and gets the total count of forks and stargazers
func calcStargazersAndForks(rawData *GeneralDataRaw, userName string) (int, int) {

	var stargazersSum int
	var forksSum int

	repoCounter := 0
	maxRepoNum := 25 //limit repo count to reduce advantage of user with many repositories

	repositorySlice := rawData.RepositoryOwner.Repositories.Edges

	for _, repo := range repositorySlice {

		repoNode := repo.Node
		owner := strings.Split(repoNode.NameWithOwner, "/")[0]

		//check if user is the owner of the repository
		if owner != userName {
			continue
		} else {

			stargazers := repoNode.Stargazers.TotalCount
			forks := repoNode.Forks.TotalCount

			stargazersSum += stargazers
			forksSum += forks

		}

		if repoCounter >= maxRepoNum {
			break
		} else {
			repoCounter++
		}

	}

	return stargazersSum, forksSum

}
