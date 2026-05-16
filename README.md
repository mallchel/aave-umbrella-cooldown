# Umbrella Cooldown Indexer

Go monorepo with:
- backend service (Chi)
- daemon indexer (go-daemon)
- PostgreSQL storage for cooldown events

The daemon indexes cooldown transactions and stores them in PostgreSQL.

## Requirements

- Go 1.26+
- PostgreSQL 14+
- Ethereum RPC endpoint

## Project Structure

- apps/backend - HTTP backend
- apps/daemon - background indexer daemon
- internal/indexer - on-chain indexing logic
- internal/storage/postgres - repository layer
- migrations - SQL migrations
- configs/umbrella/mainnet.json - chain and event config

## 1. Prepare PostgreSQL

### Create user and database

Run in psql:

```sql
CREATE ROLE umbrella_user WITH LOGIN PASSWORD 'umbrella_pass';
CREATE DATABASE umbrella_db OWNER umbrella_user;
GRANT ALL PRIVILEGES ON DATABASE umbrella_db TO umbrella_user;
```

### Apply migration

From repository root:

```bash
psql "postgresql://umbrella_user:umbrella_pass@localhost:5432/umbrella_db?sslmode=disable" -f migrations/0001_init.sql
```

### Verify tables

```bash
psql "postgresql://umbrella_user:umbrella_pass@localhost:5432/umbrella_db?sslmode=disable" -c "\\dt"
```

Expected tables include:
- raw_withdraw_requests
- indexer_state

## 2. Configure environment

Export required variables:

```bash
export POSTGRES_DSN="host=localhost port=5432 user=umbrella_user password=umbrella_pass dbname=umbrella_db sslmode=disable"
export RPC_URL="https://ethereum-rpc.publicnode.com"
```

Optional variables:

```bash
export UMBRELLA_CONFIG_PATH="./configs/umbrella/mainnet.json"
export FINALITY_DEPTH="12"
export INDEXER_BATCH_BLOCK_RANGE="2000"
```

## 3. Build and run

From repository root:

```bash
make build
```

Run backend in foreground:

```bash
make run-backend
```

Run daemon in foreground:

```bash
make run-daemon
```

Run daemon as background process (go-daemon):

```bash
make daemon
```

Stop or inspect daemon:

```bash
make daemon-stop
make daemon-status
```

## 4. Health check

When backend is running:

```bash
curl -s http://localhost:8080/healthz
```

## Notes

- Indexer checkpoint is stored in PostgreSQL table indexer_state.
- Startup block is derived from checkpoint state, not an env start block.
- Current indexing mode is cooldown-only event tracking.
