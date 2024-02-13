package data

import "database/sql"

type Models struct {
	Universities UniversityModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Universities: UniversityModel{DB: db},
	}
}
