package cookies

import (
	"net/http"
)

func Logout(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:   "access_token",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
}
