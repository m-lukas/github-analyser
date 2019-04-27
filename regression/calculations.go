package regression

import (
	"math"
	"sort"

	"github.com/m-lukas/github-analyser/db"
)

func calcK(users []*db.User, fieldType string) float64 {

	var valueArray []float64

	for _, user := range users {

		value := getField(user, fieldType)
		valueArray = append(valueArray, value)

	}

	midValue := getNearestToAvgValue(valueArray)
	if midValue == 0.0 {

		//TODO: Calculate k by highest score

		return 1.0 //ERROR
	}

	var avgScore = 50.0
	k := calcKFromY(avgScore, midValue)

	return k

}

func getField(user *db.User, fieldType string) float64 {

	switch fieldType {
	case TYPE_FOLLOWING:
		return float64(user.Following)
	case TYPE_FOLLOWERS:
		return float64(user.Followers)
	case TYPE_GISTS:
		return float64(user.Gists)
	case TYPE_ISSUES:
		return float64(user.Issues)
	case TYPE_ORGANIZATIONS:
		return float64(user.Organizations)
	case TYPE_PROJECTS:
		return float64(user.Projects)
	case TYPE_PULLREQUESTS:
		return float64(user.PullRequests)
	case TYPE_CONTRIBUTIONS:
		return float64(user.RepositoriesContributedTo)
	case TYPE_STARRED:
		return float64(user.StarredRepositories)
	case TYPE_WATCHING:
		return float64(user.Watching)
	case TYPE_COMMITCOMMENTS:
		return float64(user.CommitComments)
	case TYPE_GISTCOMMENTS:
		return float64(user.GistComments)
	case TYPE_ISSUECOMMENTS:
		return float64(user.IssueComments)
	case TYPE_REPOS:
		return float64(user.Repositories)
	case TYPE_COMMITFREQUENZ:
		return float64(user.CommitFrequenz)
	case TYPE_STARGAZERS:
		return float64(user.Stargazers)
	case TYPE_FORKS:
		return float64(user.Forks)
	default:
		return 0.0
	}

}

func getNearestToAvgValue(valueArray []float64) float64 {

	sort.Float64s(valueArray)
	average := avg(valueArray)

	midIndex := int(math.Round(float64(len(valueArray) / 2)))
	midValue := valueArray[midIndex]

	if midValue == average {
		return midValue
	}

	if midValue < average {

		//ARRAY TOO SHORT!!!!!!!
		for i := midIndex + 1; i < len(valueArray)-1; i++ {
			lastVal := valueArray[i-1]
			currentVal := valueArray[i]
			nextVal := valueArray[i+1]

			distanceFromLast := distanceToNumber(lastVal, average)
			distanceFromCurrent := distanceToNumber(currentVal, average)
			distanceFromNext := distanceToNumber(nextVal, average)

			if distanceFromCurrent > distanceFromLast {
				return lastVal
			}

			if distanceFromCurrent < distanceFromNext {
				return currentVal
			}

			if i == len(valueArray)-2 {
				return nextVal
			}
		}

	} else {

		//ARRAY TOO SHORT!!!!!!!
		for i := midIndex - 1; i > 1; i-- {
			lastVal := valueArray[i+1]
			currentVal := valueArray[i]
			nextVal := valueArray[i-1]

			distanceFromLast := distanceToNumber(lastVal, average)
			distanceFromCurrent := distanceToNumber(currentVal, average)
			distanceFromNext := distanceToNumber(nextVal, average)

			if distanceFromCurrent > distanceFromLast {
				return lastVal
			}

			if distanceFromCurrent < distanceFromNext {
				return currentVal
			}

			if i == len(valueArray)-2 {
				return nextVal
			}
		}

	}

	return 0.0

}

func calcKFromY(y float64, x float64) float64 {
	return x/y - 0.01*x

}
