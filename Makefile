.PHONY: help \
	docker-dev docker-read-logs docker-dev-stop docker-reset-volumes docker-build docker-migrations \
	build run-backend daemon-foreground daemon daemon-stop daemon-status sqlc-generate openapi-json openapi-check openapi-docs

help:
	@echo "Available targets:"
	@echo "  make docker-dev       - start docker services in dev mode"
	@echo "  make docker-read-logs - tail backend and daemon logs"
	@echo "  make docker-dev-stop  - stop docker dev services"
	@echo "  make docker-reset-volumes - stop services and remove volumes"
	@echo "  make docker-build     - build and start docker services"
	@echo "  make docker-migrations - run migrations container"
	@echo "  make build         - build backend and daemon"
	@echo "  make sqlc-generate - generate sqlc code"
	@echo "  make openapi-json  - generate OpenAPI JSON from docs/openapi/openapi.yaml"
	@echo "  make openapi-check - verify docs/openapi/openapi.json matches docs/openapi/openapi.yaml"
	@echo "  make openapi-docs  - serve Swagger UI at http://localhost:9090"
	@echo "  make run-backend   - run backend in foreground"
	@echo "  make daemon-foreground    - run daemon in foreground"
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

sqlc-generate:
	go run github.com/sqlc-dev/sqlc/cmd/sqlc@v1.29.0 generate -f internal/storage/postgres/queries/sqlc.yaml

openapi-json:
	docker run --rm -v "$$(pwd):/workdir" -w /workdir mikefarah/yq -o=json '.' docs/openapi/openapi.yaml > docs/openapi/openapi.json

openapi-check:
	@$(MAKE) openapi-json
	@git diff --exit-code -- docs/openapi/openapi.json

openapi-docs:
	@$(MAKE) openapi-json
	docker run --rm -p 9090:8080 -e SWAGGER_JSON=/spec/openapi.json -v "$$(pwd)/docs/openapi:/spec" swaggerapi/swagger-ui

run-backend:
	sh ./scripts/run-backend.sh

daemon-foreground:
	sh ./scripts/daemon-foreground.sh

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
