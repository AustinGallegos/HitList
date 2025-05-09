package handlers

import (
	"html/template"
	"net/http"
	"os"
	"spotify/api_calls"
	"spotify/cookies"
	"spotify/database"
	"spotify/handlers/models"
	"spotify/handlers/utils"
)

func Hit_Day(w http.ResponseWriter, r *http.Request) {
	if accessDenied(r) {
		redirectHome(w, r)
		return
	}

	var data models.HandlerData

	token, err := getAccessToken(r, w)
	if err != nil {
		utils.LogError("retrieving access token", err)
		redirectHome(w, r)
		return
	}

	data.IsPremium, err = getSubscription(r, w, token)
	if err != nil {
		utils.LogError("retrieving user subscription from Spotify API", err)
	}

	err = populateTrackInfo(&data)
	if err != nil {
		utils.LogError("getting Hit of the Day track info", err)
		utils.RenderErrorPage(w, "There was an error getting the Hit of the Day", http.StatusInternalServerError)
		return
	}

	data.DeviceName, err = api_calls.Get_Devices(http.DefaultClient, token, data.Tracks[0])
	if err != nil {
		utils.LogError("retrieving devices from Spotify API", err)
	}

	if err = renderHitDayTemplate(w, data); err != nil {
		utils.LogError("rendering hit-of-the-day template", err)
		utils.RenderErrorPage(w, "There was an error rendering the Hit of the Day template.", http.StatusInternalServerError)
	}
}

func accessDenied(r *http.Request) bool {
	return r.URL.Query().Get("error") == "access_denied"
}

func redirectHome(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func getAccessToken(r *http.Request, w http.ResponseWriter) (string, error) {
	verifierCookie, err := r.Cookie("verifier")
	if err != nil {
		return "", err
	}

	tokenCookie, err := r.Cookie("access_token")
	if err == nil {
		return tokenCookie.Value, nil
	}

	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	code := r.URL.Query().Get("code")

	token, err := api_calls.Get_Token(http.DefaultClient, clientID, code, verifierCookie.Value)
	if err != nil {
		return "", err
	}

	cookies.Set_token(w, token)
	return token, nil
}

func getSubscription(r *http.Request, w http.ResponseWriter, token string) (bool, error) {
	subCookie, err := r.Cookie("subscription")
	if err == nil {
		return subCookie.Value == "premium", nil
	}

	sub, err := api_calls.Get_Subscription(http.DefaultClient, token)
	if err != nil {
		return false, err
	}

	set_subscription_cookie(w, sub)
	return sub == "premium", nil
}

func populateTrackInfo(data *models.HandlerData) error {
	trackInfo, err := database.Current_hit()
	if err != nil {
		return err
	}

	data.Tracks = append(data.Tracks, trackInfo)

	return nil
}

func renderHitDayTemplate(w http.ResponseWriter, data models.HandlerData) error {
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/hit-of-the-day.html")
	if err != nil {
		return err
	}

	return tmpl.ExecuteTemplate(w, "layout.html", data)
}

func set_subscription_cookie(w http.ResponseWriter, subscription string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "subscription",
		Value:    subscription,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
}
