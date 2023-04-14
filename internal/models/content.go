package models

type Content struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	ImageURL    string `json:"image_url"`
	Genre       string `json:"genre"`
	ReleaseYear int    `json:"release_year"`
}
