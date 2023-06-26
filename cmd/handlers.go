package cmd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

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

func (app *Application) getCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	var err error
	data, err := app.GetAllCategory()
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
