package view

import (
	"database/sql"
	"fmt"
)

type ProfileResponse struct {
	ID          int      `json:"id"`
	Username    string   `json:"username"`
	Perferences []string `json:"perferences"`
}

type UserRequest struct {
	Username    string   `json:"username"`
	Password    string   `json:"password"`
	Email       string   `json:"email"`
	Phone       string   `json:"phone"`
	Birthdate   string   `json:"birthdate"`
	Perferences []string `json:"perferences"`
}

type UserUpdateRequest struct {
	Username   string   `json:"username"`
	Password   string   `json:"password"`
	Prefrences []string `json:"prefrences"`
	ID         int      `json:"id"`
}

type SourceResponse struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	Website   string         `json:"website"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at"`
}

func (s SourceResponse) String() string {
	return fmt.Sprintf("id: %v, name: %v, website: %v, created_at: %v, updated_at: %v", s.ID, s.Name, s.Website, s.CreatedAt, s.UpdatedAt)
}
