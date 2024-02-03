package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/universities", app.createUniversityHandler)
	router.HandlerFunc(http.MethodGet, "/v1/universities/:id", app.showUniversityHandler)

	return router
}
