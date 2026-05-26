package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"time"

	"1-task/internal/storage/postgres/queries"
)

// SaveIndexerState stores indexer checkpoint by name.
func (r *Repository) SaveIndexerState(ctx context.Context, name string, lastBlock uint64, processedAt time.Time) error {
	if lastBlock > math.MaxInt64 {
		return fmt.Errorf("save indexer_state: last block overflows int64: %d", lastBlock)
	}

	err := r.q.UpsertIndexerState(ctx, queries.UpsertIndexerStateParams{
		Name:               name,
		LastProcessedBlock: int64(lastBlock),
		LastProcessedTime:  processedAt,
	})
	if err != nil {
		return fmt.Errorf("save indexer_state: %w", err)
	}
	return nil
}

// GetIndexerState returns checkpoint for an indexer name.
func (r *Repository) GetIndexerState(ctx context.Context, name string) (IndexerState, error) {
	state, err := r.q.GetIndexerState(ctx, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return IndexerState{}, sql.ErrNoRows
		}
		return IndexerState{}, fmt.Errorf("get indexer_state: %w", err)
	}

	if state.LastProcessedBlock < 0 {
		return IndexerState{}, fmt.Errorf("get indexer_state: negative last_processed_block for %q", name)
	}

	return IndexerState{
		Name:               state.Name,
		LastProcessedBlock: uint64(state.LastProcessedBlock),
		LastProcessedTime:  state.LastProcessedTime,
		UpdatedAt:          state.UpdatedAt,
	}, nil
}
