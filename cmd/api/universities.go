package main

import (
	"fmt"
	"net/http"

	"github.com/liamgluna/kolehiyo/internal/data"
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

	university := data.University{
		ID:       id,
		Name:     "La Salle University Ozamiz",
		Founded:  1929,
		Location: "Ozamiz City, Misamis Occidental",
		Website:  "https://www.lsu.edu.ph",
		Campuses: []string{"Main Campus", "Integrated School Campus", "Heritage Campus"},
		Version:  1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"university": university}, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}

}
