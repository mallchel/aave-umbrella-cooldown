.PHONY: help \
	docker-dev docker-read-logs docker-dev-stop docker-reset-volumes docker-build docker-migrations \
	build run-backend run-daemon daemon daemon-stop daemon-status

help:
	@echo "Available targets:"
	@echo "  make docker-dev       - start docker services in dev mode"
	@echo "  make docker-read-logs - tail backend and daemon logs"
	@echo "  make docker-dev-stop  - stop docker dev services"
	@echo "  make docker-reset-volumes - stop services and remove volumes"
	@echo "  make docker-build     - build and start docker services"
	@echo "  make docker-migrations - run migrations container"
	@echo "  make build         - build backend and daemon"
	@echo "  make run-backend   - run backend in foreground"
	@echo "  make run-daemon    - run daemon in foreground"
	@echo "  make daemon        - start daemon in background (go-daemon)"
	@echo "  make daemon-stop   - stop background daemon"
	@echo "  make daemon-status - show daemon process status"

docker-dev:
	docker compose -f docker-compose.yml -f docker-compose.dev.yml up -d

docker-read-logs:
	docker compose -f docker-compose.yml -f docker-compose.dev.yml logs -f backend daemon

docker-dev-stop:
	docker compose -f docker-compose.yml -f docker-compose.dev.yml down

docker-reset-volumes:
	docker compose -f docker-compose.yml -f docker-compose.dev.yml down -v

docker-build:
	docker compose up --build -d

docker-migrations:
	docker compose run --rm migrate

build:
	sh ./scripts/build.sh

run-backend:
	sh ./scripts/run-backend.sh

run-daemon:
	sh ./scripts/run-daemon.sh

daemon:
	sh ./scripts/start-daemon.sh

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
