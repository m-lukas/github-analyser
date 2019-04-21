package graphql

import (
	"errors"
	"time"
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

func GetCommitData(userName string) (*CommitData, error) {

	if userName == "" {
		return nil, errors.New("username must not be empty!")
	}

	var rawData CommitDataRaw

	err := query(userName, "./graphql/queries/commit.gql", &rawData)
	if err != nil {
		return nil, err
	}

	convertedData := convertCommitData(&rawData)

	return convertedData, nil

}

func convertCommitData(rawData *CommitDataRaw) *CommitData {

	commitFrequenz := GetCommitFrequenz(rawData)

	convertedCommitData := &CommitData{
		CommitFrequenz: commitFrequenz,
	}

	return convertedCommitData

}

func GetCommitFrequenz(rawCommitData *CommitDataRaw) float64 {

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
