package auth

import (
	"context"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	RegisterUser(ctx context.Context, username, password string) (int64, error)
	LoginUser(ctx context.Context, username, password string) (int64, error)
}

type SQLRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *SQLRepository {
	return &SQLRepository{db: db}
}

func (r *SQLRepository) RegisterUser(ctx context.Context, username, password string) (int64, error) {
	query := `
		INSERT INTO users (username, password_hash) 
		VALUES ($1, $2)
		RETURNING id
	`

	var id int64
	err := r.db.QueryRowContext(ctx, query, username, password).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("creating user: %w", err)
	}

	return id, nil
}

func (r *SQLRepository) LoginUser(ctx context.Context, username, password string) (int64, error) {
	query := `
		SELECT id, password_hash
		FROM users
		WHERE username = $1
	`

	var id int64
	var hash string

	err := r.db.QueryRowContext(ctx, query, username).Scan(&id, &hash)
	if err != nil {
		return 0, fmt.Errorf("getting user: %w", err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return 0, fmt.Errorf("invalid credentials: %w", err)
	}

	return id, nil
}
