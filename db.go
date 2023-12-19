package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func openDB() {
	var err error

	db, err := sql.Open("sqlite3", "tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}

func closeDB() {
	DB.Close()
}

func setupDB() {
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		Name TEXT,
		Description TEXT,
		Comment TEXT,
		Status TEXT
	)`)
	if err != nil {
		log.Fatal(err)
	}
}
