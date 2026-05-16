package postgres

import (
	"context"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Repository provides PostgreSQL access for queue and indexer data.
type Repository struct {
	db *gorm.DB
}

// New creates a new PostgreSQL repository.
func New(ctx context.Context, dsn string) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open postgres: %w", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("open sql db from gorm: %w", err)
	}

	if err := sqlDB.PingContext(ctx); err != nil {
		_ = sqlDB.Close()
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	return &Repository{db: db}, nil
}

// Close closes the underlying DB pool.
func (r *Repository) Close() error {
	if r == nil || r.db == nil {
		return nil
	}
	sqlDB, err := r.db.DB()
	if err != nil {
		return fmt.Errorf("get sql db from gorm: %w", err)
	}
	return sqlDB.Close()
}
