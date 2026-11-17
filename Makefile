.PHONY: db-migrate db-migrate-up db-migrate-down

# Default target
.DEFAULT_GOAL := help

help: ## Display this help message
	@echo "Snippetbox Web - Makefile"
	@echo ""
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'

build: ## Build the application
	go build -o bin/web ./cmd/web

clean: ## Remove build artifacts
	rm -rf bin/

db-migrate: db-migrate-up ## Alias for db-migrate-up

# Run all pending migrations
db-migrate-up: ## Run the database migrations
	@echo "Running database migrations..."
	migrate \
		-source file://db/migrations \
		-database 'mysql://root:SuperSecret!@tcp(localhost:3306)/snippetbox' up

# Roll back the most recent migration
db-migrate-down: ## Roll back the most recent migration
	migrate \
		-source file://db/migrations \
		-database 'mysql://root:SuperSecret!@tcp(localhost:3306)/snippetbox' down 1
