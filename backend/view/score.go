package view

import "database/sql"

type ScoreResponse struct {
	MovieId   int            `json:"movie_id"`
	Rating    int            `json:"rating"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at"`
}

type ScoreResponseList struct {
	Page    int             `json:"page"`
	Size    int             `json:"size"`
	Results []ScoreResponse `json:"results"`
}
