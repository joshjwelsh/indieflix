package model

import (
	"database/sql"
	"fmt"
	"main/view"
)

func GetScoresByMovieId(db *sql.DB, movie_id int) ([]view.ScoreResponse, error) {
	rows, err := db.Query("SELECT movie_id, AVG(rating) AS rating FROM scores WHERE movie_id=$1 GROUP BY 1", movie_id)
	if err != nil {
		return nil, fmt.Errorf("error querying scores: %v", err)
	}
	defer rows.Close()

	results := []view.ScoreResponse{}
	for rows.Next() {
		result := view.ScoreResponse{}
		if err := rows.Scan(&result.MovieId, &result.Rating, &result.CreatedAt, &result.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning scores: %v", err)
		}
		results = append(results, result)
	}
	return results, nil
}
