package main

import (
	"database/sql"
	"log"

	"github.com/rss/internal/database"
)

func establishDatabaseConnection(DB_URL string) *database.Queries {

	connection, err := sql.Open("postgres", DB_URL)

	if err != nil {
		log.Fatal("Sorry an error occur connecting to database", err)
		return nil
	}

	return database.New(connection)
}
