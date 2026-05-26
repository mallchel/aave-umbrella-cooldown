# 0002. Migrate to Go Standard Project Layout

Date: 2026-05-26
Status: accepted

## Context

The repository started with executable entrypoints under `apps/`:
- `apps/backend`
- `apps/daemon`

This worked functionally but diverged from the conventional Go project structure used by many teams and tools. We wanted a layout that is easier to navigate for new contributors and more aligned with standard Go project practices.

## Decision

We migrated executable entrypoints to `cmd/` and standardized related tooling paths.

Changes made:
- moved backend entrypoint from `apps/backend/main.go` to `cmd/backend/main.go`
- moved daemon entrypoint from `apps/daemon/main.go` to `cmd/daemon/main.go`
- moved Dockerfile to `build/docker/Dockerfile`
- updated compose build references to `build/docker/Dockerfile`
- updated Air configs to build from `cmd/backend` and `cmd/daemon`
- updated Make targets to use scripts under `scripts/`
- introduced `scripts/` for local build/run helpers
- clarified `tmp/` as runtime-generated artifacts only

## Consequences

Positive:
- executable discovery follows common Go conventions (`cmd/*`)
- repository layout is more predictable for contributors and automation
- build and run commands are centralized via scripts and stable Make targets
- Docker assets now live under `build/`, reducing root-level clutter

Trade-offs:
- path updates were required across docs and tooling
- external references to old `apps/*` paths are now invalid

## Follow-up

Potential follow-up improvements:
- add `scripts/README.md` documenting script usage and intent
- add CI checks to validate `go build ./cmd/...` and prevent stale path regressions
- consider future expansion of `build/` for release packaging artifacts if needed
