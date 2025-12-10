package models

import (
	"time"

	"github.com/google/uuid"
)

type Song struct {
	ID          string    `json:"id"`
	ArtistID    string    `json:"artist_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Duration    int       `json:"duration"`
	FileKey     string    `json:"file_key"`
	FileSize    int64     `json:"file_size"`
	Genre       string    `json:"genre"`
	LikesCount  int       `json:"likes_count"`
	PlaysCount  int       `json:"plays_count"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type SongUpload struct {
	Title       string `json:"title" validate:"required,max=255"`
	Description string `json:"description" validate:"max=1000"`
	Genre       string `json:"genre" validate:"required,max=50"`
	Duration    int    `json:"duration" validate:"required,min=1"`
}

type SongResponse struct {
	ID          string      `json:"id"`
	ArtistID    string      `json:"artist_id"`
	Artist      *ArtistInfo `json:"artist"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Duration    int         `json:"duration"`
	Genre       string      `json:"genre"`
	LikesCount  int         `json:"likes_count"`
	PlaysCount  int         `json:"plays_count"`
	IsLiked     bool        `json:"is_liked"`
	CreatedAt   time.Time   `json:"created_at"`
}

func NewSong(artistID, title, description, genre string, duration int) *Song {
	now := time.Now()
	return &Song{
		ID:          uuid.New().String(),
		ArtistID:    artistID,
		Title:       title,
		Description: description,
		Genre:       genre,
		Duration:    duration,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
