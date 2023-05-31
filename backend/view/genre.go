package view

import "database/sql"

type GenreRequest struct {
	Name string `json:"name"`
}

type GenreResponse struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at"`
}
