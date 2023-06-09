package content

import (
	"animeList/internal/models"
	"context"
)

type Database interface {
	Connect(url string) error
	Close() error
	Content() RepositoryContent
	Favorites() RepositoryFavorites
	User() UserRepo
}

type RepositoryContent interface {
	Create(ctx context.Context, anime *models.Anime) error
	All(ctx context.Context, filter *models.ContentFilter) ([]*models.Anime, error)
	ByID(ctx context.Context, id int) (*models.Anime, error)
	Update(ctx context.Context, anime *models.Anime) error
	Delete(ctx context.Context, id int) error
	FilterABC(ctx context.Context) ([]*models.Anime, error)
	FilterGenre(ctx context.Context, filter *models.ContentFilter) ([]*models.Anime, error)
	FilterAuthor(ctx context.Context, filter *models.ContentFilter) ([]*models.Anime, error)
	NewAnime(ctx context.Context) ([]*models.Anime, error)
	PopularAnime(ctx context.Context) ([]*models.Anime, error)
	UserRec(ctx context.Context, id int) ([]*models.Anime, error)
}

type RepositoryFavorites interface {
	Create(ctx context.Context, userID int, contentID int) error
	All(ctx context.Context, userID int) ([]*models.Favorite, error)
	Delete(ctx context.Context, id int) error
}

type UserRepo interface {
	Create(ctx context.Context, user *models.User) error
	Parse(ctx context.Context, email string, password string) ([]*models.User, error)
	ResetPassword(ctx context.Context, email string, password string) error
}
