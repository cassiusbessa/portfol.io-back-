package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func CreateDb(db *sql.DB) error {
	fmt.Println("Creating database...")
	if exists, err := checkDb(db); err != nil {
		return err
	} else if exists {
		return nil
	}

	_, err := db.Exec(`
		CREATE DATABASE orange_portfolio;
	`)

	if err != nil {
		return fmt.Errorf("error creating database: %v", err)
	}
	err = createTables(db)
	if err != nil {
		return err
	}

	return nil
}

func checkDb(db *sql.DB) (bool, error) {

	var exists bool
	err := db.QueryRow(`
		SELECT EXISTS (
			SELECT 1
			FROM pg_database
			WHERE datname = 'orange_portfolio'
		);
	`).Scan(&exists)

	if err != nil {
		return false, fmt.Errorf("error checking if database exists: %v", err)
	}

	return exists, nil
}

func createTables(db *sql.DB) error {
	err := createUserTable(db)
	if err != nil {
		return err
	}
	err = createProjectTable(db)
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

func createProjectTable(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS projects (
		id VARCHAR(255) PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description text,
		image text,
		user_id VARCHAR(255) NOT NULL,
		created_at timestamp NOT NULL,
		updated_at timestamp NOT NULL,
		delete_at timestamp,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
	`)
	if err != nil {
		return fmt.Errorf("error creating project table: %v", err)
	}
	return nil
}
