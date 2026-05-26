# AGENTS

Repository guidance for coding agents working in this project.

## Project Layout

- Executable entrypoints:
  - `cmd/backend`
  - `cmd/daemon`
- Application internals:
  - `internal/envutil`
  - `internal/indexer`
  - `internal/storage/postgres`
- Build assets:
  - `build/docker/Dockerfile`
- Runtime and chain configs:
  - `configs/umbrella/mainnet.json`
- DB migrations:
  - `migrations/*.up.sql` and `migrations/*.down.sql`
- Architecture decisions:
  - `docs/adr/`

## Required Conventions

- Keep executables under `cmd/*`; do not reintroduce `apps/*` for Go binaries.
- Keep reusable command wrappers in `scripts/` and call them from `Makefile` targets.
- Treat `tmp/` as generated/runtime output only (pid files, logs, local build artifacts).
- Keep business logic in `internal/*`, not in `cmd/*`.
- When architecture or structure changes, add or update an ADR in `docs/adr/`.

## Common Commands

- Start Docker dev mode (hot reload):
  - `make docker-dev`
- Tail backend and daemon logs:
  - `make docker-read-logs`
- Stop Docker dev services:
  - `make docker-dev-stop`
- Reset Docker volumes (fresh DB state):
  - `make docker-reset-volumes`
- Build and start Docker services:
  - `make docker-build`
- Run migrations in Docker:
  - `make docker-migrations`

## Validation Checklist

Before finishing a change:
- run `go build ./cmd/backend ./cmd/daemon` for layout-sensitive updates
- run `go test ./...` for behavioral changes
- ensure no stale `apps/backend` or `apps/daemon` references remain
- update `README.md` and ADRs when structure or operational commands change
