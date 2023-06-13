package view

import "database/sql"

type MovieGenreResponse struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at"`
}

type MovieGenreResponseList struct {
	Page    int                  `json:"page"`
	Size    int                  `json:"size"`
	Results []MovieGenreResponse `json:"results"`
}
