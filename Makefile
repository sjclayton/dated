# Variables
BINARY_NAME=dated
OUTPUT_DIR=out
SRC_DIR=.
GO=go
GOFLAGS=-ldflags="-s -w"

# Default target
all: build

# Build the binary, ensure clean runs first
build: clean
	@mkdir -p $(OUTPUT_DIR)
	$(GO) build $(GOFLAGS) -o $(OUTPUT_DIR)/$(BINARY_NAME) $(SRC_DIR)

# Clean the output directory
clean:
	rm -rf $(OUTPUT_DIR)

# Run the application
run: build
	./$(OUTPUT_DIR)/$(BINARY_NAME)

# Test the project
test:
	$(GO) test ./...

.PHONY: all build clean run test
