.PHONY: help run run-dev lint lint-fix swagger test-integration test-unit test-coverage

help:
	@echo "Targets:"
	@echo "  run               Run the application"
	@echo "  run-dev           Run with live reload (Air)"
	@echo "  lint              Run golangci-lint (installs if missing)"
	@echo "  lint-fix          Run golangci-lint with --fix (installs if missing)"
	@echo "  swagger           Generate Swagger docs (swag init)"
	@echo "  test-integration  Run integration tests (Godog)"
	@echo "  test-unit         Run unit tests"
	@echo "  test-coverage     Run tests and generate coverage.html"

run:
	go run cmd/main.go

run-dev:
	@AIR_BIN="$$(go env GOBIN)"; \
	if [ -z "$$AIR_BIN" ]; then AIR_BIN="$$(go env GOPATH)/bin"; fi; \
	if [ ! -x "$$AIR_BIN/air" ]; then \
		echo "air not found. Installing..."; \
		go install github.com/air-verse/air@latest; \
	fi; \
	"$$AIR_BIN/air"

lint:
	@LINT_BIN="$$(go env GOBIN)"; \
	if [ -z "$$LINT_BIN" ]; then LINT_BIN="$$(go env GOPATH)/bin"; fi; \
	if [ ! -x "$$LINT_BIN/golangci-lint" ]; then \
		echo "golangci-lint not found. Installing..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi; \
	"$$LINT_BIN/golangci-lint" run

lint-fix:
	@LINT_BIN="$$(go env GOBIN)"; \
	if [ -z "$$LINT_BIN" ]; then LINT_BIN="$$(go env GOPATH)/bin"; fi; \
	if [ ! -x "$$LINT_BIN/golangci-lint" ]; then \
		echo "golangci-lint not found. Installing..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi; \
	"$$LINT_BIN/golangci-lint" run --fix

swagger:
	@SWAG_BIN="$$(go env GOBIN)"; \
	if [ -z "$$SWAG_BIN" ]; then SWAG_BIN="$$(go env GOPATH)/bin"; fi; \
	if [ ! -x "$$SWAG_BIN/swag" ]; then \
		echo "swag not found. Installing..."; \
		go install github.com/swaggo/swag/cmd/swag@latest; \
	fi; \
	"$$SWAG_BIN/swag" init -g cmd/main.go -o docs

test-integration:
	go test -count=1 -v ./tests

test-unit:
	go test -v ./...

test-coverage:
	@PKGS=$$(go list ./... | grep -v '/cmd' | grep -v '/docs' | grep -v '/tests'); \
	go test -coverprofile=coverage.out $$PKGS
	go tool cover -html=coverage.out -o coverage.html
