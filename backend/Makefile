.PHONY: test test-unit test-cover clean build-dev build-prod run

# Variables
APP_NAME=cochera
MAIN_PATH=main.go

# Development commands (always run)
test:
	go test -v ./...

test-unit:
	go test -v -short ./...

test-cover:
	go test -v -cover ./...

clean:
	rm -f $(APP_NAME)
	go clean -cache
	go clean -testcache

run:
	go run $(MAIN_PATH)

# File-based commands (don't use .PHONY)
build-dev: $(APP_NAME)

$(APP_NAME): $(shell find . -name '*.go')
	go build -o $(APP_NAME) $(MAIN_PATH)

build-prod:
	CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o $(APP_NAME) $(MAIN_PATH)