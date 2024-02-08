package data

import "time"

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
