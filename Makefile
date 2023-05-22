# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
BINARY_NAME = fucookie

# Cross-compilation targets
BUILD_FOLDER = build
BUILD_LINUX = linux
BUILD_WINDOWS = windows
BUILD_MAC = mac

# Default target
all: all clean build-linux build-windows build-mac test deps

# Clean the project
clean:
	$(GOCLEAN)
	rm -rf $(BUILD_FOLDER)

# Build the project for Linux
build-linux:
	mkdir -p build
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o build/$(BINARY_NAME)-$(BUILD_LINUX) -ldflags="-s -w" -trimpath -tags prod

# Build the project for Windows
build-windows:
	mkdir -p build
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o build/$(BINARY_NAME)-$(BUILD_WINDOWS).exe -ldflags="-s -w" -trimpath -tags prod

# Build the project for macOS
build-mac:
	mkdir -p build
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o build/$(BINARY_NAME)-$(BUILD_MAC) -ldflags="-s -w" -trimpath -tags prod

# Run tests
test:
	$(GOTEST) -v ./...

# Install dependencies
deps:
	$(GOGET) -v ./...

# Default target
.PHONY: all clean build-linux build-windows build-mac test deps
