package postgres

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// ListWithdrawFlowsFilter contains optional DB filters for list endpoint.
type ListWithdrawFlowsFilter struct {
	SenderAddress *string
	EventType     *string
	FromTime      *time.Time
	ToTime        *time.Time
	Limit         int
	Offset        int
}

// WithdrawFlowRow is a DB row model for listing withdraw/request events.
type WithdrawFlowRow struct {
	ChainID          int64
	TxHash           string
	LogIndex         int32
	BlockNumber      int64
	BlockTime        time.Time
	SenderAddress    string
	EventType        string
	AmountRaw        string
	AmountNormalized string
	UpdatedAt        time.Time
}

// ListWithdrawFlows returns flow events using DB-level filters and pagination.
func (r *Repository) ListWithdrawFlows(ctx context.Context, filter ListWithdrawFlowsFilter) ([]WithdrawFlowRow, error) {
	var query strings.Builder
	query.WriteString(`
SELECT
  chain_id,
  tx_hash,
  log_index,
  block_number,
  block_time,
  sender_address,
  event_type,
  amount_raw::text,
  amount_normalized::text,
  updated_at
FROM raw_withdraw_flows
WHERE 1=1
`)

	args := make([]any, 0, 8)
	addFilter := func(expr string, value any) {
		args = append(args, value)
		query.WriteString(" AND ")
		fmt.Fprintf(&query, expr, len(args))
	}

	if filter.SenderAddress != nil {
		addFilter("sender_address = $%d", *filter.SenderAddress)
	}
	if filter.EventType != nil {
		addFilter("event_type = $%d", *filter.EventType)
	}
	if filter.FromTime != nil {
		addFilter("block_time >= $%d", *filter.FromTime)
	}
	if filter.ToTime != nil {
		addFilter("block_time <= $%d", *filter.ToTime)
	}

	limit := filter.Limit
	if limit <= 0 {
		limit = 100
	}
	offset := max(filter.Offset, 0)

	args = append(args, limit)
	limitArgPos := len(args)
	args = append(args, offset)
	offsetArgPos := len(args)

	fmt.Fprintf(&query, " ORDER BY block_time DESC, block_number DESC, log_index DESC LIMIT $%d OFFSET $%d", limitArgPos, offsetArgPos)

	rows, err := r.db.QueryContext(ctx, query.String(), args...)
	if err != nil {
		return nil, fmt.Errorf("list raw_withdraw_flows: %w", err)
	}
	defer rows.Close()

	result := make([]WithdrawFlowRow, 0, limit)
	for rows.Next() {
		var row WithdrawFlowRow
		if err := rows.Scan(
			&row.ChainID,
			&row.TxHash,
			&row.LogIndex,
			&row.BlockNumber,
			&row.BlockTime,
			&row.SenderAddress,
			&row.EventType,
			&row.AmountRaw,
			&row.AmountNormalized,
			&row.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan raw_withdraw_flows: %w", err)
		}
		result = append(result, row)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate raw_withdraw_flows: %w", err)
	}

	return result, nil
}
