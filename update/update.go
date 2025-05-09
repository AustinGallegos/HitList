package update

import (
	"database/sql"
	"fmt"
	"log"
	"spotify/database"
	"spotify/handlers/models"
)

func Update_hit() {
	db, err := sql.Open("sqlite3", "hitlist.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	row := db.QueryRow(`
		SELECT id, track_id, track_name, artist_name, image_link
		FROM hit_submissions
		ORDER BY RANDOM() LIMIT 1
	`)

	var trackInfo models.TrackInfo
	err = row.Scan(
		&trackInfo.ID,
		&trackInfo.TrackID,
		&trackInfo.TrackName,
		&trackInfo.ArtistName,
		&trackInfo.ImageLink,
	)

	if err != nil {
		log.Fatal(err)
	}

	database.Delete()

	database.Insert(
		"current_hit",
		trackInfo.TrackID,
		trackInfo.TrackName,
		trackInfo.ArtistName,
		trackInfo.ImageLink)

	fmt.Println("Hit of the Day successfully updated.")
}
