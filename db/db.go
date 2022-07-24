package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConectDB() *sql.DB {
	connStr := "postgres://postgres:postgrespw@localhost:55000/store?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
