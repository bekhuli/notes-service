package note

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository interface {
	CreateNote(ctx context.Context, userID int64, title, content string) (int64, error)
	GetNotes(ctx context.Context, userID int64) ([]Note, error)
	UpdateNote(ctx context.Context, id int64, title, content string) (*Note, error)
	DeleteNote(ctx context.Context, id int64) error
}

type SQLRepository struct {
	db *sql.DB
}

func NewNoteRepository(db *sql.DB) *SQLRepository {
	return &SQLRepository{db: db}
}

func (r *SQLRepository) CreateNote(ctx context.Context, userID int64, title, content string) (int64, error) {
	query := `
		INSERT INTO notes (user_id, title, content)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	var id int64

	err := r.db.QueryRowContext(ctx, query, userID, title, content).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("creating note: %w", err)
	}

	return id, nil
}

type Note struct {
	ID      int64
	UserID  int64
	Title   string
	Content string
}

func (r *SQLRepository) GetNotes(ctx context.Context, userID int64) ([]Note, error) {
	query := `
		SELECT id, title, content
		FROM notes
		WHERE user_id = $1
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("error getting notes: %w", err)
	}
	defer rows.Close()

	var notes []Note
	for rows.Next() {
		var n Note
		n.UserID = userID
		if err = rows.Scan(&n.ID, &n.Title, &n.Content); err != nil {
			return nil, err
		}
		notes = append(notes, n)
	}
	return notes, nil
}

func (r *SQLRepository) UpdateNote(ctx context.Context, id int64, title, content string) (*Note, error) {
	query := `
		UPDATE notes
		SET title = $1, content = $2
		WHERE id = $3
	`

	_, err := r.db.ExecContext(ctx, query, title, content, id)
	if err != nil {
		return &Note{}, fmt.Errorf("updating note: %w", err)
	}

	query = `
		SELECT user_id
		FROM notes
		WHERE id = $1
	`

	var userID int64
	err = r.db.QueryRowContext(ctx, query, id).Scan(&userID)
	if err != nil {
		return &Note{}, fmt.Errorf("getting note's user_id: %w", err)
	}

	return &Note{id, userID, title, content}, nil
}

func (r *SQLRepository) DeleteNote(ctx context.Context, id int64) error {
	query := `
		DELETE FROM notes
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("deleting note: %w", err)
	}

	return nil
}
