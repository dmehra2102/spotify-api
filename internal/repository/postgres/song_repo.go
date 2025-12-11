package postgres

import (
	"context"
	"database/sql"

	"github.com/dmehra2102/spotify-api/internal/domain/errors"
	"github.com/dmehra2102/spotify-api/internal/domain/models"
)

type SongRepository struct {
	db *sql.DB
}

func NewSongRepository(db *sql.DB) *SongRepository {
	return &SongRepository{db: db}
}

func (r *SongRepository) Create(ctx context.Context, song *models.Song) (*models.Song, error) {
	query := `
		INSERT INTO songs (id, artist_id, title, description, duration, file_key, file_size, genre, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id, artist_id, title, description, duration, file_key, file_size, genre, likes_count, plays_count, created_at, updated_at
	`

	err := r.db.QueryRowContext(ctx, query,
		song.ID, song.ArtistID, song.Title, song.Description, song.Duration, song.FileKey, song.FileSize, song.Genre, song.CreatedAt, song.UpdatedAt,
	).Scan(
		&song.ID, &song.ArtistID, &song.Title, &song.Description, &song.Duration,
		&song.FileKey, &song.FileSize, &song.Genre, &song.LikesCount, &song.PlaysCount,
		&song.CreatedAt, &song.UpdatedAt,
	)

	return song, err
}

func (r *SongRepository) GetByID(ctx context.Context, songID string) (*models.Song, error) {
	query := `
		SELECT id, artist_id, title, description, duration, file_key, file_size, genre, likes_count, plays_count, created_at, updated_at
		FROM songs WHERE id = $1
	`

	song := &models.Song{}
	err := r.db.QueryRowContext(ctx, query, songID).Scan(
		&song.ID, &song.ArtistID, &song.Title, &song.Description, &song.Duration,
		&song.FileKey, &song.FileSize, &song.Genre, &song.LikesCount, &song.PlaysCount,
		&song.CreatedAt, &song.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.ErrNotFound
	}

	return song, nil
}

func (r *SongRepository) List(ctx context.Context, limit, offset int) ([]*models.Song, error) {
	query := `
		SELECT id, artist_id, title, description, duration, file_key, file_size, genre, likes_count, plays_count, created_at, updated_at
		FROM songs ORDER BY created_at DESC LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []*models.Song
	for rows.Next() {
		song := &models.Song{}
		err := rows.Scan(
			&song.ID, &song.ArtistID, &song.Title, &song.Description, &song.Duration,
			&song.FileKey, &song.FileSize, &song.Genre, &song.LikesCount, &song.PlaysCount,
			&song.CreatedAt, &song.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}

	return songs, nil
}

func (r *SongRepository) GetTrending(ctx context.Context, limit int) ([]*models.Song, error) {
	query := `
		SELECT id, artist_id, title, description, duration, file_key, file_size, genre, likes_count, plays_count, created_at, updated_at
		FROM songs ORDER BY likes_count DESC, plays_count DESC LIMIT $1
	`

	rows, err := r.db.QueryContext(ctx, query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []*models.Song
	for rows.Next() {
		song := &models.Song{}
		err := rows.Scan(
			&song.ID, &song.ArtistID, &song.Title, &song.Description, &song.Duration,
			&song.FileKey, &song.FileSize, &song.Genre, &song.LikesCount, &song.PlaysCount,
			&song.CreatedAt, &song.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}

	return songs, nil
}

func (r *SongRepository) IncrementPlaysCount(ctx context.Context, songID string) error {
	query := `UPDATE songs SET plays_count = plays_count + 1 WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, songID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.ErrNotFound
	}

	return err
}

func (r *SongRepository) Delete(ctx context.Context, songID string) error {
	query := `DELETE FROM songs WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, songID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.ErrNotFound
	}

	return err
}
