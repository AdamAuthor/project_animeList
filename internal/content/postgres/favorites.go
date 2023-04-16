package postgres

import (
	"animeList/internal/content"
	"animeList/internal/models"
	"context"
	"github.com/jmoiron/sqlx"
)

func (D *Database) Favorites() content.RepositoryFavorites {
	if D.favorite == nil {
		D.favorite = NewFavoritesRepository(D.conn)
	}

	return D.favorite
}

type RepositoryFavorites struct {
	conn *sqlx.DB
}

func (r RepositoryFavorites) Create(ctx context.Context, userID int, contentID int) error {
	_, err := r.conn.Exec(`INSERT INTO favorites ("userID", "contentID") VALUES ($1, $2)`, userID, contentID)
	if err != nil {
		return err
	}
	return nil
}

func (r RepositoryFavorites) All(ctx context.Context, userID int) ([]*models.Favorite, error) {
	var favorites []*models.Favorite
	err := r.conn.Select(&favorites, `SELECT favorites.id, favorites."userID", to_json(anime.*) as anime FROM favorites JOIN anime ON favorites."contentID"=anime.id WHERE favorites."userID"=$1`, userID)
	if err != nil {
		return nil, err
	}
	return favorites, nil
}

func (r RepositoryFavorites) Delete(ctx context.Context, id int) error {
	_, err := r.conn.Exec("DELETE FROM favorites WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func NewFavoritesRepository(conn *sqlx.DB) content.RepositoryFavorites {
	return &RepositoryFavorites{conn: conn}
}
