package models

import (
	"time"

	"github.com/google/uuid"
)

type Artist struct {
	ID             string    `json:"id"`
	UserID         string    `json:"user_id"`
	Name           string    `json:"name"`
	Bio            string    `json:"bio"`
	AvatarURL      string    `json:"avatar_url"`
	FollowersCount int       `json:"followers_count"`
	SongsCount     int       `json:"songs_count"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type ArtistInfo struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	AvatarURL      string `json:"avatar_url"`
	FollowersCount int    `json:"followers_count"`
}

type ArtistProfile struct {
	ID             string    `json:"id"`
	UserID         string    `json:"user_id"`
	Name           string    `json:"name"`
	Bio            string    `json:"bio"`
	AvatarURL      string    `json:"avatar_url"`
	FollowersCount int       `json:"followers_count"`
	SongsCount     int       `json:"songs_count"`
	IsFollowing    bool      `json:"is_following"`
	CreatedAt      time.Time `json:"created_at"`
}

func NewArtist(userID, name, bio string) *Artist {
	now := time.Now()
	return &Artist{
		ID:        uuid.New().String(),
		UserID:    userID,
		Name:      name,
		Bio:       bio,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
