package main

import (
	"fmt"
	"net/http"

	"github.com/kelcheone/urlshortener/cmd/handlers"
	"github.com/kelcheone/urlshortener/cmd/storage"
)

func main() {
	mux := http.NewServeMux()

	storage.Connect()

	mux.HandleFunc("/", handlers.RenderClient)

	mux.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://google.com", http.StatusSeeOther)
	})

	mux.HandleFunc("/create/", handlers.CreateUrl)

	mux.HandleFunc("/r/", handlers.Redirect)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Server listening on :8080")
	server.ListenAndServe()
}
