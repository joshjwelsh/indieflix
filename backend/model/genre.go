package model

import (
	"database/sql"
	"fmt"
	"main/view"
)

func AllGenres(db *sql.DB) ([]view.GenreResponse, error) {
	rows, err := db.Query("SELECT * FROM genres")
	if err != nil {
		return nil, fmt.Errorf("error querying genres: %v", err)
	}
	defer rows.Close()

	results := []view.GenreResponse{}
	for rows.Next() {
		result := view.GenreResponse{}
		if err := rows.Scan(&result.ID, &result.Name, &result.CreatedAt, &result.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning genres: %v", err)
		}
		results = append(results, result)
	}
	return results, nil
}

func SelectGenres(db *sql.DB, id int) (view.GenreResponse, error) {
	result := view.GenreResponse{}
	row := db.QueryRow("SELECT * FROM genres WHERE id = $1")
	if err := row.Scan(&result.ID, &result.Name, &result.CreatedAt, &result.UpdatedAt); err != nil {
		return result, fmt.Errorf("error scanning genres: %v", err)
	}
	return result, nil
}

func InsertGenres(db *sql.DB, name string) (int, error) {
	var id int
	row := db.QueryRow("INSERT INTO genres (name) VALUES ($1) RETURNING id", name)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("error scanning genres: %v", err)
	}
	return id, nil
}

func DeleteGenres(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM genres WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error deleting genres: %v", err)
	}
	return nil
}

func UpdateGenres(db *sql.DB, id int, name string) error {
	_, err := db.Exec("UPDATE genres SET name = $1 WHERE id = $2", name, id)
	if err != nil {
		return fmt.Errorf("error updating genres: %v", err)
	}
	return nil
}
