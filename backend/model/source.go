package model

import (
	"database/sql"
	"fmt"
	"main/view"
)

func AllSources(db *sql.DB) ([]view.SourceResponse, error) {
	rows, err := db.Query("SELECT * FROM sources")
	if err != nil {
		return nil, fmt.Errorf("error querying sources: %v", err)
	}

	results := []view.SourceResponse{}
	for rows.Next() {
		result := view.SourceResponse{}
		if err := rows.Scan(&result.ID, &result.Name, &result.Website, &result.CreatedAt, &result.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning sources: %v", err)
		}
		results = append(results, result)

	}
	return results, nil
}

func SelectSources(db *sql.DB, id int) (view.SourceResponse, error) {
	result := view.SourceResponse{}
	row := db.QueryRow("SELECT * FROM sources WHERE id = $1", id)
	if err := row.Scan(&result.ID, &result.Name, &result.Website, &result.CreatedAt, &result.UpdatedAt); err != nil {
		return result, fmt.Errorf("error scanning sources: %v", err)
	}
	return result, nil
}

func InsertSources(db *sql.DB, name, website string) (int, error) {
	var id int
	row := db.QueryRow("INSERT INTO sources (name, website) VALUES ($1, $2) RETURNING id", name, website)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("error scanning sources: %v", err)
	}
	return id, nil
}

func DeleteSources(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM sources WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error deleting sources: %v", err)
	}
	return nil
}

func UpdateSources(db *sql.DB, id int, name string, website string) error {
	_, err := db.Exec("UPDATE sources SET name = $1, website = $2  WHERE id = $3", name, website, id)
	if err != nil {
		return fmt.Errorf("error updating sources: %v", err)
	}
	return nil
}
