# -------- KeyForge Makefile (simple & practical) --------
# Binary name and output dir
BINARY_NAME := keyforge
BIN_DIR     := ./binaries

# Colors for test output
GREEN := $(shell printf "\033[32m")
RED   := $(shell printf "\033[31m")
RESET := $(shell printf "\033[0m")

# Always treat these as phony
.PHONY: all build build-binaries test clean run

# Default: build + test + fmt check
all: fmt build test

# Format code (entire repo)
fmt:
	@echo "🎨 Running go fmt..."
	@go fmt ./...
	@echo "✅ Formatting complete."

# Build for current system (assumes main.go at project root)
build: $(BINARY_NAME)

$(BINARY_NAME): $(shell find . -name '*.go')
	@echo "🔧 Building $(BINARY_NAME) for host..."
	@mkdir -p $(BIN_DIR)
	GOOS=$$(go env GOOS) GOARCH=$$(go env GOARCH) go build -o $(BINARY_NAME) .
	@echo "✅ Build complete: $(BINARY_NAME)"

# Cross-build for popular targets
build-binaries:
	@echo "📂 Ensuring $(BIN_DIR) exists..."
	@mkdir -p $(BIN_DIR)
	@echo "🔨 Building cross-platform binaries..."

	# Linux AMD64
	GOOS=linux  GOARCH=amd64 go build -o $(BIN_DIR)/keyforge-linux-amd64 .
	@echo "✅ Built: $(BIN_DIR)/keyforge-linux-amd64"

	# Linux ARM64 (Raspberry Pi 64-bit / Ubuntu ARM)
	GOOS=linux  GOARCH=arm64 go build -o $(BIN_DIR)/keyforge-linux-arm64 .
	@echo "✅ Built: $(BIN_DIR)/keyforge-linux-arm64"

	# Raspberry Pi 32-bit (ARMv7)
	GOOS=linux  GOARCH=arm GOARM=7 go build -o $(BIN_DIR)/keyforge-linux-armv7 .
	@echo "✅ Built: $(BIN_DIR)/keyforge-linux-armv7"

	# macOS Intel
	GOOS=darwin GOARCH=amd64 go build -o $(BIN_DIR)/keyforge-darwin-amd64 .
	@echo "✅ Built: $(BIN_DIR)/keyforge-darwin-amd64"

	# macOS Apple Silicon
	GOOS=darwin GOARCH=arm64 go build -o $(BIN_DIR)/keyforge-darwin-arm64 .
	@echo "✅ Built: $(BIN_DIR)/keyforge-darwin-arm64"

	# Windows 64-bit
	GOOS=windows GOARCH=amd64 go build -o $(BIN_DIR)/keyforge-windows-amd64.exe .
	@echo "✅ Built: $(BIN_DIR)/keyforge-windows-amd64.exe"

	@echo "🎉 All binaries built successfully!"

# Unit tests (plain & reliable)
test:
	@echo "🧪 Running tests..."
	@if go test ./... -v; then \
		echo "$(GREEN)✅ PASS$(RESET)"; \
	else \
		echo "$(RED)❌ FAIL$(RESET)"; \
		exit 1; \
	fi

# Clean artifacts
clean:
	@echo "🧹 Cleaning up build artifacts..."
	@rm -f $(BINARY_NAME)
	@rm -rf $(BIN_DIR)
	@rm -f coverage.out coverage.html
	@echo "✅ Cleaned."

# Build (if needed) then run
run: build
	@echo "🚀 Running $(BINARY_NAME)..."
	./$(BINARY_NAME)
