package postgres

import (
	"context"
	"fmt"

	"1-task/internal/storage/postgres/queries"
)

// UpsertWithdrawFlow inserts or updates a flow event by (chain_id, tx_hash, log_index).
func (r *Repository) UpsertWithdrawFlow(ctx context.Context, flow WithdrawFlow) error {
	err := r.q.UpsertWithdrawFlow(ctx, queries.UpsertWithdrawFlowParams{
		ChainID:          flow.ChainID,
		TxHash:           flow.TxHash,
		LogIndex:         flow.LogIndex,
		BlockNumber:      flow.BlockNumber,
		BlockTime:        flow.BlockTime,
		SenderAddress:    flow.SenderAddress,
		EventType:        flow.EventType,
		AmountRaw:        flow.AmountRaw,
		AmountNormalized: flow.AmountNormalized,
		CooldownEndAt:    flow.CooldownEndAt,
	})
	if err != nil {
		return fmt.Errorf("upsert raw_withdraw_flows: %w", err)
	}
	return nil
}
