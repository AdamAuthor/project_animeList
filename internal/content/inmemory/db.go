package inmemory

import (
	"animeList/internal/content"
	"animeList/internal/models"
	"context"
	"fmt"
	"sync"
)

// DB saves information about laptops
type DB struct {
	data map[int]*models.Content
	mu   *sync.RWMutex
}

// NewDB is the function for creating basic Database
func NewDB() content.Content {
	return &DB{
		data: make(map[int]*models.Content),
		mu:   new(sync.RWMutex),
	}
}

// Create for creating new element in DB
func (db *DB) Create(ctx context.Context, manga *models.Content) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.data[manga.ID] = manga

	return nil
}

// ByID is used for reading elements by id in DB
func (db *DB) ByID(ctx context.Context, id int) (*models.Content, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	manga, ok := db.data[id]
	if !ok {
		return nil, fmt.Errorf("No anime with id: %d", id)
	}

	return manga, nil
}

// Update is used for updating elements in DB
func (db *DB) Update(ctx context.Context, content *models.Content) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.data[content.ID] = content
	return nil
}

// Delete is used for deleting elements by id in DB
func (db *DB) Delete(ctx context.Context, id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	delete(db.data, id)
	return nil
}
