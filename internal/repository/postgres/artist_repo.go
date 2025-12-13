package postgres

import (
	"context"
	"database/sql"

	"github.com/dmehra2102/spotify-api/internal/domain/errors"
	"github.com/dmehra2102/spotify-api/internal/domain/models"
)

type ArtistRepistory struct {
	db *sql.DB
}

func NewArtistRepository(db *sql.DB) *ArtistRepistory {
	return &ArtistRepistory{db: db}
}

func (r *ArtistRepistory) Create(ctx context.Context, artist *models.Artist) (*models.Artist, error) {
	query := `
		INSERT INTO artists (id, user_id, name, bio, avatar_url, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, user_id, name, bio, avatar_url, followers_count, songs_count, created_at, updated_at
	`

	err := r.db.QueryRowContext(ctx, query,
		artist.ID, artist.UserID, artist.Name, artist.Bio, artist.AvatarURL, artist.CreatedAt, artist.UpdatedAt,
	).Scan(
		&artist.ID, &artist.UserID, &artist.Name, &artist.Bio, &artist.AvatarURL,
		&artist.FollowersCount, &artist.SongsCount, &artist.CreatedAt, &artist.UpdatedAt,
	)

	return artist, err
}

func (r *ArtistRepistory) GetByUserID(ctx context.Context, userID string) (*models.Artist, error) {
	query := `
		SELECT id, user_id, name, bio, avatar_url, followers_count, songs_count, created_at, updated_at FROM artists WHERE user_id = $1
	`

	artist := &models.Artist{}
	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&artist.ID, &artist.UserID, &artist.Name, &artist.Bio, &artist.AvatarURL,
		&artist.FollowersCount, &artist.SongsCount, &artist.CreatedAt, &artist.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.ErrNotFound
	}

	return artist, err
}

func (r *ArtistRepistory) GetByID(ctx context.Context, artistId string) (*models.Artist, error) {
	query := `
		SELECT id, user_id, name, bio, avatar_url, followers_count, songs_count, created_at, updated_at
		FROM artists WHERE id = $1
	`

	artist := &models.Artist{}
	err := r.db.QueryRowContext(ctx, query, artistId).Scan(
		&artist.ID, &artist.UserID, &artist.Name, &artist.Bio, &artist.AvatarURL,
		&artist.FollowersCount, &artist.SongsCount, &artist.CreatedAt, &artist.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.ErrNotFound
	}

	return artist, err
}
