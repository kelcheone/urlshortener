package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/kelcheone/urlshortener/cmd/storage"
)

type User struct {
	ID   int
	Name string
}

var usersDB = map[int]User{
	123: {ID: 123, Name: "Alice"},
	456: {ID: 456, Name: "Bob"},
}

func GetId(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/user/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, found := usersDB[id]
	if !found {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "User ID: %d\n Name: %s\n", user.ID, user.Name)
}

func CreateUrl(w http.ResponseWriter, r *http.Request) {
	original_url := r.URL.Path[len("/create/"):]

	short_url, err := storage.CreateUrl(original_url)

	if err != nil {
		// http.Error(w, "Internal Server Error:", http.StatusInternalServerError)
		//  include the error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = godotenv.Load()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	HOST := os.Getenv("HOST")

	fmt.Fprintf(w, "Short URL: %s\n", HOST+"/"+short_url)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	to := r.URL.Path[len("/"):]

	dest, err := storage.GetOriginalLink(to)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	http.Redirect(w, r, dest, http.StatusSeeOther)
}
