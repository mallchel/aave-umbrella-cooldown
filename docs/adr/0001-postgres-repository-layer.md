# 0001. Use PostgreSQL Repository Layer

Date: 2026-05-12
Status: accepted

## Context

The architecture requires:
- PostgreSQL schema for queue data
- persistence methods for raw withdraw requests
- indexer checkpoint read/write

## Decision

We document architecture decisions using MADR files under `docs/adr/`.

We keep PostgreSQL as the persistence backend and implement a repository layer in code for:
- `raw_withdraw_flows`
- `indexer_state`

Schema source of truth is the SQL migration:
- `migrations/0001_init.sql`

## Consequences

Positive:
- decision history is explicit and versioned
- future design changes can be tracked in follow-up ADRs
- repository behavior is tied to schema decisions

Trade-offs:
- one more artifact to maintain
- contributors must update ADRs when architecture changes

## Follow-up

Potential next ADRs:
- indexing strategy and partitioning for `raw_withdraw_flows`
- migration strategy for schema-breaking updates
