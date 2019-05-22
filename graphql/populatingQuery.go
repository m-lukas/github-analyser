package graphql

import (
	"errors"

	"github.com/m-lukas/github-analyser/util"
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

//GetPopulatingData gets a list of followers and following users by a rootUser login
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

//GetLoginList processes rawData and returns a list of unique login strings
func GetLoginList(rawData *PopulatingDataRaw) []string {

	var loginList []string

	repositoryOwner := rawData.RepositoryOwner

	topFollowingSlice := repositoryOwner.Following.Edges
	topFollowersSlice := repositoryOwner.Followers.Edges

	//loop through top-level following users
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

	//loop through top-level followers
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

	//loop through list and remove duplicates
	uniqueList := util.RemoveDuplicates(loginList)

	return uniqueList

}
