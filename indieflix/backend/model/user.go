package model

import (
	"database/sql"
	"fmt"
	"main/view"
)

func AllUsers(db *sql.DB, offset, limit int) ([]view.UserPublicResponse, error) {
	rows, err := db.Query("SELECT * FROM users LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error querying users: %v", err)
	}
	defer rows.Close()

	results := []view.UserPublicResponse{}
	for rows.Next() {
		result := view.UserPrivateResponse{}
		if err := rows.Scan(&result.ID, &result.Username, &result.Email, &result.PasswordHash, &result.CreatedAt, &result.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning users: %v", err)
		}
		results = append(results, result.Prep())
	}
	return results, nil
}

func SelectUsers(db *sql.DB, id int) (view.UserPublicResponse, error) {
	result := view.UserPrivateResponse{}
	row := db.QueryRow("SELECT * FROM users WHERE id = $1", id)
	if err := row.Scan(&result.ID, &result.Username, &result.Email, &result.PasswordHash, &result.CreatedAt, &result.UpdatedAt); err != nil {
		return result.Prep(), fmt.Errorf("error scanning users: %v", err)
	}
	return result.Prep(), nil
}

func SelectUsersByUsername(db *sql.DB, username string) error {
	result := view.UserPrivateResponse{}
	row := db.QueryRow("SELECT * FROM users WHERE username = $1", username)
	if err := row.Scan(&result.ID, &result.Username, &result.Email, &result.PasswordHash, &result.CreatedAt, &result.UpdatedAt); err != nil {
		return fmt.Errorf("error scanning users: %v", err)
	}
	return nil
}

func InsertUsers(db *sql.DB, username string, password string, email string) (int, error) {
	var id int
	row := db.QueryRow("INSERT INTO users (username, passwordhash, email) VALUES ($1, $2, $3) RETURNING id", username, password, email)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("error scanning users: %v", err)
	}
	return id, nil
}

func DeleteUsers(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error deleting users: %v", err)
	}
	return nil
}

func UpdateUsers(db *sql.DB, id int, username string, password string, email string) error {
	_, err := db.Exec("UPDATE users SET username = $1, passwordhash = $2, email = $3, updated_at = NOW() WHERE id = $4", username, password, email, id)
	if err != nil {
		return fmt.Errorf("error updating users: %v", err)
	}
	return nil
}
