package main

import (
	"net/http"
	"time"

	"github.com/liamgluna/kolehiyo/internal/data"
	"github.com/liamgluna/kolehiyo/internal/validator"
)

func (app *application) createUniversityHandler(w http.ResponseWriter, r *http.Request) {
	// we decode into an input struct to prevent the client
	// from providing an id and version key in the request body
	var input struct {
		Name     string    `json:"name"`
		Founded  data.Date `json:"founded"`
		Location string    `json:"location"`
		Campuses []string  `json:"campuses"`
		Website  string    `json:"website"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	university := &data.University{
		Name:     input.Name,
		Founded:  input.Founded,
		Location: input.Location,
		Campuses: input.Campuses,
		Website:  input.Website,
	}

	v := validator.New()

	if data.ValidateUniversity(v, university); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"university": university}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showUniversityHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	founded := data.Date(time.Date(1929, time.June, 16, 0, 0, 0, 0, time.UTC))
	university := &data.University{
		ID:       id,
		Name:     "La Salle University Ozamiz",
		Founded:  founded,
		Location: "Ozamiz City, Misamis Occidental",
		Website:  "https://www.lsu.edu.ph",
		Campuses: []string{"Main Campus", "Integrated School Campus"},
		Version:  1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"university": university}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
