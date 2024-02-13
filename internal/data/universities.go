package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/liamgluna/kolehiyo/internal/validator"
	"github.com/lib/pq"
)

type UniversityModel struct {
	DB *sql.DB
}

type University struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Name      string    `json:"name"`
	Founded   Date      `json:"founded"`
	Location  string    `json:"location"`
	Campuses  []string  `json:"campuses,omitempty"`
	Website   string    `json:"website"`
	Version   int32     `json:"version"`
}

func ValidateUniversity(v *validator.Validator, university *University) {
	v.Check(university.Name != "", "name", "must be provided")
	v.Check(len(university.Name) <= 150, "name", "must not be more than 150 bytes long")

	founded := time.Time(university.Founded)
	v.Check(!founded.IsZero(), "founded", "must be provided")
	v.Check(founded.Year() >= 1589, "founded", "must be greater than or equal to 1589")
	v.Check(founded.Year() <= time.Now().Year(), "founded", "must be less than or equal to the current year")

	v.Check(university.Location != "", "location", "must be provided")
	v.Check(len(university.Location) <= 500, "location", "must not be more than 500 bytes long")

	v.Check(university.Website != "", "website", "must be provided")
	v.Check(len(university.Website) <= 100, "website", "must not be more than 100 bytes long")

	v.Check(validator.Unique(university.Campuses), "campuses", "must not contain duplicate values")
}

func (m UniversityModel) Insert(university *University) error {
	query := `
		INSERT INTO universities (name, founded, location, campuses, website)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, version`

	args := []any{university.Name, time.Time(university.Founded), university.Location, pq.Array(university.Campuses), university.Website}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&university.ID, &university.CreatedAt, &university.Version)
}

func (m UniversityModel) Get(id int64) (*University, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, created_at, name, founded, location, campuses, website, version
		FROM universities
		WHERE id = $1`

	var university University

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&university.ID,
		&university.CreatedAt,
		&university.Name,
		&university.Founded,
		&university.Location,
		pq.Array(&university.Campuses),
		&university.Website,
		&university.Version)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &university, nil
}
