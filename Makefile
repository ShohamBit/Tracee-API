# Variables
NAME := TraceeClient
BUILD_DIR := ./dist
SRC :=  $(shell find . -name '*.go') # Finds all Go source files
BINARY := $(BUILD_DIR)/$(NAME)

# Default target: create build dir, build, and install
.PHONY: all
all: $(BUILD_DIR) $(BINARY) install

# Build binary only if source files have changed
$(BINARY): $(SRC) | $(BUILD_DIR)
	@echo "Building $(NAME)..."
	go build  ./...

# Install target
.PHONY: install
install: $(BINARY) ## Build and install the binary
	@echo "Installing $(NAME)..."
	go install ./...

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

# Test target
.PHONY: test
test: ## Test TraceeClient
	go test ./... -v
