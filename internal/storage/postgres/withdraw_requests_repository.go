package postgres

import (
	"context"
	"fmt"

	"gorm.io/gorm/clause"
)

// UpsertWithdrawFlow inserts or updates a flow event by (chain_id, tx_hash, log_index).
func (r *Repository) UpsertWithdrawFlow(ctx context.Context, flow WithdrawFlow) error {
	err := r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "chain_id"}, {Name: "tx_hash"}, {Name: "log_index"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"block_number",
				"block_time",
				"sender_address",
				"event_type",
				"amount_raw",
				"amount_normalized",
				"updated_at",
			}),
		}).
		Create(&flow).Error
	if err != nil {
		return fmt.Errorf("upsert raw_withdraw_flows: %w", err)
	}
	return nil
}
