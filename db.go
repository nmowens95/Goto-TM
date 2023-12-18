package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func openDB() error {
	db, err := sql.Open("sqlite3", "./task.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}

func closeDB() error {
	return DB.Close()
}

func setupDB() error {
	_, err := DB.Exec()
	if err != nil {
		return err
	}
	return nil
}
