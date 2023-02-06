package repositories

import (
	"context"
	"database/sql"
	"grpc-unary-stream/internal/entity"
)

type NoteRepository interface {
	FindById(ctx context.Context, id int64) (*entity.Note, error)
}

type note struct {
	db *sql.DB
}

// NewNoteRepository is function to create new note repository
func NewNoteRepository(db *sql.DB) NoteRepository {
	return &note{db: db}
}

func (r *note) FindById(ctx context.Context, id int64) (*entity.Note, error) {
	q := `SELECT 
			id,
			title,
			description,
			created_at,
			updated_at
		 FROM notes WHERE id = ?`

	var note entity.Note

	if err := r.db.QueryRowContext(ctx, q, id).Scan(
		&note.ID,
		&note.Title,
		&note.Description,
		&note.CreatedAt,
		&note.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &note, nil
}
