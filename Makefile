.PHONY: all prepare build clean dev test backend frontend

VERSION := $(shell git describe --tags --always --dirty)

# Default target builds both
all: build

# Prepare build environment
prepare: prepare-deps prepare-frontend

# Build everything
build: build-frontend build-backend

# Development mode
dev: dev-frontend dev-backend

# Clean build artifacts
clean: clean-frontend clean-backend

# Backend-specific targets
prepare-deps:
	@echo "Installing Go dependencies..."
	go mod download

build-backend: build-frontend
	@echo "Building backend (version: $(VERSION))..."
	go build -ldflags "-X 'github.com/QuantumNous/new-api/common.Version=$(VERSION)'" -o newapi-edu main.go

# Pure backend build (no embedded frontend)
build-backend-pure:
	@echo "Building pure backend (version: $(VERSION))..."
	go build -ldflags "-X 'github.com/QuantumNous/new-api/common.Version=$(VERSION)'" -o newapi-edu-pure main-backend.go

build-backend-pure-linux:
	@echo "Building pure backend for Linux amd64 (version: $(VERSION))..."
	GOOS=linux GOARCH=amd64 go build -ldflags "-X 'github.com/QuantumNous/new-api/common.Version=$(VERSION)'" -o newapi-edu-pure main-backend.go

build-backend-pure-linux-arm64:
	@echo "Building pure backend for Linux arm64 (version: $(VERSION))..."
	GOOS=linux GOARCH=arm64 go build -ldflags "-X 'github.com/QuantumNous/new-api/common.Version=$(VERSION)'" -o newapi-edu-pure-arm64 main-backend.go

build-backend-linux: build-frontend
	@echo "Building backend for Linux amd64 (version: $(VERSION))..."
	GOOS=linux GOARCH=amd64 go build -ldflags "-X 'github.com/QuantumNous/new-api/common.Version=$(VERSION)'" -o newapi-edu main.go

build-backend-linux-arm64: build-frontend
	@echo "Building backend for Linux arm64 (version: $(VERSION))..."
	GOOS=linux GOARCH=arm64 go build -ldflags "-X 'github.com/QuantumNous/new-api/common.Version=$(VERSION)'" -o newapi-edu-arm64 main.go

build-linux: build-backend-linux
	@echo "Linux build complete"

# Frontend-specific targets (using make -C)
prepare-frontend:
	@echo "Preparing frontend dependencies..."
	make -C web/default install

build-frontend:
	@echo "Building frontend..."
	make -C web/default build

dev-frontend:
	@echo "Starting frontend development server..."
	make -C web/default dev

dev-backend: build-frontend
	@echo "Starting backend (frontend must be built first)..."
	./newapi-edu

clean-frontend:
	@echo "Cleaning frontend..."
	make -C web/default clean

clean-backend:
	@echo "Cleaning backend..."
	rm -f newapi-edu newapi-edu-pure

# Test targets
test:
	@echo "Running tests..."
	make -C web/default test

# Additional convenience targets
prepare-all: prepare-deps prepare-frontend
	@echo "All dependencies prepared"

build-prod: build
	@echo "Production build complete"

watch:
	@echo "Starting development with file watching..."
	@echo "Run 'make dev-backend' in another terminal for backend"
	make -C web/default watch