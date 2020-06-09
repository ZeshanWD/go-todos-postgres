package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func GetConnection() *sql.DB {
	connStr := "postgres://postgres:password@localhost/todos?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}