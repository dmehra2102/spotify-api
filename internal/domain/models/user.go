package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	AvatarURL    string    `json:"avatar_url"`
	Bio          string    `json:"bio"`
	HashPassword string    `json:"hash_password"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserCredentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserProfile struct {
	ID             string    `json:"id"`
	Email          string    `json:"email"`
	Name           string    `json:"name"`
	AvatarURL      string    `json:"avatar_url"`
	Bio            string    `json:"bio"`
	FollowersCount int       `json:"followers_count"`
	FollowingCount int       `json:"following_count"`
	SongsCount     int       `json:"songs_count"`
	CreatedAt      time.Time `json:"created_at"`
}

type JWTClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
}

func NewUser(email, name string) *User {
	now := time.Now()
	return &User{
		ID:        uuid.New().String(),
		Email:     email,
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
