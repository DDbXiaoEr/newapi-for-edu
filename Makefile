.PHONY: all prepare build clean dev test backend frontend docker-backend docker-backend-arm64 docker-allinone docker-allinone-arm64

VERSION := $(shell git describe --tags --always --dirty)
BUILD_DIR := build
IMAGE ?= newapi-edu-pure
ALLINONE_IMAGE ?= newapi-edu
IMAGE_TAG ?= latest

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
	@mkdir -p $(BUILD_DIR)
	@echo "Building backend (version: $(VERSION))..."
	go build -ldflags "-X 'github.com/QuantumNous/new-api/common.Version=$(VERSION)'" -o $(BUILD_DIR)/newapi-edu main.go

# Pure backend build (no embedded frontend)
build-backend-pure:
	@mkdir -p $(BUILD_DIR)
	@echo "Building pure backend (version: $(VERSION))..."
	go build -ldflags "-X 'github.com/QuantumNous/new-api/common.Version=$(VERSION)'" -o $(BUILD_DIR)/newapi-edu-pure main-backend.go

build-backend-pure-linux:
	@mkdir -p $(BUILD_DIR)
	@echo "Building pure backend for Linux amd64 (version: $(VERSION))..."
	GOOS=linux GOARCH=amd64 go build -ldflags "-X 'github.com/QuantumNous/new-api/common.Version=$(VERSION)'" -o $(BUILD_DIR)/newapi-edu-pure main-backend.go

build-backend-pure-linux-arm64:
	@mkdir -p $(BUILD_DIR)
	@echo "Building pure backend for Linux arm64 (version: $(VERSION))..."
	GOOS=linux GOARCH=arm64 go build -ldflags "-X 'github.com/QuantumNous/new-api/common.Version=$(VERSION)'" -o $(BUILD_DIR)/newapi-edu-pure-arm64 main-backend.go

build-backend-linux: build-frontend
	@mkdir -p $(BUILD_DIR)
	@echo "Building backend for Linux amd64 (version: $(VERSION))..."
	GOOS=linux GOARCH=amd64 go build -ldflags "-X 'github.com/QuantumNous/new-api/common.Version=$(VERSION)'" -o $(BUILD_DIR)/newapi-edu main.go

build-backend-linux-arm64: build-frontend
	@mkdir -p $(BUILD_DIR)
	@echo "Building backend for Linux arm64 (version: $(VERSION))..."
	GOOS=linux GOARCH=arm64 go build -ldflags "-X 'github.com/QuantumNous/new-api/common.Version=$(VERSION)'" -o $(BUILD_DIR)/newapi-edu-arm64 main.go

build-linux: build-backend-linux
	@echo "Linux build complete"

# Pure-backend (API only) Docker image based on debian:bullseye-v1.4.1.
# The binary is cross-compiled for linux/amd64; build context is the repo root.
docker-backend: build-backend-pure-linux
	@echo "Building pure-backend image $(IMAGE):$(IMAGE_TAG)..."
	docker build -f docker/Dockerfile.backend -t $(IMAGE):$(IMAGE_TAG) .

docker-backend-arm64: build-backend-pure-linux-arm64
	@echo "Building pure-backend image $(IMAGE):$(IMAGE_TAG) (arm64)..."
	docker buildx build --platform linux/arm64 --build-arg PURE_BINARY=build/newapi-edu-pure-arm64 -f docker/Dockerfile.backend -t $(IMAGE):$(IMAGE_TAG) .

# All-in-one Docker image with the frontend embedded (built from main.go).
docker-allinone: build-backend-linux
	@echo "Building all-in-one image $(ALLINONE_IMAGE):$(IMAGE_TAG)..."
	docker build -f docker/Dockerfile.allinone -t $(ALLINONE_IMAGE):$(IMAGE_TAG) .

docker-allinone-arm64: build-backend-linux-arm64
	@echo "Building all-in-one image $(ALLINONE_IMAGE):$(IMAGE_TAG) (arm64)..."
	docker buildx build --platform linux/arm64 --build-arg APP_BINARY=build/newapi-edu-arm64 -f docker/Dockerfile.allinone -t $(ALLINONE_IMAGE):$(IMAGE_TAG) .

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
	./$(BUILD_DIR)/newapi-edu

clean-frontend:
	@echo "Cleaning frontend..."
	make -C web/default clean

clean-backend:
	@echo "Cleaning backend binaries..."
	rm -f $(BUILD_DIR)/newapi-edu $(BUILD_DIR)/newapi-edu-pure $(BUILD_DIR)/newapi-edu-arm64 $(BUILD_DIR)/newapi-edu-pure-arm64

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