.PHONY: backend
backend:
	@echo "🚗 Starting backend environment..."
	@docker compose up -d api
	@docker compose exec api bash

.PHONY: frontend
frontend:
	@echo "🎨 Starting frontend environment..."
	@docker compose up -d
	@docker compose exec -d api make -C backend run
	@docker compose exec frontend bash

.PHONY: down
down:
	@echo "🛑 Stopping environment..."
	@docker compose down

.PHONY: clean
clean:
	@echo "🧹 Stopping environment and removing volumes..."
	@docker compose down -v --remove-orphans

.PHONY: prune
prune:
	@echo "🗑️  Pruning unused Docker build cache, volumes & images..."
	@docker builder prune -f --keep-storage 2GB
	@docker volume prune -f
	@docker image prune -f
