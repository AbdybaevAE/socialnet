# ------------------------------------------------------
# VARIABLES
# ------------------------------------------------------
SHELL			:= /bin/bash -o pipefail
BUILD_DIR		?= ./bin/
CI_PROJECT_DIR	?= .

# ------------------------------------------------------
# HELP
# ------------------------------------------------------
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# ------------------------------------------------------
# MAIN COMMANDS
# ------------------------------------------------------
.PHONY: all
all: gen format lint test build

.PHONY: build
build: gen ## Build static program binary
	mkdir -es $(BUILD_DIR)
	rm -f $(BUILD_DIR)lambda
	go build -ldflags="-w -s" -o $(BUILD_DIR)lambda lambda/cmd/lambda.go

.PHONY: gen
gen: ## Generate code
	rm -f $(shell find . -type f -name "*mock_test.go" -or -name "*mock.go" -or -name "*generated.go")
	go generate ./...
	go mod tidy -go=1.18 -compat=1.18

.PHONY: format
format: ## Format the source code
	goimports -w .

.PHONY: lint
lint: ## Lint the code
	golangci-lint run

.PHONY: test
test: ## Run tests
	go test -timeout 10m -v -cover -race ./...

# ------------------------------------------------------
# CI COMMANDS
# ------------------------------------------------------
.PHONY: ci-lint
ci-lint: gen
	golangci-lint run --out-format code-climate | tee $(CI_PROJECT_DIR)/gl-code-quality-report.json | jq -r '.[] | "\(.location.path):\(.location.lines.begin) \(.description)"'

.PHONY: ci-test
ci-test: gen
	gotestsum --format standard-verbose --junitfile junit.xml -- -timeout 10m -v -coverpkg=./... -coverprofile=coverage.txt -covermode atomic -race ./...
	mv coverage.txt coverage.txt.orig && grep -Ev '_generated.go|/cmd/' coverage.txt.orig > coverage.txt && rm -f coverage.txt.orig
	go tool cover -func=coverage.txt
	gocover-cobertura < coverage.txt > $(CI_PROJECT_DIR)/gl-code-coverage-report.xml
