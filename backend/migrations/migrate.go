package migrations

import (
	"database/sql"
	"fmt"

	"github.com/pressly/goose"
)

func RunMigrations(db *sql.DB, dir string) error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return err
	}

	// Get the current version of the database.
	current, err := goose.GetDBVersion(db)
	if err != nil {
		return err
	}

	fmt.Printf("Current version is: %d\n", current)

	// Run the migrations
	err = goose.Up(db, dir)
	if err != nil {
		return err
	}

	// Print the new version after running the migrations.
	current, err = goose.GetDBVersion(db)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully migrated to version: %d\n", current)

	return nil
}
