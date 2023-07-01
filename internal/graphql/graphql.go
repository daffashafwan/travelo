package graphql

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const (
	GetAllUser string = `
	query MyQuery {
		user {
		  user_name
		  user_id
		  user_username
		  user_password
		}
	  }
	`

	GetUserByUserName string = `
	query MyQuery($username: String!) {
		user(where: {user_username: {_eq: $username}}) {
		  user_name
		  user_username
		  user_id
		  user_password
		}
	  }
	`

	InsertOneUser string = `
	mutation MyMutation ($username: String!, $name: String!, $password: String!){
		insert_user_one(object: {user_name: $name, user_password: $password, user_username: $username}) {
		  user_id
		  user_name
		  user_password
		  user_username
		}
	  }
	`
)

type GraphqlClient struct {
	url string
}

func NewGraphqlClient(url string) *GraphqlClient {
	return &GraphqlClient{
		url: url,
	}
}

func (c *GraphqlClient) Query(query string, variables map[string]interface{}) ([]byte, error) {
	requestBody := &GraphqlRequest{
		Query: query,
		Variables: variables,
	}

	jsonPayload, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("content-type", "application/json")
	req.Header.Set("x-hasura-admin-secret", "p7q8lnZNaZjoHPtSzFgQcVwzj1mrM56GF5ysp4h0Qw7xI1rhpUQg67py9PzXTPiE")

	client := &http.Client{
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second, // Set Keep-Alive duration
			}).DialContext,
			MaxIdleConns:          100,           // Maximum number of idle connections to keep alive
			IdleConnTimeout:       90 * time.Second, // Maximum duration for an idle connection to be kept alive
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			DisableKeepAlives:     false, // Enable Keep-Alive
		},
	}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.Body != nil {
		bodyRes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return bodyRes, nil
	}

	return nil, nil
}
