package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"1-task/internal/storage/postgres/queries"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// Repository provides PostgreSQL access for queue and indexer data.
type Repository struct {
	db *sql.DB
	q  *queries.Queries
}

// New creates a new PostgreSQL repository.
func New(ctx context.Context, dsn string) (*Repository, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("open postgres: %w", err)
	}

	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	return &Repository{db: db, q: queries.New(db)}, nil
}

// Close closes the underlying DB pool.
func (r *Repository) Close() error {
	if r == nil || r.db == nil {
		return nil
	}
	return r.db.Close()
}
