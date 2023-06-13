package view

import "database/sql"

type EntityListResponse struct {
	ID         int            `json:"id"`
	EntityType string         `json:"entity_type"`
	UserId     int            `json:"user_id"`
	Title      string         `json:"title"`
	CreatedAt  string         `json:"created_at"`
	UpdatedAt  sql.NullString `json:"updated_at"`
}

type EntityListResponseList struct {
	Page    int                  `json:"page"`
	Size    int                  `json:"size"`
	Results []EntityListResponse `json:"results"`
}
