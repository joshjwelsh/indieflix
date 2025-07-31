package model

import (
	"database/sql"
	"fmt"
	"main/view"
)

func GetAllCommentsByMovie(db *sql.DB, movie_id int) ([]view.CommentResponse, error) {
	rows, err := db.Query("SELECT * FROM comments WHERE movie_id = $1", movie_id)
	if err != nil {
		return nil, fmt.Errorf("error querying comments: %v", err)
	}
	defer rows.Close()

	results := []view.CommentResponse{}
	for rows.Next() {
		result := view.CommentResponse{}
		if err := rows.Scan(&result.ID, &result.MovieId, &result.UserText, &result.CreatedAt, &result.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning comments: %v", err)
		}
		results = append(results, result)
	}
	return results, nil
}

func GetAllCommentsByUser(db *sql.DB, user_id int) ([]view.CommentResponse, error) {
	rows, err := db.Query("SELECT * FROM comments WHERE user_id = $1", user_id)
	if err != nil {
		return nil, fmt.Errorf("error querying comments: %v", err)
	}
	defer rows.Close()

	results := []view.CommentResponse{}
	for rows.Next() {
		result := view.CommentResponse{}
		if err := rows.Scan(&result.ID, &result.MovieId, &result.UserText, &result.CreatedAt, &result.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning comments: %v", err)
		}
		results = append(results, result)
	}
	return results, nil
}
