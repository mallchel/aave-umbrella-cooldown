# Umbrella Cooldown Indexer

Go monorepo with:

- backend service generated from OpenAPI with oapi-codegen and Chi
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
make docker-dev
```

This starts:

- `postgres` with persistent volume
- `backend` with Air hot reload
- `daemon` with Air hot reload

Migrations are executed by `daemon` on startup before indexer logic begins.

### Edit code and see changes

- Backend edits (for example `cmd/backend` or `internal/**`) trigger rebuild/restart automatically.
- Daemon edits (for example `cmd/daemon` or `internal/**`) trigger rebuild/restart automatically.
- JSON config changes also trigger reload.

### View logs

```bash
make docker-read-logs
```

### Test on localhost

```bash
curl -s http://localhost:8888/healthz
```

Swagger UI is served by the same backend process on a separate port:

```bash
open http://localhost:9090
```

Grafana is available at:

```bash
open http://localhost:3030
```

Default local login:

- user: `admin`
- password: `admin`

Grafana datasource and dashboards are provisioned from repository files:

- `configs/grafana/provisioning/datasources/postgres.yml`
- `configs/grafana/provisioning/dashboards/umbrella.yml`
- `configs/grafana/dashboards/*.json`

To save dashboard settings in git, update or export dashboard JSON into `configs/grafana/dashboards/`, then restart the dev stack:

```bash
make docker-dev-stop
make docker-dev
```

### Stop dev mode

```bash
make docker-dev-stop
```

### Reset database volume

```bash
make docker-reset-volumes
```

## 2. Run with Docker locally

This repository includes:

- `build/docker/Dockerfile` with multi-stage builds for backend and daemon
- `docker-compose.yml` with services: `postgres`, `backend`, `daemon`

### Start everything

From repository root:

```bash
make docker-build
```

Migrations run automatically when `daemon` starts (for both `make docker-dev` and `make docker-build`).

Migration files follow golang-migrate naming:

- `migrations/000001_name.up.sql`
- `migrations/000001_name.down.sql`

Add a new migration by creating the next versioned pair (for example `000002_add_index.up.sql` and `000002_add_index.down.sql`), then restart daemon via `make docker-dev` or `make docker-build`.

### Watch logs

```bash
make docker-read-logs
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
make docker-dev-stop
```

### Reset local database volume

```bash
make docker-reset-volumes
```

### Override RPC endpoint

```bash
RPC_URL="https://your-rpc.example" make docker-build
```

## Notes

- Indexer checkpoint is stored in PostgreSQL table indexer_state.
- Startup block is derived from checkpoint state, not an env start block.
- Current indexing mode is cooldown-only event tracking.

## API Documentation

OpenAPI source of truth is stored in YAML:

- `docs/openapi/openapi.yaml`

The backend HTTP server glue and response/request types are generated from that contract with oapi-codegen:

```bash
make openapi-generate
```

This writes:

- `internal/backend/openapi.gen.go`

Generate JSON from YAML with:

```bash
make openapi-json
```

This writes:

- `docs/openapi/openapi.json`

Interactive Swagger UI is served by the backend binary without an extra container:

- `http://localhost:9090`

The API itself remains available on:

- `http://localhost:8888`
