-- name: GetIndexerState :one
SELECT
  name,
  last_processed_block,
  last_processed_time,
  updated_at
FROM indexer_state
WHERE name = $1;

-- name: UpsertIndexerState :exec
INSERT INTO indexer_state (
  name,
  last_processed_block,
  last_processed_time
) VALUES (
  $1,
  $2,
  $3
)
ON CONFLICT (name) DO UPDATE
SET
  last_processed_block = EXCLUDED.last_processed_block,
  last_processed_time = EXCLUDED.last_processed_time,
  updated_at = NOW();
