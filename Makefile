# Database config from .env
DB_HOST ?= localhost
DB_PORT ?=5432
DB_NAME ?=trucker_marketplace_go_postgres
DB_USER ?=postgres_user
DB_PASS ?=postgres_password
DB_URL = postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

# Path to migration files
MIGRATIONS_PATH = migrations

# -----------------------
# Migration commands
# -----------------------

# Run all up migrations
migrate-up:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

# Rollback last migration
migrate-down:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down 1

# Force version (fix dirty db)
migrate-force:
	@echo "Enter version to force:"
	@read version; \
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" force $$version

# Reset database (down all + up)
migrate-reset:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

.PHONY: migrate-up migrate-down migrate-force migrate-reset
