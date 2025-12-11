package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/dmehra2102/spotify-api/internal/domain/errors"
	"github.com/dmehra2102/spotify-api/internal/domain/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, email, name, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := models.NewUser(email, name)

	query := `
		INSERT INTO users (id, email, hash_password, name, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, email, name, created_at, updated_at
	`

	err = r.db.QueryRowContext(ctx, query,
		user.ID, email, string(hashedPassword), name, user.CreatedAt, user.UpdatedAt,
	).Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT id, email, name, avatar_url, bio, created_at, updated_at FROM users WHERE email = $1`

	user := &models.User{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID, &user.Email, &user.Name, &user.AvatarURL, &user.Bio, &user.CreatedAt, &user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetByID(ctx context.Context, userID string) (*models.User, error) {
	query := `SELECT id, email, name, avatar_url, bio, created_at, updated_at FROM users WHERE id = $1`

	user := &models.User{}
	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&user.ID, &user.Email, &user.Name, &user.AvatarURL, &user.Bio, &user.CreatedAt, &user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) VerifyPassword(ctx context.Context, email, password string) (string, error) {
	query := `SELECT id, hash_password FROM users WHERE email = $1`

	var userID, passwordHash string
	err := r.db.QueryRowContext(ctx, query, email).Scan(&userID, &passwordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)); err != nil {
		return "", nil
	}

	return userID, nil
}

func (r *UserRepository) Update(ctx context.Context, userID string, name, bio, avatarURL string) (*models.User, error) {
	query := `
		UPDATE users SET name = $1, bio = $2, avatar_url = $3, updated_at = $4
		WHERE id = $5
		RETURNING id, email, name, avatar_url, bio, created_at, updated_at
	`

	user := &models.User{}
	err := r.db.QueryRowContext(
		ctx, query,
		name, bio, avatarURL, time.Now(), userID,
	).Scan(
		&user.ID, &user.Email, &user.Name, &user.AvatarURL, &user.Bio, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
