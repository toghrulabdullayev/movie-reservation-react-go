project_name = boilerPlate
image_name = gofiber:latest

MIGRATION_DIR = ./migrations
DB_DSN = "postgres://postgres:postgres@localhost:5432/movie_reservation?sslmode=disable"

migration-new: ## Create a new SQL migration file (e.g. make migration-new name=create_users)
	go run github.com/pressly/goose/v3/cmd/goose@latest -dir $(MIGRATION_DIR) create $(name) sql

migration-up: ## Apply migrations locally
	go run github.com/pressly/goose/v3/cmd/goose@latest -dir $(MIGRATION_DIR) postgres $(DB_DSN) up

migration-down: ## Rollback the last migration locally
	go run github.com/pressly/goose/v3/cmd/goose@latest -dir $(MIGRATION_DIR) postgres $(DB_DSN) down

migration-status: ## Check the status of database migrations
	go run github.com/pressly/goose/v3/cmd/goose@latest -dir $(MIGRATION_DIR) postgres $(DB_DSN) status

dev:
	fiber dev -t ./cmd/api/main.go

seed:
	go run ./cmd/seed/main.go

fmt:
	golangci-lint fmt

lint:
	golangci-lint run

test:
	go test ./...

vuln:
	govulncheck ./...

help: ## This help dialog.
	@grep -F -h "##" $(MAKEFILE_LIST) | grep -F -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

run-local: ## Run the app locally
	go run app.go

requirements: ## Generate go.mod & go.sum files
	go mod tidy

clean-packages: ## Clean packages
	go clean -modcache

up: ## Run the project in a local container
	make up-silent
	make shell

build: ## Generate docker image
	docker build -t $(image_name) .

build-no-cache: ## Generate docker image with no cache
	docker build --no-cache -t $(image_name) .

up-silent: ## Run local container in background
	make delete-container-if-exist
	docker run -d -p 3000:3000 --name $(project_name) $(image_name) ./app

up-silent-prefork: ## Run local container in background with prefork
	make delete-container-if-exist
	docker run -d -p 3000:3000 --name $(project_name) $(image_name) ./app -prod

delete-container-if-exist: ## Delete container if it exists
	docker stop $(project_name) || true && docker rm $(project_name) || true

shell: ## Run interactive shell in the container
	docker exec -it $(project_name) /bin/sh

stop: ## Stop the container
	docker stop $(project_name)

start: ## Start the container
	docker start $(project_name)

