package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	// custom error handler for 404 Not Found responses
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	// custom error handler for 405 Method Not Allowed responses
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/", app.homeHandler)
	router.HandlerFunc(http.MethodGet, "/v1", app.homeHandler)

	router.HandlerFunc(http.MethodGet, "/v1/health", app.healthHandler)

	router.HandlerFunc(http.MethodGet, "/v1/universities", app.listUniversitiesHandler)
	router.HandlerFunc(http.MethodGet, "/v1/universities/:id", app.showUniversityHandler)

	// restricted access from public
	router.HandlerFunc(http.MethodPost, "/v1/universities", app.createUniversityHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/universities/:id", app.updateUniversityHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/universities/:id", app.deleteUniversityHandler)

	return app.recoverPanic(app.enableCORS(app.rateLimit(router)))
}
