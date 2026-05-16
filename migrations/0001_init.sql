BEGIN;

CREATE TABLE IF NOT EXISTS raw_withdraw_requests (
    id BIGSERIAL PRIMARY KEY,
    chain_id BIGINT NOT NULL,
    tx_hash TEXT NOT NULL,
    log_index INTEGER NOT NULL,
    block_number BIGINT NOT NULL,
    block_time TIMESTAMPTZ NOT NULL,
    user_address TEXT NOT NULL,
    asset_address TEXT NOT NULL,
    asset_symbol TEXT NOT NULL,
    amount_raw NUMERIC(78, 0) NOT NULL,
    amount_normalized NUMERIC(38, 18) NOT NULL,
    amount_usd NUMERIC(38, 18) NOT NULL,
    cooldown_start_time TIMESTAMPTZ NOT NULL,
    withdrawable_from TIMESTAMPTZ NOT NULL,
    withdrawable_until TIMESTAMPTZ NOT NULL,
    status TEXT NOT NULL,
    withdraw_tx_hash TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT raw_withdraw_requests_event_unique UNIQUE (chain_id, tx_hash, log_index)
);

CREATE INDEX IF NOT EXISTS idx_raw_withdraw_requests_block_number ON raw_withdraw_requests (block_number);
CREATE INDEX IF NOT EXISTS idx_raw_withdraw_requests_status ON raw_withdraw_requests (status);
CREATE INDEX IF NOT EXISTS idx_raw_withdraw_requests_tx_hash ON raw_withdraw_requests (tx_hash);
CREATE INDEX IF NOT EXISTS idx_raw_withdraw_requests_withdrawable_from ON raw_withdraw_requests (withdrawable_from);
CREATE INDEX IF NOT EXISTS idx_raw_withdraw_requests_withdrawable_until ON raw_withdraw_requests (withdrawable_until);

CREATE TABLE IF NOT EXISTS indexer_state (
    name TEXT PRIMARY KEY,
    last_processed_block BIGINT NOT NULL,
    last_processed_time TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

COMMIT;
