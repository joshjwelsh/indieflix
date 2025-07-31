package view

import "database/sql"

type EntityListEntryResponse struct {
	ID           int            `json:"id"`
	EntityListId int            `json:"entity_list_id"`
	MovieId      int            `json:"movie_id"`
	Description  string         `json:"description"`
	CreatedAt    string         `json:"created_at"`
	UpdatedAt    sql.NullString `json:"updated_at"`
}

type EntityListEntryResponseList struct {
	Page    int                       `json:"page"`
	Size    int                       `json:"size"`
	Results []EntityListEntryResponse `json:"results"`
}
