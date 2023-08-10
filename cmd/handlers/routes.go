package handlers

import (
	"fmt"
	"net/http"
	"strconv"
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
