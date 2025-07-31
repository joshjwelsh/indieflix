package view

import (
	"database/sql"
	"encoding/json"
)

type MovieResponse struct {
	ID        int             `json:"id"`
	Name      string          `json:"name"`
	SourceId  int             `json:"source_id"`
	Metadata  json.RawMessage `json:"metadata"`
	Showtimes json.RawMessage `json:"showtimes"`
	Available bool            `json:"available"`
	CreatedAt string          `json:"created_at"`
	UpdatedAt sql.NullString  `json:"updated_at"`
}

type MoviesResponseList struct {
	Page    int             `json:"page"`
	Size    int             `json:"size"`
	Results []MovieResponse `json:"results"`
}
