#!/usr/bin/env sh
set -eu

CGO_ENABLED=0 go build ./cmd/backend ./cmd/daemon