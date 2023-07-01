package cmd

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"travelo/internal/graphql"
	"travelo/internal/response"

	"github.com/go-redis/redis"
	"github.com/julienschmidt/httprouter"

	"travelo/internal/dto"

	"travelo/internal/utilities"

	"golang.org/x/crypto/bcrypt"
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

func (app *Application) login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user dto.User
	var res dto.User
	var graphqlUser graphql.UserQueryResponse
	var err error

	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		err = response.JSONCustom(w, res, err, http.StatusBadRequest)
		return
	}

	err = app.Validator.Struct(user)
	if err != nil {
		err = response.JSONCustom(w, res, err, http.StatusBadRequest)
		return
	}

	variables := map[string]interface{}{
		"username": user.Username,
	}

	graphqlUser, err = app.graphqlQueryUser(graphql.GetUserByUserName, variables)
	if err != nil {
		err = response.JSONCustom(w, res, err, http.StatusInternalServerError)
		return
	}

	res = dto.User{
		ID:       graphqlUser.Data.Data[0].UserID,
		Name:     graphqlUser.Data.Data[0].UserName,
		Username: graphqlUser.Data.Data[0].UserUsername,
		Password: graphqlUser.Data.Data[0].UserPassword,
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(user.Password))
	if err != nil {
		err = response.JSONCustom(w, res, errors.New("Invalid username or password"), http.StatusUnauthorized)
		return
	}

	jwt := utilities.GenerateJWT(res.Username)

	err = app.Redis.Set(jwt, "valid", time.Hour).Err()
	if err != nil {
		err = response.JSONCustom(w, res, errors.New("Failed set redis jwt"), http.StatusInternalServerError)
		return
	}

	res.JWTToken = jwt

	err = response.JSONCustom(w, res, err, http.StatusOK)
}

func (app *Application) register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user dto.User
	var res dto.User
	var graphqlUser graphql.UserMutationResponse
	var err error

	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		err = response.JSONCustom(w, res, err, http.StatusBadRequest)
		return
	}

	err = app.Validator.Struct(user)
	if err != nil {
		err = response.JSONCustom(w, res, err, http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		err = response.JSONCustom(w, res, err, http.StatusInternalServerError)
		return
	}

	variables := map[string]interface{}{
		"username": user.Username,
		"password": string(hashedPassword),
		"name":     user.Name,
	}

	graphqlUser, err = app.graphqlMutationUser(graphql.InsertOneUser, variables)
	if err != nil {
		err = response.JSONCustom(w, res, err, http.StatusInternalServerError)
		return
	}

	res = dto.User{
		ID:       graphqlUser.Data.Data.UserID,
		Name:     graphqlUser.Data.Data.UserName,
		Username: graphqlUser.Data.Data.UserUsername,
		Password: graphqlUser.Data.Data.UserPassword,
	}

	err = response.JSONCustom(w, res, err, http.StatusOK)
}

func (app *Application) getCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	var data []dto.Category

	if err := app.checkAuth(r.Header.Get("Authorization")); err != nil{
		err = response.JSONCustom(w, nil, err, http.StatusUnauthorized)
		return
	}

	data, err = app.GetAllCategory()

	err = response.JSONCustom(w, data, err, http.StatusOK)
}

func (app *Application) getCategoryByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	var data dto.Category

	if err := app.checkAuth(r.Header.Get("Authorization")); err != nil{
		err = response.JSONCustom(w, data, err, http.StatusUnauthorized)
		return
	}
	
	data, err = app.GetCategoryByID(ps.ByName("id"))

	err = response.JSONCustom(w, data, err, http.StatusOK)
}

func (app *Application) addCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Read request body
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = response.JSONCustom(w, nil, err, http.StatusBadRequest)
		return
	}

	// Unmarshal request body into a Category struct
	var category dto.Category
	err = json.Unmarshal(reqBody, &category)
	if err != nil {
		err = response.JSONCustom(w, nil, err, http.StatusBadRequest)
		return
	}

	// Validate the category struct
	err = app.Validator.Struct(category)
	if err != nil {
		err = response.JSONCustom(w, nil, err, http.StatusBadRequest)
		return
	}

	// Check authorization
	if err := app.checkAuth(r.Header.Get("Authorization")); err != nil {
		err = response.JSONCustom(w, nil, err, http.StatusUnauthorized)
		return
	}

	// Add the category
	res, err := app.PostCategory(category)
	if err != nil {
		err = response.JSONCustom(w, nil, err, http.StatusInternalServerError)
		return
	}

	// Return success response
	err = response.JSONCustom(w, res, nil, http.StatusOK)
}

func (app *Application) editCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Read request body
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = response.JSONCustom(w, nil, err, http.StatusBadRequest)
		return
	}

	// Unmarshal request body into a Category struct
	var category dto.Category
	err = json.Unmarshal(reqBody, &category)
	if err != nil {
		err = response.JSONCustom(w, nil, err, http.StatusBadRequest)
		return
	}

	// Validate the category struct
	err = app.Validator.Struct(category)
	if err != nil {
		err = response.JSONCustom(w, nil, err, http.StatusBadRequest)
		return
	}

	// Check authorization
	if err := app.checkAuth(r.Header.Get("Authorization")); err != nil {
		err = response.JSONCustom(w, nil, err, http.StatusUnauthorized)
		return
	}

	// Update the category
	res, err := app.UpdateCategory(category, ps.ByName("id"))
	if err != nil {
		err = response.JSONCustom(w, nil, err, http.StatusInternalServerError)
		return
	}

	// Return success response
	err = response.JSONCustom(w, res, nil, http.StatusOK)
}


func (app *Application) getReviews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Check authorization
	if err := app.checkAuth(r.Header.Get("Authorization")); err != nil {
		err = response.JSONCustom(w, nil, err, http.StatusUnauthorized)
		return
	}

	// Get all reviews
	data, err := app.GetAllReviews()
	if err != nil {
		err = response.JSONCustom(w, nil, err, http.StatusInternalServerError)
		return
	}

	// Return success response
	err = response.JSONCustom(w, data, nil, http.StatusOK)
}


func (app *Application) getReviewByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Check authorization
	if err := app.checkAuth(r.Header.Get("Authorization")); err != nil {
		err = response.JSONCustom(w, nil, err, http.StatusUnauthorized)
		return
	}

	// Get review by ID
	data, err := app.GetReviewByID(ps.ByName("id"))
	if err != nil {
		err = response.JSONCustom(w, nil, err, http.StatusInternalServerError)
		return
	}

	// Return success response
	err = response.JSONCustom(w, data, nil, http.StatusOK)
}


func (app *Application) addReview(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Read request body
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = response.JSONCustom(w, nil, err, http.StatusInternalServerError)
		return
	}

	var customer, res dto.Review
	// Unmarshal request body into customer
	err = json.Unmarshal(reqBody, &customer)
	if err != nil {
		err = response.JSONCustom(w, nil, err, http.StatusBadRequest)
		return
	}

	// Validate customer data
	err = app.Validator.Struct(customer)
	if err != nil {
		err = response.JSONCustom(w, nil, err, http.StatusBadRequest)
		return
	}

	// Check authorization
	if err := app.checkAuth(r.Header.Get("Authorization")); err != nil {
		err = response.JSONCustom(w, nil, err, http.StatusUnauthorized)
		return
	}

	// Post review
	res, err = app.PostReview(customer)
	if err != nil {
		err = response.JSONCustom(w, nil, err, http.StatusInternalServerError)
		return
	}

	// Return success response
	err = response.JSONCustom(w, res, nil, http.StatusOK)
}


func (app *Application) editReview(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Read request body
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = response.JSONCustom(w, nil, err, http.StatusInternalServerError)
		return
	}

	var customer, res dto.Review
	// Unmarshal request body into customer
	err = json.Unmarshal(reqBody, &customer)
	if err != nil {
		err = response.JSONCustom(w, nil, err, http.StatusBadRequest)
		return
	}

	// Validate customer data
	err = app.Validator.Struct(customer)
	if err != nil {
		err = response.JSONCustom(w, nil, err, http.StatusBadRequest)
		return
	}

	// Check authorization
	if err := app.checkAuth(r.Header.Get("Authorization")); err != nil {
		err = response.JSONCustom(w, nil, err, http.StatusUnauthorized)
		return
	}

	// Update review
	res, err = app.UpdateReview(customer, ps.ByName("id"))
	if err != nil {
		err = response.JSONCustom(w, nil, err, http.StatusInternalServerError)
		return
	}

	// Return success response
	err = response.JSONCustom(w, res, nil, http.StatusOK)
}


func (app *Application) checkAuth(authHeader string) error {

	// Check if the Authorization header is present
	if authHeader == "" {
		return errors.New("Auth Header not present")
	}

	jwtToken := strings.TrimPrefix(authHeader, "Bearer ")

	err := app.Redis.Get(jwtToken).Err()
	if err == redis.Nil {
		return errors.New("JWT not found")
	} else if err != nil {
		return err
	}

	return nil
}
