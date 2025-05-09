package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderErrorPage(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)

	if err := renderErrorTemplate(w, message, code); err != nil {
		LogError("rendering error template", err)
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}
}

func renderErrorTemplate(w http.ResponseWriter, message string, code int) error {
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/error.html")
	if err != nil {
		return err
	}

	return tmpl.ExecuteTemplate(w, "layout.html", map[string]interface{}{
		"Message": message,
		"Code":    code,
	})
}

func LogError(context string, err error) {
	fmt.Printf("error %s: %v\n", context, err)
}
