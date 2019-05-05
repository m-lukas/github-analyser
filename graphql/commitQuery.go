package graphql

import (
	"errors"
	"time"

	"github.com/m-lukas/github-analyser/util"
)

type CommitDataRaw struct {
	RateLimit       rateLimit
	RepositoryOwner struct {
		Repositories struct {
			Edges []struct {
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

type CommitData struct {
	CommitFrequenz float64
}

type CommitDataResponse struct {
	Data  *CommitData
	Error error
}

func GetCommitData(userName string) CommitDataResponse {

	if userName == "" {
		return CommitDataResponse{Data: nil, Error: errors.New("username must not be empty!")}
	}

	var rawData CommitDataRaw

	err := query(userName, "./graphql/queries/commit.gql", &rawData)
	if err != nil {
		return CommitDataResponse{Data: nil, Error: err}
	}

	convertedData := convertCommitData(&rawData)

	return CommitDataResponse{Data: convertedData, Error: nil}

}

func convertCommitData(rawData *CommitDataRaw) *CommitData {

	commitFrequenz := GetCommitFrequenz(rawData, time.Now())

	convertedCommitData := &CommitData{
		CommitFrequenz: commitFrequenz,
	}

	return convertedCommitData

}

func GetCommitFrequenz(rawCommitData *CommitDataRaw, today time.Time) float64 {

	var commitDates dateSlice
	var commitTimeDifferences []float64

	repositorySlice := rawCommitData.RepositoryOwner.Repositories.Edges
	for _, repo := range repositorySlice {

		commitSlice := repo.Node.Ref.Target.History.Edges

		for _, commit := range commitSlice {

			commitDate := commit.Node.CommittedDate
			commitDates = append(commitDates, commitDate)

		}
	}

	if len(commitDates) < 1 {
		return -1.0
	}

	commitDates = sortDatesAsc(commitDates)
	var maxCheckedItems int

	recentCommitDate := commitDates[len(commitDates)-1]
	hoursDiff := getHoursDifference(today, recentCommitDate)
	commitTimeDifferences = append(commitTimeDifferences, hoursDiff)

	for index := len(commitDates) - 1; index > 0; index-- {

		if maxCheckedItems >= 100 {
			break
		} else {
			maxCheckedItems++
		}

		commitDate := commitDates[index]
		commitDateBefore := commitDates[index-1]

		hoursDiff = getHoursDifference(commitDate, commitDateBefore)
		commitTimeDifferences = append(commitTimeDifferences, hoursDiff)
	}

	frequenz := util.Avg(commitTimeDifferences)

	return frequenz

}

func getHoursDifference(endDate time.Time, startDate time.Time) float64 {
	timeDiff := endDate.Sub(startDate)
	hoursDiff := timeDiff.Hours()
	return hoursDiff
}
