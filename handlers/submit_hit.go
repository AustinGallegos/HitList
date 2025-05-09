package handlers

import (
	"html/template"
	"net/http"
	"spotify/handlers/utils"
)

func SubmitHit(w http.ResponseWriter, r *http.Request) {
	if err := renderSubmitHitTemplate(w); err != nil {
		utils.LogError("rendering template search.html", err)
		utils.RenderErrorPage(w, "There was an error rendering the submit hit page.", http.StatusInternalServerError)
		return
	}
}

func renderSubmitHitTemplate(w http.ResponseWriter) error {
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/submit-a-hit.html")
	if err != nil {
		return err
	}

	return tmpl.ExecuteTemplate(w, "layout.html", nil)
}
