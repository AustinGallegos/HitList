package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"spotify/database"
	"spotify/handlers/utils"
)

type TrackData struct {
	TrackID    string `json:"trackID"`
	TrackName  string `json:"trackName"`
	ArtistName string `json:"artist"`
	ImageLink  string `json:"image"`
}

func Success(w http.ResponseWriter, r *http.Request) {
	if !isPostRequest(r) {
		redirectToHome(w, r)
		return
	}

	trackData, err := parseTrackData(r)
	if err != nil {
		utils.LogError("decoding JSON", err)
		redirectToHome(w, r)
		return
	}

	if err := saveTrackData(trackData); err != nil {
		utils.LogError("submitting hit", err)
		utils.RenderErrorPage(w, "There was an error submitting the hit.", http.StatusInternalServerError)
		return
	}

	if err := renderSuccessTemplate(w, trackData); err != nil {
		utils.LogError("rendering template success.html", err)
		utils.RenderErrorPage(w, "There was an error rendering the success page.", http.StatusInternalServerError)
		return
	}
}

func isPostRequest(r *http.Request) bool {
	return r.Method == http.MethodPost
}

func redirectToHome(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func parseTrackData(r *http.Request) (TrackData, error) {
	var data TrackData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	return data, err
}

func saveTrackData(data TrackData) error {
	return database.Insert(
		"hit_submissions",
		data.TrackID,
		data.TrackName,
		data.ArtistName,
		data.ImageLink,
	)
}

func renderSuccessTemplate(w http.ResponseWriter, data TrackData) error {
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/success.html")
	if err != nil {
		return err
	}

	return tmpl.ExecuteTemplate(w, "layout.html", data)
}
