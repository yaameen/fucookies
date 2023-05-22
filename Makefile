# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
BINARY_NAME = fucookie

# Cross-compilation targets
BUILD_FOLDER = build
BUILD_LINUX = $(BUILD_FOLDER)/linux
BUILD_WINDOWS = $(BUILD_FOLDER)/windows
BUILD_MAC = $(BUILD_FOLDER)/mac

# Default target
all: all clean build-linux build-windows build-mac test deps

# Clean the project
clean:
	$(GOCLEAN)
	rm -rf $(BUILD_FOLDER)

# Build the project for Linux
build-linux:
	mkdir -p $(BUILD_LINUX)
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BUILD_LINUX)/$(BINARY_NAME) -ldflags="-s -w" -trimpath -tags prod

# Build the project for Windows
build-windows:
	mkdir -p $(BUILD_WINDOWS)
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BUILD_WINDOWS)/$(BINARY_NAME).exe -ldflags="-s -w" -trimpath -tags prod

# Build the project for macOS
build-mac:
	mkdir -p $(BUILD_MAC)
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BUILD_MAC)/$(BINARY_NAME) -ldflags="-s -w" -trimpath -tags prod

# Run tests
test:
	$(GOTEST) -v ./...

# Install dependencies
deps:
	$(GOGET) -v ./...

# Default target
.PHONY: all clean build-linux build-windows build-mac test deps
