package model

import (
	"database/sql"
	"fmt"
	"main/view"
)

func GetAllEntityListsByUserId(db *sql.DB, user_id int, offset int, limit int) ([]view.EntityListResponse, error) {
	rows, err := db.Query("SELECT * FROM entity_lists WHERE user_id = $1 OFFSET $2 LIMIT $3", user_id, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("error querying entity_lists: %v", err)
	}
	defer rows.Close()

	results := []view.EntityListResponse{}
	for rows.Next() {
		result := view.EntityListResponse{}
		if err := rows.Scan(&result.ID, &result.EntityType, &result.UserId, &result.CreatedAt, &result.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning entity_lists: %v", err)
		}
		results = append(results, result)
	}
	return results, nil
}
