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

.PHONY: release
release:
	@if [ -z "$(VERSION)" ]; then \
		echo "❌ Falta VERSION. Uso: make release VERSION=X.Y.Z"; \
		exit 1; \
	fi
	@echo "$(VERSION)" | grep -Eq '^[0-9]+\.[0-9]+\.[0-9]+$$' || { \
		echo "❌ VERSION debe tener el formato X.Y.Z (recibido: $(VERSION))"; \
		exit 1; \
	}
	@[ -z "$$(git status --porcelain)" ] || { \
		echo "❌ El working tree tiene cambios sin commitear"; \
		exit 1; \
	}
	@[ "$$(git rev-parse --abbrev-ref HEAD)" = "main" ] || { \
		echo "❌ Los releases se hacen desde main (rama actual: $$(git rev-parse --abbrev-ref HEAD))"; \
		exit 1; \
	}
	@git rev-parse "v$(VERSION)" >/dev/null 2>&1 && { \
		echo "❌ El tag v$(VERSION) ya existe"; \
		exit 1; \
	} || true
	@echo "🔖 Actualizando versión a $(VERSION)..."
	@sed -i 's/return "[0-9]*\.[0-9]*\.[0-9]*"/return "$(VERSION)"/' backend/application/version.go
	@sed -i 's/"version": "[0-9]*\.[0-9]*\.[0-9]*"/"version": "$(VERSION)"/' frontend/package.json
	@git add backend/application/version.go frontend/package.json
	@git commit -m "chore(release): v$(VERSION)"
	@git tag -a "v$(VERSION)" -m "v$(VERSION)"
	@echo "✅ Release v$(VERSION) preparado. Para publicarlo:"
	@echo "   git push origin main --follow-tags"
