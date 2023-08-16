package handlers

import (
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
	"github.com/kelcheone/urlshortener/cmd/storage"
)

// render client

func RenderClient(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "./client/index.html")
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, nil)
}

func CreateUrl(w http.ResponseWriter, r *http.Request) {
	original_url := r.FormValue("url")

	fmt.Println(original_url)

	short_url, err := storage.CreateUrl(original_url)
	fmt.Println(short_url)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = godotenv.Load()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	HOST := os.Getenv("HOST")

	tmpl := template.Must(template.ParseFiles("template/index.html"))

	tmpl.ExecuteTemplate(w, "short-url", HOST+"/r/"+short_url)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	to := r.URL.Path[len("/r/"):]

	dest, err := storage.GetOriginalLink(to)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	http.Redirect(w, r, dest, http.StatusSeeOther)
}
