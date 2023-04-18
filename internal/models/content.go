package models

type Anime struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Author      string `json:"author" db:"author"`
	Genre       string `json:"genre" db:"genre"`
	ReleaseYear int    `json:"year" db:"year"`
	ImageURL    string `json:"image" db:"image"`
	Views       int    `json:"views" db:"views"`
}

type Genre struct {
	Name  string `db:"genre"`
	Count int    `db:"count"`
}
