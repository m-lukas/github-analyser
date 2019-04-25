package db

import "time"

type User struct {
	Login                     string    `bson:"login"`
	Name                      string    `bson:"name"`
	Email                     string    `bson:"email"`
	Bio                       string    `bson:"bio"`
	Company                   string    `bson:"company"`
	Location                  string    `bson:"location"`
	AvatarURL                 string    `bson:"avatar_url"`
	WebsiteURL                string    `bson:"website_url"`
	GithubJoinDate            time.Time `bson:"github_join_date"`
	IsCampusExpert            bool      `bson:"is_campus_expert"`
	IsDeveloperProgramMember  bool      `bson:"is_developer_program"`
	IsEmployee                bool      `bson:"is_employee"`
	Following                 int       `bson:"following"`
	Followers                 int       `bson:"followers"`
	Gists                     int       `bson:"gists"`
	Issues                    int       `bson:"issues"`
	Organizations             int       `bson:"organizations"`
	Projects                  int       `bson:"projects"`
	PullRequests              int       `bson:"pull_requests"`
	RepositoriesContributedTo int       `bson:"repos_contributed_to"`
	StarredRepositories       int       `bson:"starred_repos"`
	Watching                  int       `bson:"watching"`
	CommitComments            int       `bson:"commit_comments"`
	GistComments              int       `bson:"gist_comments"`
	IssueComments             int       `bson:"issue_comments"`
	Repositories              int       `bson:"repos"`
	CommitFrequenz            float64   `bson:"commit_frequenz"`
	Stargazers                int       `bson:"stargazers"`
	Forks                     int       `bson:"forks"`
	ActivityScore             float64   `bson:"activity_score"`
	PopularityScorefloat64    float64   `bson:"popularity_score"`
	UpdatedAt                 time.Time `bson:"updated_at"`
}

type ScoreParams struct {
	FollowingK      float64
	FollowersK      float64
	GistsK          float64
	IssuesK         float64
	OrganizationsK  float64
	ProjectsK       float64
	PullRequestsK   float64
	ContributionsK  float64
	StarredK        float64
	Watchingk       float64
	CommitCommentsK float64
	GistCommentsK   float64
	IssueCommentsK  float64
	ReposK          float64
	CommitFrequenzK float64
	StargazersK     float64
	ForksK          float64
}
