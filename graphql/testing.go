package graphql

import (
	"encoding/json"
)

func generalQueryTestResult() (*GeneralDataRaw, error) {

	var rawData *GeneralDataRaw
	resBody := []byte(`{
		  "repositoryOwner": {
			"login": "m-lukas",
			"name": "Lukas Müller",
			"email": "mail@lukasmueller.de",
			"repositories": { "totalCount": 31,
			  "edges": [
				{
					"node": {
						"nameWithOwner": "m-lukas/github_analyser",
						"stargazers": { "totalCount": 50 },
						"forks": { "totalCount": 23 }
					}
				},
				{
					"node": {
						"nameWithOwner": "m-lukas/ibm_cloudfoundry",
						"stargazers": { "totalCount": 44 },
						"forks": { "totalCount": 15 }
					}
				},
				{
					"node": {
						"nameWithOwner": "m-lukas/sign-map",
						"stargazers": { "totalCount": 10 },
						"forks": { "totalCount": 3 }
					}
				},
				{
					"node": {
						"nameWithOwner": "somebody/some-project",
						"stargazers": { "totalCount": 5 },
						"forks": { "totalCount": 2 }
					}
				}
			  ]
		}}
	}`)
	err := json.Unmarshal(resBody, &rawData)
	if err != nil {
		return nil, err
	}

	return rawData, nil
}

func commitQueryTestResult() (*CommitDataRaw, error) {

	var rawData *CommitDataRaw
	resBody := []byte(`{
		  "repositoryOwner": {
			"repositories": {
			  "edges": [
				{ "node": { "ref": { "target": { "history": { "edges": 
						[
							{ "node": { "committedDate": "2019-03-25T18:08:49Z" } },
							{ "node": { "committedDate": "2019-03-25T18:08:30Z" } },
							{ "node": { "committedDate": "2019-03-25T18:08:12Z" } },
							{ "node": { "committedDate": "2019-03-25T18:04:21Z" } },
							{ "node": { "committedDate": "2019-03-25T17:56:02Z" } },
							{ "node": { "committedDate": "2019-03-25T16:27:30Z" } },
							{ "node": { "committedDate": "2019-03-25T16:27:17Z" } },
							{ "node": { "committedDate": "2019-03-25T14:37:16Z" } }
						]
				}}}}},
				{ "node": { "ref": { "target": { "history": { "edges": 
						[
							{ "node": { "committedDate": "2019-02-18T15:50:58Z" } },
							{ "node": { "committedDate": "2019-02-15T16:40:05Z" } },
							{ "node": { "committedDate": "2019-02-09T17:51:27Z" } },
							{ "node": { "committedDate": "2019-02-08T21:47:57Z" } },
							{ "node": { "committedDate": "2019-02-08T18:53:22Z" } },
							{ "node": { "committedDate": "2019-01-24T18:30:13Z" } },
							{ "node": { "committedDate": "2019-01-22T17:12:38Z" } },
							{ "node": { "committedDate": "2019-01-22T16:33:29Z" } }
						]
				}}}}},
				{ "node": { "ref": { "target": { "history": { "edges": 
						[
							{ "node": { "committedDate": "2018-12-30T22:15:17Z" } },
							{ "node": { "committedDate": "2018-12-21T22:10:16Z" } },
							{ "node": { "committedDate": "2018-11-28T09:05:44Z" } },
							{ "node": { "committedDate": "2018-11-24T17:46:23Z" } }
						]
				}}}}}
			  ]
		}}
	  }`)
	err := json.Unmarshal(resBody, &rawData)
	if err != nil {
		return nil, err
	}

	return rawData, nil
}

func populatingQueryTestResult() (*PopulatingDataRaw, error) {

	var rawData *PopulatingDataRaw
	resBody := []byte(`{
		  "repositoryOwner": {
			"following": {
			  "edges": [
				{
				  "node": {
					"login": "LBeul",
					"following": { "edges": [
						{ "node": { "login": "fluidsonic" } },
						{ "node": { "login": "m-lukas" } },
						{ "node": { "login": "florianstahr" } }
					  ]
					},
					"followers": { "edges": [
						{ "node": { "login": "m-lukas" } },
						{ "node": { "login": "florianstahr" } }
					  ]
					}
				  }
				},
				{
				  "node": {
					"login": "toorusr",
					"following": { "edges": [
						{ "node": { "login": "m-lukas" } },
						{ "node": { "login": "Nickramas" } }
					  ]
					}
				  }
				}
			  ]
		}}
	  }`)
	err := json.Unmarshal(resBody, &rawData)
	if err != nil {
		return nil, err
	}

	return rawData, nil
}
