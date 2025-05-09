package handlers

import (
	"html/template"
	"net/http"
	"spotify/api_calls"
	"spotify/handlers/models"
	"spotify/handlers/utils"
)

func Search(w http.ResponseWriter, r *http.Request) {
	var data models.HandlerData

	accessToken, err := getAccessTokenFromCookie(r)
	if err != nil {
		redirectHome(w, r)
		return
	}

	if data.IsPremium, err = isUserPremium(r); err != nil {
		utils.LogError("retrieving user subscription", err)
	}

	if data.IsPremium {
		query := r.URL.Query().Get("query")
		if data.Tracks, err = getTrackResults(query, accessToken); err != nil {
			utils.LogError("retrieving track results", err)
			utils.RenderErrorPage(w, "There was an error searching your Hit.", http.StatusInternalServerError)
			return
		}
	}

	if err = renderSearchTemplate(w, data); err != nil {
		utils.LogError("rendering template search.html", err)
		utils.RenderErrorPage(w, "There was an error rendering the search page.", http.StatusInternalServerError)
		return
	}
}

func getAccessTokenFromCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func isUserPremium(r *http.Request) (bool, error) {
	cookie, err := r.Cookie("subscription")
	if err != nil {
		return false, err
	}
	return cookie.Value == "premium", nil
}

func getTrackResults(query, accessToken string) ([]models.TrackInfo, error) {
	var results []models.TrackInfo

	searchItems, err := api_calls.SearchTrack(query, accessToken)
	if err != nil {
		return nil, err
	}

	for _, item := range searchItems {
		trackData, err := api_calls.GetTrackInfo(http.DefaultClient, item.ID, accessToken)
		if err != nil {
			return nil, err
		}

		trackInfo := models.TrackInfo{
			TrackID:    item.ID,
			ArtistName: trackData["artistName"],
			TrackName:  trackData["trackName"],
			ImageLink:  trackData["imageLink"],
		}
		results = append(results, trackInfo)
	}
	return results, nil
}

func renderSearchTemplate(w http.ResponseWriter, data models.HandlerData) error {
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/search.html")
	if err != nil {
		return err
	}

	return tmpl.ExecuteTemplate(w, "layout.html", data)
}
