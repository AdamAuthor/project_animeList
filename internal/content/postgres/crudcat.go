package postgres

import (
	"animeList/internal/content"
	"animeList/internal/models"
	"context"
	"fmt"
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
	_, err := D.conn.Exec("INSERT INTO content(title, author, image_url, genre, release_year) VALUES ($1, $2, $3, $4, $5)", content.Title, content.Author, content.ImageURL, content.Genre, content.ReleaseYear)

	if err != nil {
		return err
	}
	return nil
}

func (D *DB) All(ctx context.Context, filter *models.ContentFilter) ([]*models.Anime, error) {
	con := make([]*models.Anime, 0)
	basicQuery := "SELECT * FROM content"

	if filter.Query != nil {
		basicQuery = fmt.Sprintf("%s WHERE name ILIKE $1", basicQuery)

		if err := D.conn.Select(&con, basicQuery, "%"+*filter.Query+"%"); err != nil {
			return nil, err
		}

		return con, nil
	}

	if err := D.conn.Select(&con, basicQuery); err != nil {
		return nil, err
	}

	return con, nil
}

func (D *DB) ByID(ctx context.Context, id int) (*models.Anime, error) {
	con := new(models.Anime)
	if err := D.conn.Get(con, "SELECT * FROM content WHERE id=$1", id); err != nil {
		return nil, err
	}
	return con, nil
}

func (D *DB) Update(ctx context.Context, con *models.Anime) error {
	_, err := D.conn.Exec("UPDATE content SET title = $1, author = $2, image_url = $3, genre = $4, release_year = $5 WHERE id = $6", con.Title, con.Author, con.ImageURL, con.Genre, con.ReleaseYear, con.ID)

	if err != nil {
		return err
	}
	return nil
}

func (D *DB) Delete(ctx context.Context, id int) error {
	_, err := D.conn.Exec("DELETE FROM content WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
