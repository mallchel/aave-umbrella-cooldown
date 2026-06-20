# 0003. Generate HTTP Server From OpenAPI

Date: 2026-06-18
Status: accepted

## Context

The backend already had an OpenAPI description under `docs/openapi/`, but HTTP routes and response models were wired manually in `internal/backend`. This made the OpenAPI document useful as documentation, but not as the source of truth for the server surface.

The implementation requirement is to generate HTTP handlers from Swagger/OpenAPI with `oapi-codegen` and serve Swagger UI from the backend service itself on a separate port without adding another container.

## Decision

We use `docs/openapi/openapi.yaml` as the source of truth for the backend HTTP API.

Generated artifacts:
- `internal/backend/openapi.gen.go` is generated with `oapi-codegen` using `types,chi-server`
- `docs/openapi/openapi.json` remains generated from the YAML spec for JSON consumers

Runtime behavior:
- API routes are registered through the generated Chi server wrapper
- `internal/backend.Server` implements the generated `ServerInterface`
- the backend binary serves API traffic on `HTTP_ADDR`, default `:8888`
- the same backend binary serves Swagger UI and embedded OpenAPI specs on `SWAGGER_ADDR`, default `:9090`

## Consequences

Positive:
- server route wiring now follows the OpenAPI contract
- request parameter binding uses generated code instead of hand-parsing every query parameter from `http.Request`
- Swagger UI is available without a separate container
- API and docs can be started with the same backend process

Trade-offs:
- generated code must be refreshed when `docs/openapi/openapi.yaml` changes
- reviewers need to check generated diffs separately from hand-written handler logic

## Follow-up

Potential follow-up improvements:
- add CI validation that `make openapi-generate` produces no diff
- add contract tests for generated route registration and parameter validation