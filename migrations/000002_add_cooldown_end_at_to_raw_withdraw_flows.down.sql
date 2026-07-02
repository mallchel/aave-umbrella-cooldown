BEGIN;

DROP INDEX IF EXISTS idx_raw_withdraw_flows_cooldown_end_at;

ALTER TABLE raw_withdraw_flows
DROP COLUMN IF EXISTS cooldown_end_at;

COMMIT;
