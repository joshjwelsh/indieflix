package view

import "database/sql"

type UserRequest struct {
	Username  string         `json:"username"`
	Password  string         `json:"password"`
	Email     string         `json:"email"`
	Phone     string         `json:"phone"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at"`
}

// This includes the password hash
type UserPrivateResponse struct {
	ID           int            `json:"id"`
	Username     string         `json:"username"`
	PasswordHash string         `json:"password"`
	Email        string         `json:"email"`
	CreatedAt    string         `json:"created_at"`
	UpdatedAt    sql.NullString `json:"updated_at"`
}

// This does not include the password hash
type UserPublicResponse struct {
	ID        int            `json:"id"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at"`
}

func (u UserPrivateResponse) Prep() UserPublicResponse {
	response := UserPublicResponse{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	return response
}

type UsersResponseList struct {
	Page    int
	Size    int
	Results []UserPublicResponse `json:"results"`
}
