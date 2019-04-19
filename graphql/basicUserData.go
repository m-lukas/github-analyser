package graphql

import (
	"context"
	"time"
)

type BasicUserData struct {
	RepositoryOwner struct {
		Login      string
		Email      string
		Bio        string
		Name       string
		Company    string
		Location   string
		AvatarURL  string
		WebsiteURL string
		Followers  struct {
			TotalCount int
		}
		Following struct {
			TotalCount int
		}
		Repositories struct {
			TotalCount int
		}
		Organizations struct {
			TotalCount int
		}
	}
}

func GetBasicUserData(userName string) (*BasicUserData, error) {

	client := NewClient("https://api.github.com/graphql", nil)

	query, err := readQuery("./graphql/queries/userData.gql")
	if err != nil {
		return nil, err
	}

	request := NewRequest(query)
	request.Var("name", userName)

	var basicUserData BasicUserData

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Run(ctx, request, &basicUserData)
	if err != nil {
		return nil, err
	}

	return &basicUserData, nil

}
