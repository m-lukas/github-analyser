package db

import "time"

//User schema
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
	Scores                    *Scores   `bson:"scores"`
	ActivityScore             float64   `bson:"activity_score"`
	PopularityScore           float64   `bson:"popularity_score"`
	UpdatedAt                 time.Time `bson:"updated_at"`
}

//ScoreParams schema
type ScoreParams struct {
	FollowingK      float64
	FollowingW      float64
	FollowersK      float64
	FollowersW      float64
	GistsK          float64
	GistsW          float64
	IssuesK         float64
	IssuesW         float64
	OrganizationsK  float64
	OrganizationsW  float64
	ProjectsK       float64
	ProjectsW       float64
	PullRequestsK   float64
	PullRequestsW   float64
	ContributionsK  float64
	ContributionsW  float64
	StarredK        float64
	StarredW        float64
	Watchingk       float64
	WatchingW       float64
	CommitCommentsK float64
	CommitCommentsW float64
	GistCommentsK   float64
	GistCommentsW   float64
	IssueCommentsK  float64
	IssueCommentsW  float64
	ReposK          float64
	ReposW          float64
	CommitFrequenzK float64
	CommitFrequenzW float64
	StargazersK     float64
	StargazersW     float64
	ForksK          float64
	ForksW          float64
}

//Scores schema
type Scores struct {
	FollowingScore      float64
	FollowersScore      float64
	GistsScore          float64
	IssuesScore         float64
	OrganizationsScore  float64
	ProjectsScore       float64
	PullRequestsScore   float64
	ContributionsScore  float64
	StarredScore        float64
	WatchingScore       float64
	CommitCommentsScore float64
	GistCommentsScore   float64
	IssueCommentsScore  float64
	ReposScore          float64
	CommitFrequenzScore float64
	StargazersScore     float64
	ForksScore          float64
}
