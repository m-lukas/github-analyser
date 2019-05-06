package metrix

import (
	"math"
	"sort"

	"github.com/m-lukas/github-analyser/db"
	"github.com/m-lukas/github-analyser/util"
)

const (
	LOOP_DIR_UP   = "LOOP_DIR_UP"
	LOOP_DIR_DOWN = "LOOP_DIR_DOWN"
)

func calcK(users []*db.User, fieldType string) float64 {

	var valueArray []float64

	for _, user := range users {

		value := getField(user, fieldType)
		valueArray = append(valueArray, value)

	}

	sort.Float64s(valueArray)

	midValue := getNearestToAvgValue(valueArray)
	if midValue == 0.0 {

		biggestValue := util.BiggestValueSorted(valueArray)

		k := calcKFromY(99.9999, biggestValue)
		return k
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

	average := util.Avg(valueArray)

	midIndex := int(math.Round(float64(len(valueArray) / 2)))
	midValue := valueArray[midIndex]

	//len(0) or perfectly spreaded
	if midValue == average {
		return midValue
	}

	if len(valueArray) == 2 {
		return average
	}

	if len(valueArray) == 3 || len(valueArray) == 4 {

		var loopStart = 1
		var loopEnd = len(valueArray) - 2
		var loopDirection = LOOP_DIR_UP

		nearestValue := compareDistanceLoop(valueArray, loopStart, loopEnd, loopDirection, average)
		return nearestValue

	}

	if midValue < average {

		var loopStart = midIndex + 1
		var loopEnd = len(valueArray) - 2
		var loopDirection = LOOP_DIR_UP

		nearestValue := compareDistanceLoop(valueArray, loopStart, loopEnd, loopDirection, average)
		return nearestValue

	} else {

		var loopStart = midIndex - 1
		var loopEnd = 1
		var loopDirection = LOOP_DIR_DOWN

		nearestValue := compareDistanceLoop(valueArray, loopStart, loopEnd, loopDirection, average)
		return nearestValue

	}

	return 0.0

}

func calcKFromY(y float64, x float64) float64 {
	return x/y - 0.01*x
}

func compareDistanceLoop(valueArray []float64, startIndex int, endIndex int, direction string, average float64) float64 {

	switch direction {
	case LOOP_DIR_UP:

		for i := startIndex; i <= endIndex; i++ {

			lastVal := valueArray[i-1]
			currentVal := valueArray[i]
			nextVal := valueArray[i+1]

			nearest := nearestDistance(average, lastVal, currentVal, nextVal)
			if nearest != nextVal {
				return nearest
			}

			//last iteration
			if i == endIndex {
				return nextVal
			}
		}

	case LOOP_DIR_DOWN:

		for i := startIndex; i >= endIndex; i-- {
			lastVal := valueArray[i+1]
			currentVal := valueArray[i]
			nextVal := valueArray[i-1]

			nearest := nearestDistance(average, lastVal, currentVal, nextVal)
			if nearest != nextVal {
				return nearest
			}

			//last iteration
			if i == endIndex {
				return nextVal
			}
		}

	default:
		return 0.0
	}

	return 0.0

}

func nearestDistance(target float64, first float64, second float64, third float64) float64 {

	distanceFromFirst := util.DistanceToNumber(first, target)
	distanceFromSecond := util.DistanceToNumber(second, target)
	distanceFromThird := util.DistanceToNumber(third, target)

	if distanceFromSecond > distanceFromFirst {
		return first
	}

	if distanceFromSecond < distanceFromThird {
		return second
	}

	return third

}
