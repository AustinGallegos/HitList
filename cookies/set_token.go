package cookies

import (
	"net/http"
)

func Set_token(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    token,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   false, // change to true before deployment
		SameSite: http.SameSiteLaxMode,
	})
}
