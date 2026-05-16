package postgres

import (
	"context"
	"fmt"

	"gorm.io/gorm/clause"
)

// UpsertWithdrawRequest inserts or updates a withdraw request by (chain_id, tx_hash, log_index).
func (r *Repository) UpsertWithdrawRequest(ctx context.Context, req WithdrawRequest) error {
	err := r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "chain_id"}, {Name: "tx_hash"}, {Name: "log_index"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"block_number",
				"block_time",
				"user_address",
				"asset_symbol",
				"amount_raw",
				"amount_normalized",
				"amount_usd",
				"cooldown_start_time",
				"withdrawable_from",
				"withdrawable_until",
				"status",
				"withdraw_tx_hash",
				"updated_at",
			}),
		}).
		Create(&req).Error
	if err != nil {
		return fmt.Errorf("upsert raw_withdraw_requests: %w", err)
	}
	return nil
}
