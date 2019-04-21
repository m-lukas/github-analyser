package db

import "time"

type User struct {
	Login                     string
	Name                      string
	Email                     string
	Bio                       string
	Company                   string
	Location                  string
	AvatarURL                 string
	WebsiteURL                string
	GithubJoinDate            time.Time
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
	CommitFrequenz            float64
	Stargazers                int
	Forks                     int
	UpdatedAt                 time.Time
}
