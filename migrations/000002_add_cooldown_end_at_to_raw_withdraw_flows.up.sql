BEGIN;

ALTER TABLE raw_withdraw_flows
ADD COLUMN IF NOT EXISTS cooldown_end_at BIGINT NOT NULL DEFAULT 0;

CREATE INDEX IF NOT EXISTS idx_raw_withdraw_flows_cooldown_end_at
    ON raw_withdraw_flows (cooldown_end_at);

COMMIT;
