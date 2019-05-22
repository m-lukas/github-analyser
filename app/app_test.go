package app

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/m-lukas/github-analyser/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_App(t *testing.T) {

	var err error

	//init router and test server
	m := InitRouter("/api")
	testServer := httptest.NewServer(m)
	defer testServer.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//initialize mongo database and test collection
	db.TestRoot = &db.DatabaseRoot{}
	collectionName := "test_getuser"
	mongoClient, collection := setupMongoTest(t, db.TestRoot, collectionName, ctx)

	elasticClient := setupElasticTest(t, db.TestRoot, ctx)
	elasticIndex := elasticClient.Config.DefaultIndex

	//Insert test data
	err = mongoClient.Insert(&db.User{Login: "test1", Bio: "keyword"}, collectionName)
	require.Nil(t, err)
	err = mongoClient.Insert(&db.User{Login: "test2", Bio: "blabla coda", ActivityScore: 20.2, PopularityScore: 25.55}, collectionName)
	require.Nil(t, err)
	err = mongoClient.Insert(&db.User{Login: "test3", Bio: "code", ActivityScore: 15.1, PopularityScore: 5.01}, collectionName)
	require.Nil(t, err)
	err = mongoClient.Insert(&db.User{Login: "test4", Bio: "something", ActivityScore: 87.9, PopularityScore: 66.78}, collectionName)
	require.Nil(t, err)
	err = mongoClient.Insert(&db.User{Login: "test5", Bio: "It's a bio!", ActivityScore: 55.1, PopularityScore: 33.23}, collectionName)
	require.Nil(t, err)

	_, err = elasticClient.Insert(&db.ElasticUser{Login: "testE", Bio: "This is for CODE!"}, elasticIndex)
	require.Nil(t, err)

	//setTestScoreConfig(db.TestRoot)

	//BASEPATH: /api

	t.Run("/", func(t *testing.T) {
		resp, _, err := testRequest(testServer, "GET", "/api", nil)
		require.Nil(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})

	//ENDPOINT:	/user/<user>

	t.Run("/user/<user> with existing user", func(t *testing.T) {
		//cannot send request to github endpoint and has to use cache
		resp, body, err := testRequest(testServer, "GET", "/api/user/test1", nil)
		require.Nil(t, err)
		require.Equal(t, 200, resp.StatusCode)
		require.Contains(t, body, `"Login":"test1"`)
	})
	t.Run("/user/<user> with non-existing user", func(t *testing.T) {
		//cannot send request to github endpoint and has to use cache
		resp, _, err := testRequest(testServer, "GET", "/api/user/test99", nil)
		require.Nil(t, err)
		require.Equal(t, 500, resp.StatusCode)
	})

	//ENDPOINT:	/user/<user>/score

	t.Run("/user/<user>/score with existing user", func(t *testing.T) {
		//cannot send request to github endpoint and has to use cache
		resp, _, err := testRequest(testServer, "GET", "/api/user/test1/score", nil)
		require.Nil(t, err)
		require.Equal(t, 200, resp.StatusCode)
		//require.Equal(t, `{"ActivityScore":45.5,"PopularityScore":23.3}`, body) -> scores are allways recalculated
	})
	t.Run("/user/<user>/score with non-existing user", func(t *testing.T) {
		//cannot send request to github endpoint and has to use cache
		resp, _, err := testRequest(testServer, "GET", "/api/user/test99/score", nil)
		require.Nil(t, err)
		require.Equal(t, 500, resp.StatusCode)
	})

	//ENDPOINT:	/score/<score>

	t.Run("/score/<score> with score in range", func(t *testing.T) {
		//cannot send request to github endpoint and has to use cache
		resp, body, err := testRequest(testServer, "GET", "/api/score/50", nil)
		require.Nil(t, err)
		require.Equal(t, 200, resp.StatusCode)
		require.Contains(t, body, `"ActivityScore":55.1`)
	})
	t.Run("/score/<score> with score out of range", func(t *testing.T) {
		//cannot send request to github endpoint and has to use cache
		resp, _, err := testRequest(testServer, "GET", "/api/score/120", nil)
		require.Nil(t, err)
		require.Equal(t, 400, resp.StatusCode)
	})

	//ENDPOINT:	/score/<score>/next/<entries>

	t.Run("/score/<score>/next/<entries> with score in range", func(t *testing.T) {
		//cannot send request to github endpoint and has to use cache
		resp, body, err := testRequest(testServer, "GET", "/api/score/56/next/2", nil)
		require.Nil(t, err)
		require.Equal(t, 200, resp.StatusCode)
		require.Contains(t, body, `"ActivityScore":87.9`)
		require.NotContains(t, body, `"ActivityScore":55.1`)
	})
	t.Run("/score/<score>/next/<entries> with score out of range", func(t *testing.T) {
		//cannot send request to github endpoint and has to use cache
		resp, _, err := testRequest(testServer, "GET", "/api/score/120/next/100", nil)
		require.Nil(t, err)
		require.Equal(t, 400, resp.StatusCode)
	})

	//ENDPOINT:	/score/<score>/previous/<entries>

	t.Run("/score/<score>/previous/<entries> with score in range", func(t *testing.T) {
		//cannot send request to github endpoint and has to use cache
		resp, body, err := testRequest(testServer, "GET", "/api/score/55/previous/2", nil)
		require.Nil(t, err)
		require.Equal(t, 200, resp.StatusCode)
		require.Contains(t, body, `"ActivityScore":15.1`)
		require.NotContains(t, body, `"ActivityScore":55.1`)
	})
	t.Run("/score/<score>/previous/<entries> with score out of range", func(t *testing.T) {
		//cannot send request to github endpoint and has to use cache
		resp, _, err := testRequest(testServer, "GET", "/api/score/-20/previous/100", nil)
		require.Nil(t, err)
		require.Equal(t, 400, resp.StatusCode)
	})

	//ENDPOINT:	/search

	time.Sleep(1 * time.Second)
	t.Run("/search", func(t *testing.T) {
		resp, body, err := testRequest(testServer, "GET", "/api/search?query=CODE", nil)
		require.Nil(t, err)
		require.Equal(t, 200, resp.StatusCode)
		require.Contains(t, body, `CODE`)
	})

	clearMongoTestCollection(t, collection, ctx)

	_, err = elasticClient.Client.DeleteIndex(elasticClient.Config.DefaultIndex).Do(ctx)
	require.Nil(t, err, "index deletion failed")

	db.TestRoot = nil
}

//https://github.com/go-chi/chi/blob/master/mux_test.go
func testRequest(testServer *httptest.Server, method, path string, body io.Reader) (*http.Response, string, error) {

	req, err := http.NewRequest(method, testServer.URL+path, body)
	if err != nil {
		return nil, "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, "", err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	return resp, string(respBody), nil
}
