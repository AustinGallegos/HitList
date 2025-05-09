package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Create() {
	db, err := sql.Open("sqlite3", "hitlist.db")
	if err != nil {
		log.Fatal("error opening database", err)
	}
	defer db.Close()

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS hit_submissions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		track_id TEXT NOT NULL,
		track_name TEXT NOT NULL,
		artist_name TEXT NOT NULL,
		image_link TEXT NOT NULL
	);
	`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal("error creating hit_submissions database table", err)
	}

	sqlStmt = `
	CREATE TABLE IF NOT EXISTS current_hit (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		track_id TEXT NOT NULL,
		track_name TEXT NOT NULL,
		artist_name TEXT NOT NULL,
		image_link TEXT NOT NULL
	);
	`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal("error creating current_hit database table", err)
	}
}
