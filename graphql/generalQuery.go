package graphql

import (
	"errors"
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
}

type GeneralDataResponse struct {
	Data  *GeneralData
	Error error
}

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

func convertGeneralData(rawData *GeneralDataRaw) *GeneralData {

	data := rawData.RepositoryOwner

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
	}

	return convertedData

}
