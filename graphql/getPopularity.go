package graphql

import (
	"context"
	"errors"
	"strings"
	"time"
)

type PopularityRaw struct {
	RateLimit struct {
		Cost      int
		Remaining int
	}
	RepositoryOwner struct {
		Followers struct {
			TotalCount int
		}
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
	Followers  int
	Stargazers int
	Forks      int
}

func GetPopularity(userName string) (*Popularity, error) {

	if userName == "" {
		return nil, errors.New("username must not be empty!")
	}

	popularityData, err := queryPopularity(userName)
	if err != nil {
		return nil, err
	}

	convertedPopularity := convertPopularity(popularityData, userName)

	return convertedPopularity, nil

}

func queryPopularity(userName string) (*PopularityRaw, error) {

	client := NewClient("https://api.github.com/graphql", nil)

	query, err := readQuery("./graphql/queries/popularity.gql")
	if err != nil {
		return nil, err
	}

	request := NewRequest(query)
	request.Var("name", userName)

	var popularityData PopularityRaw

	ctx, cancel := context.WithTimeout(context.Background(), 12*time.Second)
	defer cancel()
	err = client.Run(ctx, request, &popularityData)
	if err != nil {
		return nil, err
	}

	return &popularityData, nil

}

func convertPopularity(popularityData *PopularityRaw, userName string) *Popularity {

	data := popularityData.RepositoryOwner
	stargazers, forks := CalcStargazersAndForks(popularityData, userName)

	convertedPopularity := &Popularity{
		Followers:  data.Followers.TotalCount,
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
