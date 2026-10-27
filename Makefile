# Variables
BINARY_NAME=ai-shell
MODELS_DIR=models
VERSION=1.0.0
BUILD_FLAGS=-ldflags="-s -w"

.PHONY: all clean dirs build-windows build-linux build-mac test

all: dirs build-windows build-linux build-mac

dirs:
	@mkdir -p $(MODELS_DIR)
	@echo "[ OK ] Directory structure created. Place your GGUF model in ./$(MODELS_DIR)/"

build-windows:
	@echo "Building for Windows (AMD64)..."
	GOOS=windows GOARCH=amd64 go build $(BUILD_FLAGS) -o $(BINARY_NAME).exe .
	@echo "[ OK ] Windows executable ready: $(BINARY_NAME).exe"

build-linux:
	@echo "Building for Linux (AMD64)..."
	GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS) -o $(BINARY_NAME)-linux .
	@echo "[ OK ] Linux binary ready: $(BINARY_NAME)-linux"

build-mac:
	@echo "Building for macOS (Apple Silicon)..."
	GOOS=darwin GOARCH=arm64 go build $(BUILD_FLAGS) -o $(BINARY_NAME)-mac .
	@echo "[ OK ] macOS binary ready: $(BINARY_NAME)-mac"

test:
	go vet ./...
	go test ./...

clean:
	@rm -f $(BINARY_NAME).exe $(BINARY_NAME)-linux $(BINARY_NAME)-mac
	@echo "[ OK ] Clean complete."