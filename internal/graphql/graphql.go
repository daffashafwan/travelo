package graphql

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type GraphqlClient struct {
	url string
}

func NewGraphqlClient(url string) *GraphqlClient {
	return &GraphqlClient{
		url: url,
	}
}

func (c *GraphqlClient) Query(query string) ([]byte, error) {
	requestBody := &GraphqlRequest{
		Query: query,
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

	client := &http.Client{}
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
