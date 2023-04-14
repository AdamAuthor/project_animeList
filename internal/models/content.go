package models

type Anime struct {
	ID          int    `json:"id" postgres:"id"`
	Title       string `json:"title" postgres:"title"`
	Author      string `json:"author" postgres:"author"`
	ImageURL    string `json:"image_url" postgres:"image_url"`
	Genre       string `json:"genre" postgres:"genre"`
	ReleaseYear int    `json:"release_year" postgres:"release_year"`
}

type ContentFilter struct {
	Query *string `json:"query"`
}
