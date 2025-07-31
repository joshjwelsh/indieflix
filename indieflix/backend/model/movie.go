package model

import (
	"database/sql"
	"fmt"
	"main/view"
)

func AllMovies(db *sql.DB, offset, limit int) ([]view.MovieResponse, error) {
	rows, err := db.Query("SELECT * FROM movies LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error querying movies: %v", err)
	}
	defer rows.Close()

	results := []view.MovieResponse{}
	for rows.Next() {
		result := view.MovieResponse{}
		if err := rows.Scan(&result.ID, &result.Name, &result.SourceId, &result.Metadata, &result.Showtimes, &result.Available); err != nil {
			return nil, fmt.Errorf("error scanning movies: %v", err)
		}
		results = append(results, result)
	}
	return results, nil
}

func SelectMovies(db *sql.DB, id int) (view.MovieResponse, error) {
	result := view.MovieResponse{}
	row := db.QueryRow("SELECT * FROM movies WHERE id = $1", id)
	if err := row.Scan(&result.ID, &result.Name, &result.SourceId, &result.Metadata, &result.Showtimes, &result.Available); err != nil {
		return result, fmt.Errorf("error scanning movies: %v", err)
	}
	return result, nil
}
