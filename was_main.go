package main

import (
	"fmt"
	"net/http"

	"github.com/kelcheone/urlshortener/cmd/handlers"
)

func Idk() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the homepage!")
	})

	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is the about page.")
	})

	mux.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://google.com", http.StatusSeeOther)
	})

	mux.HandleFunc("/user/", handlers.GetId)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Server listening on :8080")
	server.ListenAndServe()
}
