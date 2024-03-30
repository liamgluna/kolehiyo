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
		ImgURL   string    `json:"img_url,omitempty"`
		ImgCite  string    `json:"img_cite,omitempty"`
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
		ImgURL:   input.ImgURL,
		ImgCite:  input.ImgCite,
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
		ImgURL   *string    `json:"img_url,omitempty"`
		ImgCite  *string    `json:"img_cite,omitempty"`
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
	if input.ImgURL != nil {
		university.ImgURL = *input.ImgURL
	}
	if input.ImgCite != nil {
		university.ImgCite = *input.ImgCite
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

func (app *application) deleteUniversityHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Universities.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "university successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listUniversitiesHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.Name = app.readString(qs, "name", "")
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.SortSafelist = []string{"id", "name", "founded", "-id", "-name", "-founded"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	universities, metadata, err := app.models.Universities.GetAll(input.Name, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"universities": universities, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
