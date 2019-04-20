package tests

import (
	"encoding/json"
	"testing"

	"github.com/m-lukas/github-analyser/graphql"
)

func TestStargazerForkSum(t *testing.T) {

	inputString := `{"repositoryOwner":{
						"followers":{"totalCount":32115},
						"repositories":{"edges":[
							{"node":{"nameWithOwner":"sindresorhus/awesome","stargazers":{"totalCount":106862},"forks":{"totalCount":13300}}},
							{"node":{"nameWithOwner":"google/material-design-lite","stargazers":{"totalCount":31169},"forks":{"totalCount":5189}}},
							{"node":{"nameWithOwner":"sindresorhus/awesome-nodejs","stargazers":{"totalCount":29644},"forks":{"totalCount":3453}}},
							{"node":{"nameWithOwner":"google/web-starter-kit","stargazers":{"totalCount":18558},"forks":{"totalCount":3020}}},
							{"node":{"nameWithOwner":"sindresorhus/awesome-electron","stargazers":{"totalCount":17183},"forks":{"totalCount":1436}}}
						]}
					}}`
	userName := "sindresorhus"

	var popularityData graphql.PopularityRaw

	err := json.Unmarshal([]byte(inputString), &popularityData)
	if err != nil {
		t.Errorf("error while unmarshaling input string!")
	}

	stargazers, forks := graphql.CalcStargazersAndForks(&popularityData, userName)

	if stargazers != 153689 {
		t.Errorf("false return value for stargazer sum!")
	}

	if forks != 18189 {
		t.Errorf("false return value for fork sum!")
	}

}
