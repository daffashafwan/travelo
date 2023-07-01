package cmd

import (
	"encoding/json"
	"errors"
	"travelo/internal/graphql"
)

func (app *Application) graphqlQueryUser(query string, variables map[string]interface{}) (graphql.UserQueryResponse, error){	
	response := graphql.UserQueryResponse{}

	if query != graphql.GetAllUser && variables == nil{
		return graphql.UserQueryResponse{}, errors.New("variables is empty")
	}

	res, err := app.GraphqlClient.Query(query, variables)
	if err != nil {
		return graphql.UserQueryResponse{}, err
	}

	err = json.Unmarshal(res, &response)
	if err != nil {
		return graphql.UserQueryResponse{}, err
	}

	return response,nil
}

func (app *Application) graphqlMutationUser(query string, variables map[string]interface{}) (graphql.UserMutationResponse, error){
	response := graphql.UserMutationResponse{}

	if variables == nil{
		return graphql.UserMutationResponse{}, errors.New("variables is empty")
	}

	res, err := app.GraphqlClient.Query(query, variables)
	if err != nil {
		return graphql.UserMutationResponse{}, err
	}

	err = json.Unmarshal(res, &response)
	if err != nil {
		return graphql.UserMutationResponse{}, err
	}

	if len(response.Errors) > 0 {
		app.CustomLogger.Log(graphql.CollectErrorMessage(response.Errors))
		return graphql.UserMutationResponse{}, errors.New("Error On Graphql Occured")
	}

	
	return response, nil
}