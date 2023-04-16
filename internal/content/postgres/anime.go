package postgres

import (
	"animeList/internal/content"
	"animeList/internal/models"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func (D *Database) Content() content.RepositoryContent {
	if D.content == nil {
		D.content = NewContentRepository(D.conn)
	}

	return D.content
}

type RepositoryContent struct {
	conn *sqlx.DB
}

func NewContentRepository(conn *sqlx.DB) content.RepositoryContent {
	return &RepositoryContent{conn: conn}
}

func (D *RepositoryContent) Create(ctx context.Context, content *models.Anime) error {
	_, err := D.conn.Exec("INSERT INTO anime(title, author, genre, year, image) VALUES ($1, $2, $3, $4, $5)", content.Title, content.Author, content.Genre, content.ReleaseYear, content.ImageURL)

	if err != nil {
		return err
	}
	return nil
}

func (D *RepositoryContent) All(ctx context.Context, filter *models.ContentFilter) ([]*models.Anime, error) {
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

func (D *RepositoryContent) ByID(ctx context.Context, id int) (*models.Anime, error) {
	con := new(models.Anime)
	if err := D.conn.Get(con, "SELECT * FROM anime WHERE id=$1", id); err != nil {
		return nil, err
	}
	return con, nil
}

func (D *RepositoryContent) Update(ctx context.Context, con *models.Anime) error {
	_, err := D.conn.Exec("UPDATE anime SET title = $1, author = $2, image = $3, genre = $4, year = $5 WHERE id = $6", con.Title, con.Author, con.ImageURL, con.Genre, con.ReleaseYear, con.ID)

	if err != nil {
		return err
	}
	return nil
}

func (D *RepositoryContent) Delete(ctx context.Context, id int) error {
	_, err := D.conn.Exec("DELETE FROM anime WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
