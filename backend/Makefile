APP_NAME=cochera
MAIN_PATH=./cmd/server/main.go

GOLANGCI_LINT_VERSION = v2.4.0
GOBIN := $(shell go env GOPATH)/bin

DATABASE_URL=$(DEV_DB_URL)
MIGRATIONS_PATH=./migrations
MIGRATE_VERSION := v4.18.1
MIGRATE_URL := https://github.com/golang-migrate/migrate/releases/download/$(MIGRATE_VERSION)/migrate.linux-amd64.tar.gz

.PHONY: test
test: migrate-test-up test-unit test-integration test-e2e

.PHONY: test-unit
test-unit:
	@echo "🦠 Running unit tests..."
	@go test -short ./tests/unit/...

.PHONY: test-integration
test-integration: migrate-test-up
	@echo "🔌 Running integration tests..."
	@go test -run Integration ./tests/integration/...

.PHONY: test-e2e
test-e2e: migrate-test-up
	@echo "↔️  Running end to end tests..."
	@go test -run EndToEnd ./tests/end_to_end/...

.PHONY: test-cover
test-cover:
	@echo "☂️  Checking app's code coverage..."
	@go test -v -cover ./...

.PHONY: migrate-test-up
migrate-test-up: DATABASE_URL = $(TEST_DB_URL)
migrate-test-up: migrate-up

.PHONY: migrate-prod-up
migrate-prod-up: DATABASE_URL = $(PROD_DB_URL)
migrate-prod-up:
	@echo "🤔 Checking current migration version..."
	@./migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" version || true
	@echo "⏭️  Running migrations..."
	@./migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" up

.PHONY: migrate-up
migrate-up:
	@echo "🤔 Checking current migration version..."
	@migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" version || true
	@echo "⏭️  Running migrations..."
	@migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" up

.PHONY: migrate-down
migrate-down:
	@echo "⏮️ Rolling back migrations to previous version..."
	@migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" down 1

.PHONY: migrate-test-reset
migrate-test-reset: DATABASE_URL = $(TEST_DB_URL)
migrate-test-reset: migrate-reset

.PHONY: migrate-reset
migrate-reset:
	@echo "🔁 Resetting database to original state..."
	@migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" drop -f
	@migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" up

# Forzar una versión específica (cuidado!)
.PHONY: migrate-force
migrate-force:
	@read -p "Enter version to force: " version; \
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" force $$version

.PHONY: migrate-test-version
migrate-test-version: DATABASE_URL = $(TEST_DB_URL)
migrate-test-version: migrate-version

.PHONY: migrate-version
migrate-version:
	@echo "🤔 Checking current migration version..."
	@migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" version

.PHONY: migrate-create
migrate-create:
	@echo "✍️  Creating new database migration..."
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $$name

.PHONY: install-migrate-prod
install-migrate-prod:
	curl -L $(MIGRATE_URL) | tar xvz
	chmod +x migrate

.PHONY: clean
clean:
	@echo "🧹 Cleaning app's executables & cache..."
	@rm -f $(APP_NAME)
	@go clean -cache
	@go clean -testcache

.PHONY: lint-install
lint-install:
	@echo "⚙️  Installing golangci-lint $(GOLANGCI_LINT_VERSION)..."
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
	  | sh -s -- -b $(GOBIN) $(GOLANGCI_LINT_VERSION)
	@echo "✅ golangci-lint installed at $(GOBIN)/golangci-lint"
	@$(GOBIN)/golangci-lint version

.PHONY: lint
lint:
	@echo "🔍 Linting project's code..."
	@golangci-lint run

.PHONY: run
run: migrate-up
	@echo "🚗 Starting Cochera Mazzocato API..."
	@go run $(MAIN_PATH)

.PHONY: run-prod
run-prod: install-migrate-prod migrate-prod-up
	@echo "🚗 Starting Cochera Mazzocato API..."
	@./$(APP_NAME)

# File-based commands (don't use .PHONY)
.PHONY: build-dev
build-dev: $(APP_NAME)

$(APP_NAME): $(shell find . -name '*.go')
	@echo "📦 Building app for development environment..."
	@go build -o $(APP_NAME) $(MAIN_PATH)

.PHONY: build-prod
build-prod:
	@echo "🎁 Building app for production environment..."
	@GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o $(APP_NAME) $(MAIN_PATH)