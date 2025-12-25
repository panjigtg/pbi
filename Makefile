include .env
export

MIGRATE_DIR=./internal/config/migrations

.PHONY: migrate-create migrate-up migrate-down migrate-force migrate-version

run: 
	go run ./app/web/main.go

migrate-create:
	migrate create -ext sql -dir $(MIGRATE_DIR) -seq $(name)

migrate-up:
	migrate -path $(MIGRATE_DIR) -database "$(DB_SOURCE)" up

migrate-down:
	migrate -path $(MIGRATE_DIR) -database "$(DB_SOURCE)" down 1

migrate-force:
	migrate -path $(MIGRATE_DIR) -database "$(DB_SOURCE)" force $(version)

migrate-version:
	migrate -path $(MIGRATE_DIR) -database "$(DB_SOURCE)" version