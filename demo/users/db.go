package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func DB(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatalln(err)
	} else if db == nil {
		log.Fatalln("Could not connect to the database.")
	}
	return db
}
