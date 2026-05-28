# Umbrella Cooldown Indexer

Go monorepo with:

- backend service (Chi)
- daemon indexer (go-daemon)
- PostgreSQL storage for cooldown events

The daemon indexes cooldown transactions and stores them in PostgreSQL.

## Requirements

- Docker
- Docker Compose

## Project Structure

- cmd/backend - HTTP backend entrypoint
- cmd/daemon - background indexer entrypoint
- internal/indexer - on-chain indexing logic
- internal/storage/postgres - repository layer
- build/docker/Dockerfile - multi-stage backend/daemon images
- scripts - local build/run helpers used by Make targets
- migrations - SQL migrations
- configs/umbrella/mainnet.json - chain and event config
- tmp - runtime-generated artifacts (PID files, daemon logs, Air binaries)

## 1. Docker dev mode (hot reload)

For day-to-day backend development, use the compose override file so source changes are bind-mounted and services auto-restart on code edits.

### Start dev mode

```bash
docker compose -f docker-compose.yml -f docker-compose.dev.yml up -d
```

This starts:

- `postgres` with persistent volume
- `migrate` one-shot migration container
- `backend` with Air hot reload
- `daemon` with Air hot reload

### Edit code and see changes

- Backend edits (for example `cmd/backend` or `internal/**`) trigger rebuild/restart automatically.
- Daemon edits (for example `cmd/daemon` or `internal/**`) trigger rebuild/restart automatically.
- JSON config changes also trigger reload.

### View logs

```bash
docker compose -f docker-compose.yml -f docker-compose.dev.yml logs -f backend daemon
```

### Test on localhost

```bash
curl -s http://localhost:8888/healthz
```

### Stop dev mode

```bash
docker compose -f docker-compose.yml -f docker-compose.dev.yml down
```

### Reset database volume

```bash
docker compose -f docker-compose.yml -f docker-compose.dev.yml down -v
```

## 2. Run with Docker locally

This repository includes:

- `build/docker/Dockerfile` with multi-stage builds for backend and daemon
- `docker-compose.yml` with services: `postgres`, `migrate`, `backend`, `daemon`

### Start everything

From repository root:

```bash
docker compose up --build -d
```

### Run migrations only

```bash
docker compose run --rm migrate
```

Migration files follow golang-migrate naming:

- `migrations/000001_name.up.sql`
- `migrations/000001_name.down.sql`

Add a new migration by creating the next versioned pair (for example `000002_add_index.up.sql` and `000002_add_index.down.sql`), then run:

```bash
docker compose run --rm migrate
```

### Watch logs

```bash
docker compose logs -f migrate backend daemon
```

### Check backend

```bash
curl -s http://localhost:8888/healthz
```

### Query withdraw flows (GET + DB filters)

```bash
curl -s "http://localhost:8888/withdraw-flows?event_type=request&limit=20"
```

### Query daily series raw JSON data

```bash
curl -s "http://localhost:8888/daily-series-data"
```

Returns a JSON array of daily points with fields:

- day
- requested
- withdrawn
- request_count

Available query params:

- sender_address (string)
- event_type (request|withdraw)
- from_time (RFC3339)
- to_time (RFC3339)
- limit (1-500, default 100)
- offset (>=0, default 0)

### Verify tables from host (optional)

```bash
psql "postgresql://umbrella_user:umbrella_pass@localhost:5432/umbrella_db?sslmode=disable" -c "\\dt"
```

Expected tables include:

- raw_withdraw_flows
- indexer_state

### Stop stack

```bash
docker compose down
```

### Reset local database volume

```bash
docker compose down -v
```

### Override RPC endpoint

```bash
RPC_URL="https://your-rpc.example" docker compose up --build -d daemon
```

## Notes

- Indexer checkpoint is stored in PostgreSQL table indexer_state.
- Startup block is derived from checkpoint state, not an env start block.
- Current indexing mode is cooldown-only event tracking.
