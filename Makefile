# Makefile for UnIaC - Universal Infrastructure as Code tool

# Variables
BINARY_NAME = uniac
GO_FILES = $(shell find . -name "*.go" -not -path "./vendor/*")
PKG = github.com/arian-press2015/uniac
GO = go

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
	@echo "-------------------"
	@echo "Usage: make <target>"

# Phony targets to avoid conflicts with files
.PHONY: all deps build dev run clean help