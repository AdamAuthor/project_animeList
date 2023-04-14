package content

import (
	"animeList/internal/models"
	"context"
)

type Content interface {
	Create(ctx context.Context, anime *models.Content) error
	ByID(ctx context.Context, id int) (*models.Content, error)
	Update(ctx context.Context, anime *models.Content) error
	Delete(ctx context.Context, id int) error
}
