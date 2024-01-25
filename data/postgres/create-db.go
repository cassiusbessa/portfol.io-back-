package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func CreateDb(db *sql.DB) error {
	fmt.Println("Creating database...")
	// _, err := db.Exec(`
	// CREATE DATABASE orange_portfolio;
	// `)
	// if err != nil {
	// 	return fmt.Errorf("error creating database: %v", err)
	// }
	err := createUserTable(db)
	if err != nil {
		return err
	}
	return nil
}

func createUserTable(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(255) PRIMARY KEY,
		full_name VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		image text,
		created_at timestamp NOT NULL,
		updated_at timestamp NOT NULL,
		delete_at timestamp
	);
	`)
	if err != nil {
		return fmt.Errorf("error creating user table: %v", err)
	}
	return nil
}
