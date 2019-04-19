package tests

import (
	"encoding/json"
	"testing"

	"github.com/m-lukas/github-analyser/graphql"
)

func TestCommitFrequenz(t *testing.T) {

	inputString := `{"rateLimit":{"cost":1,"remaining":4996},"repositoryOwner":{"following":{"totalCount":8},"gists":{"totalCount":0},"issues":{"totalCount":9},"organizations":{"totalCount":3},"projects":{"totalCount":0},"pullRequests":{"totalCount":28},"repositories":{"totalCount":31,"edges":[{"node":{"ref":{"target":{"history":{"edges":[{"node":{"committedDate":"2019-04-05T14:08:52Z"}}]}}}}},{"node":{"ref":{"target":{"history":{"edges":[{"node":{"committedDate":"2019-03-25T18:08:49Z"}},{"node":{"committedDate":"2019-03-25T18:08:30Z"}}]}}}}}]},"repositoriesContributedTo":{"totalCount":10},"starredRepositories":{"totalCount":10},"watching":{"totalCount":40},"commitComments":{"totalCount":1},"gistComments":{"totalCount":0},"issueComments":{"totalCount":0}}}`

	var activityData graphql.ActivityRaw

	err := json.Unmarshal([]byte(inputString), &activityData)
	if err != nil {
		t.Errorf("error while unmarshaling input string!")
	}

	frequenz := graphql.GetCommitFrequenz(&activityData)

	if frequenz != 130.00305555555556 {
		t.Errorf("false return value!")
	}

}
