package cookies

import (
	"net/http"
)

func Set_verifier(w http.ResponseWriter, verifier string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "verifier",
		Value:    verifier,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // change before prod
		SameSite: http.SameSiteLaxMode,
	})
}
