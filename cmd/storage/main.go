package storage

import (
	"github.com/kelcheone/urlshortener/cmd/utils"
)

func CreateUrl(original_url string) (string, error) {
	charStr := utils.RandomCharGenerator()
	db := GetDb()

	if !utils.ValidateUrl(original_url) {
		return "", &utils.InvalidUrlError{Url: original_url}
	}
	for {
		var id string
		err := db.QueryRow("SELECT short_url FROM urls WHERE short_url=$1", charStr).Scan(&id)
		if err != nil {
			break
		}
		charStr = utils.RandomCharGenerator()
	}

	_, err := db.Exec("INSERT INTO urls (short_url, original_url) VALUES ($1, $2)", charStr, original_url)

	if err != nil {
		return "", err
	}

	return charStr, nil
}

func GetOriginalLink(short_url string) (string, error) {
	db := GetDb()
	var dest string
	err := db.QueryRow("SELECT original_url FROM urls WHERE short_url=$1", short_url).Scan(&dest)
	if err != nil {
		return "", err
	}
	return dest, nil

}
