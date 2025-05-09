package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func Insert(tableName, trackID, trackName, artistName, imageLink string) error {
	db, err := sql.Open("sqlite3", "hitlist.db")
	if err != nil {
		return err
	}
	defer db.Close()

	sqlStmt := fmt.Sprintf("INSERT INTO %s (track_id, track_name, artist_name, image_link) VALUES(:track_id, :track_name, :artist_name, :image_link)", tableName)

	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		sql.Named("track_id", trackID),
		sql.Named("track_name", trackName),
		sql.Named("artist_name", artistName),
		sql.Named("image_link", imageLink),
	)

	if err != nil {
		return err
	}

	return nil
}
