FILENAME=
DB_USER=postgres
DB_PASSWORD=postgres
DB_HOST=localhost
DB_PORT=5432
DB_NAME=go_skeleton
APP_ENV=development
BINARY=go-skeleton

tasks:
	@echo Usage: make [task]
	@echo -------------------
	@echo
	@cat Makefile

migrate-up:
	bin/migrate -path db/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)_$(APP_ENV)?sslmode=disable" -verbose up

migrate-down:
	bin/migrate -path db/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)_$(APP_ENV)?sslmode=disable" -verbose down

generate-migration:
	bin/migrate create -ext sql -dir db/migrations -format unix $(FILENAME)

create-database:
	docker-compose exec db psql -U $(DB_USER) -c "create database $(DB_NAME)_$(APP_ENV)"
	@echo "Create database $(DB_NAME)_$(APP_ENV) successfully."

drop-database:
	docker-compose exec db psql -U $(DB_USER) -c "drop database $(DB_NAME)_$(APP_ENV)"
	@echo "Drop database $(DB_NAME)_$(APP_ENV) successfully."

setup-test:
	$(MAKE) create-database APP_ENV=test
	$(MAKE) migrate-up APP_ENV=test

remove-setup-test:
	$(MAKE) drop-database APP_ENV=test

sqlc:
	bin/sqlc generate

test:
	go test -v -coverprofile coverage/cover.out  ./...
	go tool cover -html=coverage/cover.out -o coverage/cover.html

server:
	APP_ENV=$(APP_ENV) go run main.go

build:
	@echo "Build binary file..."
	go build -o build/$(BINARY) main.go
	@echo "Copy config..."
	cp config.yaml.example build/config.yaml

mock:
	bin/mockgen -package mockdb  -destination db/mock/store.go github.com/phongtv-1971/go-skeleton/db/sqlc Store

linter:
	bin/golangci-lint run ./...

.PHONY: tasks migrate-up migrate-down generate-migration create-database drop-database setup-test remove-setup-test sqlc test server build mock linter
