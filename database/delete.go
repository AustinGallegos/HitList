package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Delete() {
	db, err := sql.Open("sqlite3", "hitlist.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmts := []string{
		`DELETE FROM hit_submissions;`,
		`DELETE FROM current_hit;`,
		`DELETE FROM sqlite_sequence WHERE name='hit_submissions';`,
		`DELETE FROM sqlite_sequence WHERE name='current_hit';`,
	}

	for _, stmt := range sqlStmts {
		_, err = db.Exec(stmt)
		if err != nil {
			log.Fatal(err)
		}
	}
}
