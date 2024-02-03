package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	err := app.writeJSON(w, r, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
