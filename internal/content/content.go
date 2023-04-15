package content

import (
	"animeList/internal/models"
	"context"
)

type Content interface {
	Connect(url string) error
	Close() error
	Create(ctx context.Context, anime *models.Anime) error
	All(ctx context.Context, filter *models.ContentFilter) ([]*models.Anime, error)
	ByID(ctx context.Context, id int) (*models.Anime, error)
	Update(ctx context.Context, anime *models.Anime) error
	Delete(ctx context.Context, id int) error
}
