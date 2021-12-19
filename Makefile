.PHONY: server migrate-up migrate-down generate-migration create-database drop-database setup-test remove-setup-test sqlc test build mock linter help

include config.mk

server: ## Start server
	APP_ENV=$(APP_ENV) go run main.go

migrate-up: ## Migration up
	bin/migrate -path db/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)_$(APP_ENV)?sslmode=disable" -verbose up

migrate-down: ## Migration down
	bin/migrate -path db/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)_$(APP_ENV)?sslmode=disable" -verbose down

generate-migration: ## Generate migration file
	bin/migrate create -ext sql -dir db/migrations -format unix $(name)

create-database: ## Create database with docker db instance
	docker-compose exec db psql -U $(DB_USER) -c "create database $(DB_NAME)_$(APP_ENV)"
	@echo "Create database $(DB_NAME)_$(APP_ENV) successfully."

drop-database: ## Drop database with docker db instance
	docker-compose exec db psql -U $(DB_USER) -c "drop database $(DB_NAME)_$(APP_ENV)"
	@echo "Drop database $(DB_NAME)_$(APP_ENV) successfully."

setup-test: ## Setup environment for test
	$(MAKE) create-database APP_ENV=test
	$(MAKE) migrate-up APP_ENV=test

remove-setup-test: ## Remove setup environment for test
	$(MAKE) drop-database APP_ENV=test

sqlc: ## Generate sqlc code
	bin/sqlc generate

test: ## Execute test
	go test -v -coverprofile coverage/cover.out  ./...
	go tool cover -html=coverage/cover.out -o coverage/cover.html

build: ## Build binary file
	@echo "Build binary file..."
	go build -o build/$(BINARY) main.go
	@echo "Copy config..."
	cp config.yaml.example build/config.yaml

mock: ## Generate mock code
	bin/mockgen -package mockdb  -destination db/mock/store.go github.com/phongtv-1971/go-skeleton/db/sqlc Store

linter: ## Check linter
	bin/golangci-lint run ./...

help:
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' ./Makefile | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo
	@echo You can modify the default settings for this Makefile creating a file config.mk
	@echo
