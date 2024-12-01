package db

import (
	"database/sql"
	"log"
)

func NewSQLiteStorage() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
