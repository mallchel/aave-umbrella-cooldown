-- name: UpsertWithdrawFlow :exec
INSERT INTO raw_withdraw_flows (
  chain_id,
  tx_hash,
  log_index,
  block_number,
  block_time,
  sender_address,
  event_type,
  amount_raw,
  amount_normalized,
  cooldown_end_at
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7,
  $8,
  $9,
  $10
)
ON CONFLICT (chain_id, tx_hash, log_index) DO UPDATE
SET
  block_number = EXCLUDED.block_number,
  block_time = EXCLUDED.block_time,
  sender_address = EXCLUDED.sender_address,
  event_type = EXCLUDED.event_type,
  amount_raw = EXCLUDED.amount_raw,
  amount_normalized = EXCLUDED.amount_normalized,
  cooldown_end_at = EXCLUDED.cooldown_end_at,
  updated_at = NOW();
