package view

import (
	"database/sql"
)

type SourceResponse struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	Website   string         `json:"website"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at"`
}

type SourceRequest struct {
	Name    string `json:"name"`
	Website string `json:"website"`
}

type SourcesResponseList struct {
	Page    int
	Size    int
	Results []SourceResponse `json:"results"`
}
