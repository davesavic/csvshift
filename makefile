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

# Get the latest git tag
CURRENT_VERSION=$(shell git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")

# Extract major, minor, and patch numbers
MAJOR=$(shell echo $(CURRENT_VERSION) | cut -d. -f1 | tr -d v)
MINOR=$(shell echo $(CURRENT_VERSION) | cut -d. -f2)
PATCH=$(shell echo $(CURRENT_VERSION) | cut -d. -f3)

current-version:
	@echo $(CURRENT_VERSION)

.PHONY: current-version

release:
	@if [ "$(type)" = "major" ]; then \
		NEW_MAJOR=$$(( $(MAJOR) + 1 )); \
		NEW_VERSION="v$$NEW_MAJOR.0.0"; \
	elif [ "$(type)" = "minor" ]; then \
		NEW_MINOR=$$(( $(MINOR) + 1 )); \
		NEW_VERSION="v$(MAJOR).$$NEW_MINOR.0"; \
	elif [ "$(type)" = "patch" ]; then \
		NEW_PATCH=$$(( $(PATCH) + 1 )); \
		NEW_VERSION="v$(MAJOR).$(MINOR).$$NEW_PATCH"; \
	else \
		echo "Invalid release type. Use major, minor, or patch."; \
		exit 1; \
	fi; \
	trap 'echo "Error encountered, cleaning up tags..."; git tag -d $$NEW_VERSION; git push origin :refs/tags/$$NEW_VERSION;' ERR; \
	git tag -a $$NEW_VERSION -m "$(message)"; \
	git push origin $$NEW_VERSION; \
	goreleaser --clean; \
	trap - ERR;

.PHONY: release
