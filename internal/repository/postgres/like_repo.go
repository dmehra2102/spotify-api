package postgres

import (
	"context"
	"database/sql"

	"github.com/dmehra2102/spotify-api/internal/domain/models"
)

type LikeRepository struct {
	db *sql.DB
}

func NewLikeRepository(db *sql.DB) *LikeRepository {
	return &LikeRepository{db: db}
}

func (r *LikeRepository) Create(ctx context.Context, like *models.Like) (*models.Like, error) {
	query := `
		INSERT INTO likes (id, user_id, song_id, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id, user_id, song_id, created_at
	`

	err := r.db.QueryRowContext(ctx, query, like.ID, like.UserID, like.SongID, like.CreatedAt).Scan(
		&like.ID, &like.UserID, &like.SongID, &like.CreatedAt,
	)

	return like, err
}

func (r *LikeRepository) Delete(ctx context.Context, userID, songID string) error {
	query := `DELETE FROM likes WHERE user_id = $1 AND song_id = $2`
	_, err := r.db.ExecContext(ctx, query, userID, songID)
	return err
}

func (r *LikeRepository) IsLiked(ctx context.Context, userID, songID string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM likes WHERE user_id = $1 AND song_id = $2)`
	var exists bool
	err := r.db.QueryRowContext(ctx, query, userID, songID).Scan(&exists)
	return exists, err
}
