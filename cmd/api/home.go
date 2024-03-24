package main

import "net/http"

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "https://api.kolehiyo.live/v0", http.StatusMovedPermanently)
		return
	}

	data := envelope{
		"message":      "Welcome to Kolehiyo, a RESTful API for universities in the Philippines.",
		"universities": "https://api.kolehiyo.live/v0/universities",
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
