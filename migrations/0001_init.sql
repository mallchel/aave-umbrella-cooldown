BEGIN;

CREATE TABLE IF NOT EXISTS raw_withdraw_flows (
    id BIGSERIAL PRIMARY KEY,
    chain_id BIGINT NOT NULL,
    tx_hash TEXT NOT NULL,
    log_index INTEGER NOT NULL,
    block_number BIGINT NOT NULL,
    block_time TIMESTAMPTZ NOT NULL,
    sender_address TEXT NOT NULL,
    event_type TEXT NOT NULL,
    amount_raw NUMERIC(78, 0) NOT NULL,
    amount_normalized NUMERIC(38, 18) NOT NULL,
    amount_usdt NUMERIC(38, 18) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT raw_withdraw_flows_event_unique UNIQUE (chain_id, tx_hash, log_index)
);

CREATE INDEX IF NOT EXISTS idx_raw_withdraw_flows_block_number ON raw_withdraw_flows (block_number);
CREATE INDEX IF NOT EXISTS idx_raw_withdraw_flows_block_time ON raw_withdraw_flows (block_time);
CREATE INDEX IF NOT EXISTS idx_raw_withdraw_flows_sender_address ON raw_withdraw_flows (sender_address);
CREATE INDEX IF NOT EXISTS idx_raw_withdraw_flows_event_type ON raw_withdraw_flows (event_type);

CREATE TABLE IF NOT EXISTS indexer_state (
    name TEXT PRIMARY KEY,
    last_processed_block BIGINT NOT NULL,
    last_processed_time TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

COMMIT;
