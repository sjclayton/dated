# Variables
BINARY_NAME=dated
OUTPUT_DIR=out
SRC_DIR=.
GO=go
GOFLAGS=-ldflags="-s -w"

# Default target
all: build

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(OUTPUT_DIR)
	@$(GO) build $(GOFLAGS) -o $(OUTPUT_DIR)/$(BINARY_NAME) $(SRC_DIR)
	@echo "Built successfully in $(OUTPUT_DIR)/"

# Clean the output directory
clean:
	@echo "Cleaning up..."
	@rm -rf $(OUTPUT_DIR)
	@echo "Done cleaning."

# Run the application (with optional args)
run: build
	@echo "Running $(BINARY_NAME) with args: $(filter-out run,$(MAKECMDGOALS))\n"
	@./$(OUTPUT_DIR)/$(BINARY_NAME) $(filter-out run,$(MAKECMDGOALS))

# Run all tests
test:
	@echo "Running tests..."
	@$(GO) test -v ./...

%:
	@:

.PHONY: all build clean run test
