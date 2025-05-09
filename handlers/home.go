package handlers

import (
	"html/template"
	"net/http"
	"os"
	"spotify/api_calls"
	"spotify/cookies"
	"spotify/handlers/models"
	"spotify/handlers/utils"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if !isRootPath(r) {
		utils.RenderErrorPage(w, "The page you're looking for doesn't exist.", http.StatusNotFound)
		return
	}

	var data models.HandlerData

	if hasAccessToken(r) {
		data.TokenCookie = true
	} else {
		data.TokenCookie = false

		redirectURL, verifier, err := getSpotifyRedirect()
		if err != nil {
			utils.LogError("getting Spotify redirect link", err)
			utils.RenderErrorPage(w, "There was an error getting the Spotify login page.", http.StatusInternalServerError)
			return
		}

		cookies.Set_verifier(w, verifier)
		data.RedirectURL = redirectURL
	}

	if err := renderHomeTemplate(w, data); err != nil {
		utils.LogError("rendering template index.html", err)
		utils.RenderErrorPage(w, "There was an error rendering the home page.", http.StatusInternalServerError)
		return
	}
}

func isRootPath(r *http.Request) bool {
	return r.URL.Path == "/"
}

func hasAccessToken(r *http.Request) bool {
	_, err := r.Cookie("access_token")
	return err == nil
}

func getSpotifyRedirect() (string, string, error) {
	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	return api_calls.Get_Redirect(clientID)
}

func renderHomeTemplate(w http.ResponseWriter, data models.HandlerData) error {
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/index.html")
	if err != nil {
		return err
	}

	return tmpl.ExecuteTemplate(w, "layout.html", data)
}
