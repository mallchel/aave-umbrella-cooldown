package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SaveIndexerState stores indexer checkpoint by name.
func (r *Repository) SaveIndexerState(ctx context.Context, name string, lastBlock uint64, processedAt time.Time) error {
	state := IndexerState{
		Name:               name,
		LastProcessedBlock: lastBlock,
		LastProcessedTime:  processedAt,
	}

	err := r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "name"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"last_processed_block",
				"last_processed_time",
				"updated_at",
			}),
		}).
		Create(&state).Error
	if err != nil {
		return fmt.Errorf("save indexer_state: %w", err)
	}
	return nil
}

// GetIndexerState returns checkpoint for an indexer name.
func (r *Repository) GetIndexerState(ctx context.Context, name string) (IndexerState, error) {
	var s IndexerState
	err := r.db.WithContext(ctx).Where("name = ?", name).Take(&s).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return IndexerState{}, sql.ErrNoRows
		}
		return IndexerState{}, fmt.Errorf("get indexer_state: %w", err)
	}
	return s, nil
}
