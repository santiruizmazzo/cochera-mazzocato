.PHONY: test test-unit test-integration test-e2e test-cover clean lint build-dev build-prod run

# Variables
APP_NAME=cochera
MAIN_PATH=./cmd/server/main.go

# Development commands (always run)
test:
	go test ./...

test-unit:
	go test -short ./...

test-integration:
	go test -run Integration ./...

test-e2e:
	go test -run EndToEnd ./...

test-cover:
	go test -v -cover ./...

# Ejecutar todas las migraciones pendientes
migrate-up:
	migrate -path db/migrations -database "$(DB_URL)" up

# Rollback una migración
migrate-down:
	migrate -path db/migrations -database "$(DB_URL)" down 1

# Forzar una versión específica (cuidado!)
migrate-force:
	@read -p "Enter version to force: " version; \
	migrate -path db/migrations -database "$(DB_URL)" force $$version

# Ver versión actual
migrate-version:
	migrate -path db/migrations -database "$(DB_URL)" version

# Crear nueva migración
migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir db/migrations -seq $$name

clean:
	rm -f $(APP_NAME)
	go clean -cache
	go clean -testcache

lint:
	golangci-lint run

run:
	go run $(MAIN_PATH)

# File-based commands (don't use .PHONY)
build-dev: $(APP_NAME)

$(APP_NAME): $(shell find . -name '*.go')
	go build -o $(APP_NAME) $(MAIN_PATH)

build-prod:
	CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o $(APP_NAME) $(MAIN_PATH)