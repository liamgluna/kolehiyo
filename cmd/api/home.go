package main

import "net/http"

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	data := envelope{
		"message": "Welcome to Kolehiyo, a RESTful API for universities in the Philippines.",
		"universities": "https://kolehiyo.live/v1/universities",
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
