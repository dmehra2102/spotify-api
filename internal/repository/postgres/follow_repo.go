package postgres

import (
	"context"
	"database/sql"

	"github.com/dmehra2102/spotify-api/internal/domain/models"
)

type FollowRepository struct {
	db *sql.DB
}

func (r *FollowRepository) Create(ctx context.Context, follow *models.Follow) (*models.Follow, error) {
	query := `
		INSERT INTO follows (id, follower_id, artist_id, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id, follower_id, artist_id, created_at
	`

	err := r.db.QueryRowContext(ctx, query,
		follow.ID, follow.FollowerID, follow.ArtistID, follow.CreatedAt,
	).Scan(
		&follow.ID, &follow.FollowerID, &follow.ArtistID, &follow.CreatedAt,
	)

	return follow, err
}

func (r *FollowRepository) Delete(ctx context.Context, followerID, artistID string) error {
	query := `DELETE FROM follows WHERE follower_id = $1 AND artist_id = $2`
	_, err := r.db.ExecContext(ctx, query, followerID, artistID)
	return err
}

func (r *FollowRepository) IsFollowing(ctx context.Context, followerID, artistID string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM follows WHERE follower_id = $1 AND artist_id = $2)`
	var exists bool
	err := r.db.QueryRowContext(ctx, query, followerID, artistID).Scan(&exists)
	return exists, err
}
