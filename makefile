# Makefile

# The directory containing the files you want to modify
TARGET_DIR := gen

# The 'all' target will be the default target if 'make' is invoked without arguments
all: replace

replace:
	@echo "Replacing text in files..."
	@find $(TARGET_DIR) -type f -exec sed -i '' 's|github.com/antlr/antlr4/runtime/Go/antlr/v4|github.com/antlr4-go/antlr|g' {} +
	@echo "Replacement done."

.PHONY: all replace
