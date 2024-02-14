package main

import (
	"errors"
	"fmt"
	"net/http"

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

	err = app.models.Universities.Insert(university)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/universities/%d", university.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"university": university}, headers)
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

	university, err := app.models.Universities.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"university": university}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) updateUniversityHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	university, err := app.models.Universities.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// to handle partial updates, we use pointers
	// to distinguish between a field that was not provided
	var input struct {
		Name     *string    `json:"name"`
		Founded  *data.Date `json:"founded"`
		Location *string    `json:"location"`
		Campuses []string   `json:"campuses"`
		Website  *string    `json:"website"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// if the input field is nil, we don't update the field
	if input.Name != nil {
		university.Name = *input.Name
	}
	if input.Founded != nil {
		university.Founded = *input.Founded
	}
	if input.Location != nil {
		university.Location = *input.Location
	}
	if input.Campuses != nil {
		university.Campuses = input.Campuses
	}
	if input.Website != nil {
		university.Website = *input.Website
	}

	v := validator.New()

	if data.ValidateUniversity(v, university); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Universities.Update(university)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"university": university}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
