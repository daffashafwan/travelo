package cmd

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) routes() http.Handler {
	mux := httprouter.New()

	mux.NotFound = http.HandlerFunc(app.notFound)
	mux.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowed)

	mux.HandlerFunc("GET", "/status", app.status)

	//category
	mux.HandlerFunc("GET", "/categories", app.getCategories)
	mux.GET("/category/:id", app.getCategoryByID)
	mux.HandlerFunc("POST", "/category", app.addCategory)
	mux.PUT("/category/:id", app.editCategory)


	return app.recoverPanic(mux)
}
