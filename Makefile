# Load .env configurations
ifneq (,$(wildcard ./.env))
    include .env
    export
    MIGRATIONS_PATH = db/migrations
	DATABASE_URL = postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)
endif

# run rest http
run:
	go run cmd/main.go

# test is using for unit_test
test:
	go test ./... -covermode=count -coverprofile coverage.out

# will install all tools needed in this repository
tools:
	bash ./install_tools

# migrate-create will create migration file. this file will include for `drop table` and `create table`
migrate-create:
	migrate create -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

# migrate-up will create all table inside migration_path
migrate-up:
	migrate --path $(MIGRATIONS_PATH) --database "$(DATABASE_URL)" up

# migrate-down will drop all table
migrate-down:
	migrate --path $(MIGRATIONS_PATH) --database "$(DATABASE_URL)" down --all
