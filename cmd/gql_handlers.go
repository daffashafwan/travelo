package cmd

import (
	"encoding/json"
	"travelo/internal/graphql"
)

func (app *Application) getAllUser() ([]graphql.User, error){

	response := make(map[string]graphql.UserResponse)

	query := `
	query MyQuery {
		user {
		  user_name
		  user_id
		  user_username
		  user_password
		}
	  }
	`

	res, err := app.GraphqlClient.Query(query)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &response)
	if err == nil {
		return response["data"].User, nil
	}

	
	return nil, nil
}