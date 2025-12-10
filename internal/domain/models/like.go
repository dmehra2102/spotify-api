package models

import (
	"time"

	"github.com/google/uuid"
)

type Like struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	SongID    string    `json:"song_id"`
	CreatedAt time.Time `json:"created_at"`
}

func NewLike(userID, songID string) *Like {
	return &Like{
		ID:        uuid.New().String(),
		UserID:    userID,
		SongID:    songID,
		CreatedAt: time.Now(),
	}
}
