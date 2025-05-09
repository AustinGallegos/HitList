package database

import (
	"database/sql"
	"spotify/handlers/models"
)

func Current_hit() (models.TrackInfo, error) {
	db, err := sql.Open("sqlite3", "hitlist.db")
	if err != nil {
		return models.TrackInfo{}, err
	}
	defer db.Close()

	row := db.QueryRow(`SELECT id, track_id, track_name, artist_name, image_link FROM current_hit`)
	var trackInfo models.TrackInfo
	err = row.Scan(
		&trackInfo.ID,
		&trackInfo.TrackID,
		&trackInfo.TrackName,
		&trackInfo.ArtistName,
		&trackInfo.ImageLink,
	)
	if err != nil {
		return models.TrackInfo{}, err
	}

	return trackInfo, nil
}
