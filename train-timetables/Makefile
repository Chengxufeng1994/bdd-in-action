.PHONY: help test test-verbose test-cucumber-report clean build

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build all packages
	@echo "Building all packages..."
	@go build ./...

test: ## Run acceptance tests with default format
	@echo "Running acceptance tests..."
	@go test ./tests/acceptancetests

test-verbose: ## Run acceptance tests with verbose output
	@echo "Running acceptance tests (verbose)..."
	@go test -v ./tests/acceptancetests -run TestFeatures

test-cucumber-report: ## Generate Cucumber JSON coverage report
	@echo "Generating Cucumber coverage report..."
	@go test -v ./tests/acceptancetests -run TestFeaturesWithCucumberReport
	@echo "Cucumber coverage report generated: tests/acceptancetests/cucumber-report.json"

fmt: ## Format all Go files
	@echo "Formatting Go files..."
	@gofmt -w -s .
	@goimports -w .

vet: ## Run go vet
	@echo "Running go vet..."
	@go vet ./...

lint: fmt vet ## Run formatting and vetting

tidy: ## Tidy up dependencies
	@echo "Tidying up dependencies..."
	@go mod tidy

clean: ## Clean test artifacts
	@echo "Cleaning test artifacts..."
	@rm -f tests/acceptancetests/cucumber-report.json

deps: ## Install/update dependencies
	@echo "Installing dependencies..."
	@go get -u github.com/cucumber/godog@latest
	@go mod tidy

all: lint build test ## Run lint, build, and test
