.PHONY: all
# Default target
all: build install

# Variables
BINARY_NAME=client
SRC_DIR=./...
BUILD_DIR=dist

# Create the build directory if it doesn't exist
$(BUILD_DIR):
	mkdir -p $@

# Build the binary
build: $(BUILD_DIR)  # Ensure BUILD_DIR exists before building
	@echo "Building the binary..."
	go build -o $(BUILD_DIR) $(SRC_DIR)

# Install the binary
install:
	@echo "Installing the binary..."
	go install $(SRC_DIR)

# Clean the build
clean:
	@echo "Cleaning up..."
	rm -f $(BUILD_DIR)/$(BINARY_NAME)

# Help command
help:
	@echo "Makefile commands:"
	@echo "  make build    Build the binary"
	@echo "  make install  Install the binary"
	@echo "  make clean    Remove the built binary"
	@echo "  make help     Show this help message"


