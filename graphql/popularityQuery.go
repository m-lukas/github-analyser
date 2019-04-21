package graphql

import (
	"errors"
	"strings"
)

type PopularityRaw struct {
	RateLimit       rateLimit
	RepositoryOwner struct {
		Repositories struct {
			Edges []struct {
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

type Popularity struct {
	Stargazers int
	Forks      int
}

func GetPopularity(userName string) (*Popularity, error) {

	if userName == "" {
		return nil, errors.New("username must not be empty!")
	}

	var popularityData PopularityRaw

	err := query(userName, "./graphql/queries/popularity.gql", &popularityData)
	if err != nil {
		return nil, err
	}

	convertedPopularity := convertPopularity(&popularityData, userName)

	return convertedPopularity, nil

}

func convertPopularity(popularityData *PopularityRaw, userName string) *Popularity {

	stargazers, forks := CalcStargazersAndForks(popularityData, userName)

	convertedPopularity := &Popularity{
		Stargazers: stargazers,
		Forks:      forks,
	}

	return convertedPopularity

}

func CalcStargazersAndForks(popularityData *PopularityRaw, userName string) (int, int) {

	var stargazersSum int
	var forksSum int

	repoCounter := 0
	maxRepoNum := 25

	repositorySlice := popularityData.RepositoryOwner.Repositories.Edges

	for _, repo := range repositorySlice {

		repoNode := repo.Node
		owner := strings.Split(repoNode.NameWithOwner, "/")[0]

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
