# Makefile for UnIaC - Unified Infrastructure as Code tool

# Variables
BINARY_NAME = uniac
GO_FILES = $(shell find . -name "*.go" -not -path "./vendor/*")
PKG = github.com/arian-press2015/uniac
GO = go
VERSION ?= $(shell git describe --tags --abbrev=0 2>/dev/null | sed 's/^v//' || echo "0.0.0")
NEXT_VERSION = $(shell \
  major=$$(echo "$(VERSION)" | awk -F'.' '{print $$1}'); \
  minor=$$(echo "$(VERSION)" | awk -F'.' '{print $$2}'); \
  patch=$$(echo "$(VERSION)" | awk -F'.' '{print $$3}'); \
  if [ "$(INCREMENT)" = "MAJOR" ]; then \
    printf "%d.0.0" $$((major + 1)); \
  else if [ "$(INCREMENT)" = "MINOR" ]; then \
    printf "%d.%d.0" "$$major" $$((minor + 1)); \
  else if [ "$(INCREMENT)" = "PATCH" ]; then \
    printf "%d.%d.%d" "$$major" "$$minor" $$((patch + 1)); \
  else \
    printf "%s" "$(VERSION)"; \
  fi; fi; fi)

# Default target
all: help

# Install dependencies
deps:
	@echo "Installing dependencies..."
	$(GO) mod tidy
	$(GO) get -u ./...

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	$(GO) build -o $(BINARY_NAME) ./cmd/uniac

# Development execution
dev:
	@echo "Running $(BINARY_NAME) in development mode..."
	$(GO) run ./cmd/uniac/main.go

# Run the binary
run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BINARY_NAME) infra.yaml

# Clean up generated files
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)
	$(GO) clean

# Release the project
release:
	@echo "Calculating next version based on $(VERSION)..."
	@if [ -z "$(INCREMENT)" ] || [ "$(INCREMENT)" != "MAJOR" ] && [ "$(INCREMENT)" != "MINOR" ] && [ "$(INCREMENT)" != "PATCH" ]; then \
		echo "Error: INCREMENT must be set to MAJOR, MINOR, or PATCH (e.g., make release INCREMENT=MINOR)"; \
		exit 1; \
	fi
	@echo "New version: v$(NEXT_VERSION)"
	@echo "Tagging version v$(NEXT_VERSION)..."
	git tag v$(NEXT_VERSION)
	git push origin v$(NEXT_VERSION)
	@echo "Release triggered. Check GitHub Actions for progress."

# Help target to display available commands
help:
	@echo "UnIaC Makefile Help"
	@echo "-------------------"
	@echo "Available targets:"
	@echo "  all       - Show this help (default)"
	@echo "  deps      - Install or update dependencies"
	@echo "  build     - Build the UnIaC binary"
	@echo "  dev       - Run in development mode"
	@echo "  run       - Build and run the UnIaC binary with infra.yaml"
	@echo "  clean     - Remove binary and clean up"
	@echo "  release   - Create a Git tag and trigger a release (set INCREMENT, e.g., make release INCREMENT=MINOR)"	
	@echo "-------------------"
	@echo "Usage: make <target>"

# Phony targets to avoid conflicts with files
.PHONY: all deps build dev run clean release help