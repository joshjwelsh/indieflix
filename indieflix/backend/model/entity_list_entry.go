package model

import (
	"database/sql"
	"fmt"
	"main/view"
)

func GetAllEntityListEntriesByEntityListId(db *sql.DB, entity_list int) ([]view.EntityListEntryResponse, error) {
	rows, err := db.Query("SELECT * FROM entity_list_entries WHERE entity_list_id = $1", entity_list)
	if err != nil {
		return nil, fmt.Errorf("error querying entity_list_entry with entity_list(id=%v): %v", entity_list, err)
	}
	defer rows.Close()

	results := []view.EntityListEntryResponse{}
	for rows.Next() {
		result := view.EntityListEntryResponse{}
		if err := rows.Scan(&result.ID, &result.EntityListId, &result.MovieId, &result.Description, &result.CreatedAt, &result.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning entity_list_entry with entity_list(id=%v): %v", entity_list, err)
		}
		results = append(results, result)
	}
	return results, nil
}
