package data

import (
	"time"

	"github.com/liamgluna/kolehiyo/internal/validator"
)

type University struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Name      string    `json:"name"`
	Founded   int32     `json:"founded"`
	Location  string    `json:"location"`
	Campuses  []string  `json:"campuses,omitempty"`
	Website   string    `json:"website"`
	Version   int32     `json:"version"`
}

func ValidateUniversity(v *validator.Validator, university *University) {
	v.Check(university.Name != "", "name", "must be provided")
	v.Check(len(university.Name) <= 150, "name", "must not be more than 150 bytes long")

	v.Check(university.Founded != 0, "founded", "must be provided")
	v.Check(university.Founded >= 1589, "founded", "must be greater than or equal to 1589")
	v.Check(university.Founded <= int32(time.Now().Year()), "founded", "must be less than or equal to the current year")

	v.Check(university.Location != "", "location", "must be provided")
	v.Check(len(university.Location) <= 500, "location", "must not be more than 500 bytes long")

	v.Check(university.Website != "", "website", "must be provided")
	v.Check(len(university.Website) <= 100, "website", "must not be more than 100 bytes long")

	v.Check(validator.Unique(university.Campuses), "campuses", "must not contain duplicate values")
}
