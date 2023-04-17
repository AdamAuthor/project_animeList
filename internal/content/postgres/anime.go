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

func (D *RepositoryContent) UserRec(ctx context.Context, id int) ([]*models.Anime, error) {
	var genres []*models.Genre
	var recAnime []*models.Anime
	err := D.conn.Select(&genres, `SELECT unnest(string_to_array(genre, ',')) as genre, COUNT(*) as count
FROM anime
GROUP BY unnest(string_to_array(genre, ','))
ORDER BY count DESC
LIMIT 3;
`)

	err = D.conn.Select(&recAnime, `SELECT a.id, a.title, a.views, a.author, a.year, a.image, a.genre
FROM anime a
LEFT JOIN favorites f ON a.id = f.animeid
WHERE f.animeid IS NULL
AND a.genre IN (
    SELECT genre
    FROM anime
    JOIN favorites ON anime.id = favorites.animeid
    WHERE favorites.userid = $1
    GROUP BY genre
    ORDER BY COUNT(*) DESC
    LIMIT 3
)
ORDER BY 
  CASE WHEN a.genre LIKE '%' || $2 || '%' AND a.genre LIKE '%' || $3 || '%' AND a.genre LIKE '%' || $4 || '%' THEN 1
       WHEN (a.genre LIKE '%' || $2 || '%' AND a.genre LIKE '%' || $3 || '%') OR
            (a.genre LIKE '%' || $2 || '%' AND a.genre LIKE '%' || $4 || '%') OR
            (a.genre LIKE '%' || $3 || '%' AND a.genre LIKE '%' || $4 || '%') THEN 2
       ELSE 3
  END

LIMIT 4;`, id, genres[0].Name, genres[1].Name, genres[2].Name)

	if err != nil {
		return nil, err
	}
	return recAnime, nil
}

func (D *RepositoryContent) PopularAnime(ctx context.Context) ([]*models.Anime, error) {
	var popularAnime []*models.Anime
	err := D.conn.Select(&popularAnime, `SELECT * FROM anime ORDER BY views DESC LIMIT 4;`)
	if err != nil {
		return nil, err
	}
	return popularAnime, nil
}

func (D *RepositoryContent) NewAnime(ctx context.Context) ([]*models.Anime, error) {
	var newAnime []*models.Anime
	err := D.conn.Select(&newAnime, `SELECT * FROM anime ORDER BY id DESC LIMIT 4;
`)
	if err != nil {
		return nil, err
	}
	return newAnime, nil
}

func (D *RepositoryContent) FilterAuthor(ctx context.Context, filter *models.ContentFilter) ([]*models.Anime, error) {
	var anime []*models.Anime
	basicQuery := "SELECT * FROM anime"

	if filter.Query != nil {
		basicQuery = fmt.Sprintf("%s WHERE author ILIKE $1", basicQuery)
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

func (D *RepositoryContent) FilterABC(ctx context.Context) ([]*models.Anime, error) {
	var anime []*models.Anime
	err := D.conn.Select(&anime, `SELECT * FROM anime ORDER BY title ASC`)
	if err != nil {
		return nil, err
	}
	return anime, nil
}

func (D *RepositoryContent) FilterGenre(ctx context.Context, filter *models.ContentFilter) ([]*models.Anime, error) {
	var anime []*models.Anime
	basicQuery := "SELECT * FROM anime"

	if filter.Query != nil {
		basicQuery = fmt.Sprintf("%s WHERE genre ILIKE $1", basicQuery)
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

func NewContentRepository(conn *sqlx.DB) content.RepositoryContent {
	return &RepositoryContent{conn: conn}
}

func (D *RepositoryContent) Create(ctx context.Context, content *models.Anime) error {
	_, err := D.conn.Exec("INSERT INTO anime(title, author, genre, year, image, views) VALUES ($1, $2, $3, $4, $5, $6)", content.Title, content.Author, content.Genre, content.ReleaseYear, content.ImageURL, content.Views)

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
	_, err := D.conn.Exec("UPDATE anime SET title = $1, author = $2, image = $3, genre = $4, year = $5, views = $6 WHERE id = $7", con.Title, con.Author, con.ImageURL, con.Genre, con.ReleaseYear, con.Views, con.ID)

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
