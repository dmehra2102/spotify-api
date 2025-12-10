package models

import (
	"time"

	"github.com/google/uuid"
)

type Follow struct {
	ID         string    `json:"id"`
	FollowerID string    `json:"follower_id"`
	ArtistID   string    `json:"artist_id"`
	CreatedAt  time.Time `json:"created_at"`
}

func NewFollow(followerID, artistID string) *Follow {
	return &Follow{
		ID:         uuid.New().String(),
		FollowerID: followerID,
		ArtistID:   artistID,
		CreatedAt:  time.Now(),
	}
}
