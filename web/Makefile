.PHONY: install build clean dev test

install:
	@echo "Installing dependencies with bun..."
	bun install

build:
	@echo "Building frontend with bun..."
	bun run build

clean:
	@echo "Cleaning build artifacts..."
	rm -rf dist

dev:
	@echo "Starting development server..."
	bun run dev

test:
	@echo "Running tests..."
	bun run test

install-all: install
	@echo "Installing all dependencies including dev dependencies..."
	bun install --frozen-lockfile

build-prod: build
	@echo "Production build complete"

watch:
	@echo "Starting development server with file watcher..."
	bun run dev --watch