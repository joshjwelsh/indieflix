package model

import (
	"database/sql"
	"fmt"
	"main/view"
)

func GetAllMovieGenresByMovieId(db *sql.DB, movie_id int) ([]view.MovieGenreResponse, error) {
	rows, err := db.Query("SELECT * FROM movie_genres WHERE movie_id = $1", movie_id)
	if err != nil {
		return nil, fmt.Errorf("error querying movie_genres with movie(id=%v): %v", movie_id, err)
	}
	defer rows.Close()

	results := []view.MovieGenreResponse{}
	for rows.Next() {
		result := view.MovieGenreResponse{}
		if err := rows.Scan(&result.ID, &result.Name, &result.CreatedAt, &result.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning movie_genres with movie(id=%v): %v", movie_id, err)
		}
		results = append(results, result)
	}
	return results, nil
}
