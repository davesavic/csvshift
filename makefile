# Makefile

replace-antlr4:
	@echo "Replacing imports in gen files..."
	@find gen/ -type f -exec sed -i 's|github.com/antlr/antlr4/runtime/Go/antlr/v4|github.com/antlr4-go/antlr|g' {} +
	@echo "Replacement done."

.PHONY: replace-antlr4

OS := $(shell uname)

build:
	@echo "Building..."
ifeq ($(OS),Darwin)
	@go build -o bin/csvshift ./...
endif
ifeq ($(OS),Linux)
	@go build -o bin/csvshift ./...
endif
ifeq ($(OS),Windows_NT)
	@go build -o bin/csvshift.exe ./...
endif
	@echo "Build complete."

.PHONY: build

test:
	@echo "Testing..."
	@go test -v ./...
	@echo "Test complete."

.PHONY: test