package models

import "encoding/json"

type Favorite struct {
	ID        int             `json:"id" db:"id"`
	UserID    int             `json:"userID" db:"userID"`
	ContentID int             `json:"contentID" db:"contentID"`
	Anime     json.RawMessage `json:"content" db:"anime" `
}
