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

	//user
	mux.HandlerFunc("POST", "/user/login", app.login)
	mux.HandlerFunc("POST", "/user/register", app.register)

	//category
	mux.HandlerFunc("GET", "/categories", app.getCategories)
	mux.GET("/category/:id", app.getCategoryByID)
	mux.HandlerFunc("POST", "/category", app.addCategory)
	mux.PUT("/category/:id", app.editCategory)

	//review
	mux.HandlerFunc("GET", "/reviews", app.getReviews)
	mux.GET("/review/:id", app.getReviewByID)
	mux.HandlerFunc("POST", "/review", app.addReview)
	mux.PUT("/review/:id", app.editReview)


	return app.recoverPanic(mux)
}
