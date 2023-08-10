package storage

import "github.com/kelcheone/urlshortener/cmd/storage"

func GetURL(id string) {
	db := storage.Connect()
	db
}
