package postgres

import (
	"animeList/internal/content"
	"animeList/internal/models"
	"context"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/jackc/pgx/v5"
	"github.com/jmoiron/sqlx"
	"log"
)

type DB struct {
	conn    *sqlx.DB
	content content.Content
}

func NewDB() content.Content {
	return &DB{}
}

func (D *DB) Connect(url string) error {
	conn, err := sqlx.Connect("pgx", url)
	if err != nil {
		return err
	}

	if err := conn.Ping(); err != nil {
		return err
	}

	D.conn = conn
	log.Println("DB has successful pinging")
	return nil
}

func (D *DB) Close() error {
	return D.conn.Close()
}

func (D *DB) Create(ctx context.Context, content *models.Anime) error {
	_, err := D.conn.Exec("INSERT INTO anime(title, author, genre, year, image) VALUES ($1, $2, $3, $4, $5)", content.Title, content.Author, content.Genre, content.ReleaseYear, content.ImageURL)

	if err != nil {
		return err
	}
	return nil
}

func (D *DB) All(ctx context.Context, filter *models.ContentFilter) ([]*models.Anime, error) {
	var anime []*models.Anime
	basicQuery := "SELECT * FROM anime"

	if filter.Query != nil {
		basicQuery = fmt.Sprintf("%s WHERE title ILIKE $1", basicQuery)
		queryArg := "%" + *filter.Query + "%"
		if err := D.conn.Select(&anime, basicQuery, queryArg); err != nil {
			return nil, err
		}

		return anime, nil
	}

	if err := D.conn.Select(&anime, basicQuery); err != nil {
		return nil, err
	}

	return anime, nil
}

func (D *DB) ByID(ctx context.Context, id int) (*models.Anime, error) {
	con := new(models.Anime)
	if err := D.conn.Get(con, "SELECT * FROM anime WHERE id=$1", id); err != nil {
		return nil, err
	}
	return con, nil
}

func (D *DB) Update(ctx context.Context, con *models.Anime) error {
	_, err := D.conn.Exec("UPDATE anime SET title = $1, author = $2, image_url = $3, genre = $4, release_year = $5 WHERE id = $6", con.Title, con.Author, con.ImageURL, con.Genre, con.ReleaseYear, con.ID)

	if err != nil {
		return err
	}
	return nil
}

func (D *DB) Delete(ctx context.Context, id int) error {
	_, err := D.conn.Exec("DELETE FROM anime WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
