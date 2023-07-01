package cmd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"travelo/internal/graphql"
	"travelo/internal/response"

	"github.com/julienschmidt/httprouter"

	"travelo/internal/dto"
)

func (app *Application) status(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Status": "OK",
	}

	err := response.JSON(w, http.StatusOK, data)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *Application) login(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user, res dto.User
	var graphqlUser graphql.UserQueryResponse
	var err error
	json.Unmarshal(reqBody, &user)

	err = app.Validator.Struct(user)

	if err == nil {
		variables := map[string]interface{}{
			"username" : user.Username,
			"password" : user.Password,
		}
		graphqlUser, err = app.graphqlQueryUser(graphql.GetUserByUserNamePassword, variables)
		if err == nil {
			res = dto.User{
				ID: graphqlUser.Data.Data[0].UserID,
				Name: graphqlUser.Data.Data[0].UserName,
				Username: graphqlUser.Data.Data[0].UserUsername,
				Password: graphqlUser.Data.Data[0].UserPassword,
			}
		}
	}

	err = response.JSONCustom(w, res, err)
}

func (app *Application) getCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	var data []dto.Category

	data, err = app.GetAllCategory()

	err = response.JSONCustom(w, data, err)
}

func (app *Application) getCategoryByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	data, err := app.GetCategoryByID(ps.ByName("id"))

	err = response.JSONCustom(w, data, err)
}

func (app *Application) addCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var customer, res dto.Category
	var err error
	json.Unmarshal(reqBody, &customer)

	err = app.Validator.Struct(customer)

	if err == nil {
		res, err = app.PostCategory(customer)
	}

	err = response.JSONCustom(w, res, err)
}

func (app *Application) editCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var customer, res dto.Category
	var err error
	json.Unmarshal(reqBody, &customer)

	err = app.Validator.Struct(customer)

	if err == nil {
		res, err = app.UpdateCategory(customer, ps.ByName("id"))
	}

	err = response.JSONCustom(w, res, err)
}

func (app *Application) getReviews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	var err error
	data, err := app.GetAllReviews()
	err = response.JSONCustom(w, data, err)
}

func (app *Application) getReviewByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	data, err := app.GetReviewByID(ps.ByName("id"))

	err = response.JSONCustom(w, data, err)
}

func (app *Application) addReview(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var customer, res dto.Review
	var err error
	json.Unmarshal(reqBody, &customer)

	err = app.Validator.Struct(customer)

	if err == nil {
		res, err = app.PostReview(customer)
	}

	err = response.JSONCustom(w, res, err)
}

func (app *Application) editReview(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var customer, res dto.Review
	var err error
	json.Unmarshal(reqBody, &customer)

	err = app.Validator.Struct(customer)

	if err == nil {
		res, err = app.UpdateReview(customer, ps.ByName("id"))
	}

	err = response.JSONCustom(w, res, err)
}