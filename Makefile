# Variables
NAME := TraceeClient
BUILD_DIR := ./dist
SRC := ./...

# Default target: create build dir, build, and install
.PHONY: all
all: $(BUILD_DIR) build install

# Build target
.PHONY: build
build: ## Build the Go binary
	@echo "Building $(NAME)..."
	go build -o $(BUILD_DIR) $(SRC)

# Install target
.PHONY: install
install: ## Build and install the binary
	@echo "Installing $(NAME)..."
	go install $(SRC)

# Directory creation
$(BUILD_DIR):
	@echo "Creating build directory..."
	mkdir -p $(BUILD_DIR)

# Clean target
.PHONY: clean
clean: ## Remove binary and build directory
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)

# Environment target
.PHONY: env
env: ## Display Go environment variables
	@echo "Go environment variables:"
	go env

# Help target
.PHONY: help
help: ## Show this help message
	@echo "Available make targets:"
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} \
    /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)
