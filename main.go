package main

import (
	"fmt"
	"net/http"
	"spotify/database"
	"spotify/handlers"
	"spotify/update"
	"time"
)

func main() {
	database.Create()
	database.List("hit_submissions")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/hit-of-the-day", handlers.Hit_Day)
	http.HandleFunc("/logout", handlers.Logout)
	http.HandleFunc("/submit-a-hit", handlers.SubmitHit)
	http.HandleFunc("/search", handlers.Search)
	http.HandleFunc("/success", handlers.Success)

	update.Update_hit()

	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		for range ticker.C {
			update.Update_hit()
		}
	}()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed:", err)
		return
	}
}
