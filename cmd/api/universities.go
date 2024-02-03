package main

import (
	"fmt"
	"net/http"
)

func (app *application) createUniversityHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new university")
}

func (app *application) showUniversityHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Show details of university %d\n", id)
}


