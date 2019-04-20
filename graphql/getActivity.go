package graphql

import (
	"context"
	"errors"
	"time"
)

type ActivityRaw struct {
	RateLimit struct {
		Cost      int
		Remaining int
	}
	RepositoryOwner struct {
		Following struct {
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
					CreatedAt time.Time
					Ref       struct {
						Target struct {
							History struct {
								Edges []struct {
									Node struct {
										CommittedDate time.Time
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

type Activity struct {
	Following               int
	Gists                   int
	Issues                  int
	Organizations           int
	Projects                int
	PullRequests            int
	RepositoryContributedTo int
	StarredRepositories     int
	Watching                int
	CommitComments          int
	GistComments            int
	IssueComments           int
	Repositories            int
	CommitFrequenz          float64
}

func GetActivity(userName string) (*Activity, error) {

	if userName == "" {
		return nil, errors.New("username must not be empty!")
	}

	activityData, err := queryActivity(userName)
	if err != nil {
		return nil, err
	}

	convertedActivity := convertActivity(activityData)

	return convertedActivity, nil

}

func queryActivity(userName string) (*ActivityRaw, error) {

	client := NewClient("https://api.github.com/graphql", nil)

	query, err := readQuery("./graphql/queries/activity.gql")
	if err != nil {
		return nil, err
	}

	request := NewRequest(query)
	request.Var("name", userName)

	var activityData ActivityRaw

	ctx, cancel := context.WithTimeout(context.Background(), 12*time.Second)
	defer cancel()
	err = client.Run(ctx, request, &activityData)
	if err != nil {
		return nil, err
	}

	return &activityData, nil

}

func convertActivity(activityData *ActivityRaw) *Activity {

	data := activityData.RepositoryOwner
	commitFrequenz := GetCommitFrequenz(activityData)

	convertedActivity := &Activity{
		Following:               data.Following.TotalCount,
		Gists:                   data.Gists.TotalCount,
		Issues:                  data.Issues.TotalCount,
		Organizations:           data.Organizations.TotalCount,
		Projects:                data.Projects.TotalCount,
		PullRequests:            data.PullRequests.TotalCount,
		RepositoryContributedTo: data.RepositoriesContributedTo.TotalCount,
		StarredRepositories:     data.StarredRepositories.TotalCount,
		Watching:                data.Watching.TotalCount,
		CommitComments:          data.CommitComments.TotalCount,
		GistComments:            data.GistComments.TotalCount,
		IssueComments:           data.IssueComments.TotalCount,
		Repositories:            data.Repositories.TotalCount,
		CommitFrequenz:          commitFrequenz,
	}

	return convertedActivity

}

func GetCommitFrequenz(activityData *ActivityRaw) float64 {

	var commitDates dateSlice
	var commitTimeDifferences []float64

	repositorySlice := activityData.RepositoryOwner.Repositories.Edges
	for _, repo := range repositorySlice {

		commitSlice := repo.Node.Ref.Target.History.Edges

		for _, commit := range commitSlice {

			commitDate := commit.Node.CommittedDate
			commitDates = append(commitDates, commitDate)

		}

	}

	commitDates = sortDatesAsc(commitDates)

	var maxCheckedItems int

	for index := len(commitDates) - 1; index > 0; index-- {

		if maxCheckedItems >= 100 {
			break
		} else {
			maxCheckedItems++
		}

		commitDate := commitDates[index]
		commitDateBefore := commitDates[index-1]

		timeDiff := commitDate.Sub(commitDateBefore)
		hoursDiff := timeDiff.Hours()

		commitTimeDifferences = append(commitTimeDifferences, hoursDiff)

	}

	frequenz := avg(commitTimeDifferences)

	return frequenz

}
