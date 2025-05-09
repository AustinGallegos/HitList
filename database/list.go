package database

import (
	"database/sql"
	"fmt"
	"log"
	"spotify/handlers/models"

	_ "github.com/mattn/go-sqlite3"
)

func List(tableName string) {
	db, err := sql.Open("sqlite3", "hitlist.db")
	if err != nil {
		log.Fatal("error opening database", err)
	}
	defer db.Close()

	sqlStmt := fmt.Sprintf("SELECT id, track_id, track_name, artist_name, image_link FROM %s", tableName)

	rows, err := db.Query(sqlStmt)
	if err != nil {
		log.Fatal("error querying database table to list results", err)
	}
	defer rows.Close()

	for rows.Next() {
		var trackInfo models.TrackInfo

		err := rows.Scan(
			&trackInfo.ID,
			&trackInfo.TrackID,
			&trackInfo.TrackName,
			&trackInfo.ArtistName,
			&trackInfo.ImageLink,
		)

		if err != nil {
			log.Fatal("error scanning database table rows to list results", err)
		}

		fmt.Printf(
			"ID: %d\n"+
				"Track ID: %s\n"+
				"Track Name: %s\n"+
				"Artist Name: %s\n"+
				"Image Link: %s\n\n",
			trackInfo.ID,
			trackInfo.TrackID,
			trackInfo.TrackName,
			trackInfo.ArtistName,
			trackInfo.ImageLink,
		)

	}

	err = rows.Err()
	if err != nil {
		log.Fatal("error iterating database table rows to list results", err)
	}
}
