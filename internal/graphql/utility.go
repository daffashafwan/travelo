package graphql

import (
	"encoding/json"
	"fmt"
	"strings"
)

func CollectErrorMessage(errMsg []ErrorMessage) string{
	errMessages:= make(map[string]string)
	for _, er := range errMsg {
		errMessages[er.Message] = " [code]: " + er.Extensions.Code + " [path]: " + er.Extensions.Path
	}
	return fmt.Sprint(errMessages)
}

func BuildQueryString(query string, variables map[string]interface{}) string {
	// Create a struct to hold the query and variables
	queryData := struct {
		Query     string                 `json:"query"`
		Variables map[string]interface{} `json:"variables"`
	}{
		Query:     query,
		Variables: variables,
	}

	// Serialize the struct into JSON
	jsonData, err := json.Marshal(queryData)
	if err != nil {
		panic(err)
	}

	// Convert the JSON data to string
	queryString := string(jsonData)

	// Remove unnecessary characters added during serialization
	queryString = strings.Replace(queryString, "\\u0026", "&", -1)

	return queryString
}

func GraphqlResponseDTO(response interface{}) interface{}{
	if _, ok := response.(*UserQueryResponse); ok {
		return UserQueryResponse{}
	}

	if _, ok := response.(*UserMutationResponse); ok {
		return UserMutationResponse{}
	}

	return nil
}