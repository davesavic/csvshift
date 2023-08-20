# Makefile

replace-antlr4:
	@echo "Replacing imports in gen files..."
	@find gen/ -type f -exec sed -i 's|github.com/antlr/antlr4/runtime/Go/antlr/v4|github.com/antlr4-go/antlr|g' {} +
	@echo "Replacement done."

.PHONY: replace-antlr4

test:
	@echo "Testing..."
	@go test -v ./...
	@echo "Test complete."

.PHONY: test