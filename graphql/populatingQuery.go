package graphql

import (
	"errors"
)

type PopulatingDataRaw struct {
	RateLimit       rateLimit
	RepositoryOwner struct {
		Login     string
		Following struct {
			Edges []struct {
				Node struct {
					Login     string
					Following struct {
						Edges []struct {
							Node struct {
								Login string
							}
						}
					}
					Followers struct {
						Edges []struct {
							Node struct {
								Login string
							}
						}
					}
				}
			}
		}
		Followers struct {
			Edges []struct {
				Node struct {
					Login     string
					Following struct {
						Edges []struct {
							Node struct {
								Login string
							}
						}
					}
					Followers struct {
						Edges []struct {
							Node struct {
								Login string
							}
						}
					}
				}
			}
		}
	}
}

func GetPopulatingData(rootUser string) ([]string, error) {

	if rootUser == "" {
		return nil, errors.New("username must not be empty!")
	}

	var rawData PopulatingDataRaw

	err := query(rootUser, "./graphql/queries/populating.gql", &rawData)
	if err != nil {
		return nil, err
	}

	loginList := GetLoginList(&rawData)

	return loginList, nil

}

func GetLoginList(rawData *PopulatingDataRaw) []string {

	var loginList []string

	repositoryOwner := rawData.RepositoryOwner

	topFollowingSlice := repositoryOwner.Following.Edges
	topFollowersSlice := repositoryOwner.Followers.Edges

	for _, edge := range topFollowingSlice {

		user := edge.Node
		login := user.Login

		loginList = append(loginList, login)

		bottomFollowingSlice := user.Following.Edges
		bottomFollowersSlice := user.Followers.Edges

		for _, edge := range bottomFollowingSlice {

			user := edge.Node
			login := user.Login

			loginList = append(loginList, login)

		}

		for _, edge := range bottomFollowersSlice {

			user := edge.Node
			login := user.Login

			loginList = append(loginList, login)

		}
	}

	for _, edge := range topFollowersSlice {

		user := edge.Node
		login := user.Login

		loginList = append(loginList, login)

		bottomFollowingSlice := user.Following.Edges
		bottomFollowersSlice := user.Followers.Edges

		for _, edge := range bottomFollowingSlice {

			user := edge.Node
			login := user.Login

			loginList = append(loginList, login)

		}

		for _, edge := range bottomFollowersSlice {

			user := edge.Node
			login := user.Login

			loginList = append(loginList, login)

		}
	}

	uniqueList := removeDuplicates(loginList)

	return uniqueList

}
