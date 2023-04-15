package models

type Anime struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	ImageURL    string `json:"image_url" db:"image_url"`
	Author      string `json:"author" db:"author"`
	Genre       string `json:"genre" db:"genre"`
	ReleaseYear int    `json:"release_year" db:"release_year"`
}

type ContentFilter struct {
	Query *string `json:"query"`
}
