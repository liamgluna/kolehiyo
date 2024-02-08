package main

import (
	"fmt"
	"net/http"

	"github.com/liamgluna/kolehiyo/internal/data"
)

func (app *application) createUniversityHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string   `json:"name"`
		Founded  int32    `json:"founded"`
		Location string   `json:"location"`
		Campuses []string `json:"campuses"`
		Website  string   `json:"website"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showUniversityHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
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
		app.serverErrorResponse(w, r, err)
	}

}
