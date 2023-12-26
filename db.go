package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func openDB() {
	var err error

	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		Name TEXT,
		Description TEXT,
		Comment Text,
		Status TEXT,
		UserID INTEGER,
		FOREIGN KEY (UserID) REFERENCES users(ID)
	)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users {
		ID INTERGER PRIMARY KEY AUTOINCREMENT,
		Email TEXT
	}`)
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
