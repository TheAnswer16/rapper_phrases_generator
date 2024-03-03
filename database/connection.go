package database

import (
	"database/sql"
	"os"
	"fmt"
	_ "github.com/lib/pq"
)

	
	func Connect() (*sql.DB, error) {

		// Get the database connection details from environment variables
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")

		// Create the connection string
		connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", dbHost, dbPort, dbName, dbUser, dbPassword)
		fmt.Println(connStr)
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			return nil, err
		}

		// Test the database connection
		err = db.Ping()
		if err != nil {
			return nil, err
		}

		return db, nil
	}
