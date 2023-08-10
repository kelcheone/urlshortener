package models

type URL struct {
	ID           int    `json:"id"`
	Original_url string `json:"original_url"`
	Short_url    string `json:"short_url"`
	Clicks       int    `json:"clicks"`
}
