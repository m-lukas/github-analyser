package tests

import (
	"encoding/json"
	"testing"

	"github.com/m-lukas/github-analyser/graphql"
)

func TestCommitFrequenz(t *testing.T) {

	inputString := `{"rateLimit":{"cost":1,"remaining":4996},"repositoryOwner":{"repositories":{"edges":[{"node":{"ref":{"target":{"history":{"edges":[{"node":{"committedDate":"2019-04-05T14:08:52Z"}}]}}}}},{"node":{"ref":{"target":{"history":{"edges":[{"node":{"committedDate":"2019-03-25T18:08:49Z"}},{"node":{"committedDate":"2019-03-25T18:08:30Z"}}]}}}}}]}}}`

	var commitData graphql.CommitDataRaw

	err := json.Unmarshal([]byte(inputString), &commitData)
	if err != nil {
		t.Errorf("error while unmarshaling input string!")
	}

	frequenz := graphql.GetCommitFrequenz(&commitData)

	if frequenz != 130.00305555555556 {
		t.Errorf("false return value!")
	}

}
