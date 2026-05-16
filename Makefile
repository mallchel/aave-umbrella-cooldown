.PHONY: help fmt tidy build test run-backend run-daemon daemon daemon-stop daemon-status

help:
	@echo "Available targets:"
	@echo "  make fmt           - format Go sources"
	@echo "  make tidy          - tidy go modules"
	@echo "  make build         - build backend and daemon"
	@echo "  make test          - run tests"
	@echo "  make run-backend   - run backend in foreground"
	@echo "  make run-daemon    - run daemon in foreground"
	@echo "  make daemon        - start daemon in background (go-daemon)"
	@echo "  make daemon-stop   - stop background daemon"
	@echo "  make daemon-status - show daemon process status"

build:
	go build ./apps/backend ./apps/daemon

run-backend:
	go run ./apps/backend

run-daemon:
	DAEMON_FOREGROUND=1 go run ./apps/daemon

daemon:
	mkdir -p ./tmp
	go run ./apps/daemon

daemon-stop:
	@if [ -f ./tmp/umbrella-daemon.pid ]; then \
		kill -TERM "$$(cat ./tmp/umbrella-daemon.pid)"; \
		echo "daemon stopped"; \
	else \
		echo "pid file not found: ./tmp/umbrella-daemon.pid"; \
	fi

daemon-status:
	@if [ -f ./tmp/umbrella-daemon.pid ]; then \
		PID="$$(cat ./tmp/umbrella-daemon.pid)"; \
		ps -p "$$PID"; \
	else \
		echo "pid file not found: ./tmp/umbrella-daemon.pid"; \
	fi
