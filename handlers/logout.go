package handlers

import (
	"net/http"
	"spotify/cookies"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	cookies.Logout(w)
}
