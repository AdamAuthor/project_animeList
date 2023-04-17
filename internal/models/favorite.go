package models

import "encoding/json"

type Favorite struct {
	ID      int             `json:"id" db:"id"`
	UserID  int             `json:"userID" db:"userid"`
	AnimeID int             `json:"animeID" db:"animeid"`
	Anime   json.RawMessage `json:"content"`
}
