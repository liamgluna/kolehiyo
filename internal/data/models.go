package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Universities UniversityModel
	Users		UserModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Universities: UniversityModel{DB: db},
		Users: UserModel{DB: db},
	}
}
