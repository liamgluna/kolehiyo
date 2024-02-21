package main

import "net/http"

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" || r.URL.Path == "/api"{
		http.Redirect(w, r, "https://kolehiyo.live/api/v0", http.StatusMovedPermanently)
		return
	}

	data := envelope{
		"message":      "Welcome to Kolehiyo, a RESTful API for universities in the Philippines.",
		"universities": "https://kolehiyo.live/api/v0/universities",
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
