package graphql

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Client struct {
	Endpoint          string
	HTTPClient        *http.Client
	AutorizationToken string
	CloseRequest      bool
}

type Request struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

type Response struct {
	Data   interface{}
	Errors interface{}
}

func newClient(endpoint string, httpClient *http.Client) *Client {

	client := &Client{
		Endpoint:          endpoint,
		AutorizationToken: os.Getenv("GITHUB_TOKEN"),
		HTTPClient:        httpClient,
		CloseRequest:      true,
	}

	if httpClient == nil {
		client.HTTPClient = http.DefaultClient
	}

	return client

}

func newRequest(query string) *Request {
	request := &Request{
		Query:     query,
		Variables: make(map[string]interface{}),
	}

	return request
}

func (r *Request) variable(key string, value interface{}) {
	r.Variables[key] = value
}

func (c *Client) run(ctx context.Context, request *Request, responseData interface{}) error {

	if request == nil {
		return errors.New("Request object must not be nil!")
	}

	marshaledData, err := json.Marshal(request)
	if err != nil {
		return err
	}

	httpRequest, err := http.NewRequest(http.MethodPost, c.Endpoint, bytes.NewBuffer(marshaledData))
	if err != nil {
		return err
	}

	httpRequest.Close = c.CloseRequest

	httpRequest.Header.Set("Content-Type", "application/json; charset=utf-8")
	httpRequest.Header.Set("Accept", "application/json; charset=utf-8")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("bearer %s", c.AutorizationToken))

	if ctx == nil {
		ctx = context.Background()
	}

	httpRequest = httpRequest.WithContext(ctx)

	httpResponse, err := c.HTTPClient.Do(httpRequest)
	if err != nil {
		return err
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("Retrieved a non-ok statuscode: %d", httpResponse.StatusCode))
	}

	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return err
	}

	response := &Response{
		Data: responseData,
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	if response.Errors != nil {
		log.Println(response.Errors)
		return errors.New("Retrieved a GraphQL error!")
	}

	return nil

}

func query(userName string, query string, object interface{}) error {

	client := newClient("https://api.github.com/graphql", nil)

	query, err := readQuery(query)
	if err != nil {
		return err
	}

	request := newRequest(query)
	request.variable("name", userName)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	err = client.run(ctx, request, object)
	if err != nil {
		return err
	}

	return nil

}

func queryPop(userName string, object interface{}) error {

	client := newClient("https://api.github.com/graphql", nil)

	query := `query POPULATING($name: String!) {
		rateLimit {
		  cost
		  remaining
		}
		repositoryOwner(login: $name) {
			... on User {
			following(first: 100) {
			  edges{
				node{
				  login
				  ... on User {
					following(first: 100) {
					  edges{
						node{
						  login
						}
					  }
					}
					followers(first: 100) {
					  edges{
						node{
						  login
						}
					  }
					}
				  }
				}
			  }
			}
			followers(first: 100) {
			  edges{
				node{
				  login
				  ... on User {
					following(first: 100) {
					  edges{
						node{
						  login  
						}
					  }
					}
					followers(first: 100) {
					  edges{
						node{
						  login
						}
					  }
					}
				  }
				}
			  }
			}
		  }
		}
	  }`

	request := newRequest(query)
	request.variable("name", userName)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	err := client.run(ctx, request, object)
	if err != nil {
		return err
	}

	return nil

}
