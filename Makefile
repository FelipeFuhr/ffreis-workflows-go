.DEFAULT_GOAL := help
SHELL         := /usr/bin/env bash

.PHONY: help fmt fmt-check lint test validate plan secrets-scan-staged lefthook-bootstrap lefthook-install hooks setup

help: ## Show available targets
	@awk 'BEGIN {FS = ":.*##"} /^[a-zA-Z_-]+:.*##/ {printf "\033[36m%-18s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

fmt: ## Format Go example source
	cd examples/hello && gofmt -w .

fmt-check: ## Check Go example formatting (no modifications)
	@files=$$(cd examples/hello && gofmt -l .); \
	if [ -n "$$files" ]; then echo "Files need formatting:\n$$files"; echo "Run 'make fmt' to fix."; exit 1; fi

lint: ## Lint workflow YAML + golangci-lint on examples
	@command -v actionlint >/dev/null 2>&1 && actionlint || \
	  python3 -c "import sys, glob, yaml; [yaml.safe_load(open(f)) for f in glob.glob('.github/workflows/*.yml')]; print('YAML valid')"
	@if command -v golangci-lint >/dev/null 2>&1; then \
	  cd examples/hello && golangci-lint run ./...; \
	else \
	  echo "golangci-lint not found; skipping example lint"; \
	fi

test: ## Run tests in examples/hello
	cd examples/hello && go test ./...

validate: lint ## Alias for lint

plan: ## Not applicable
	@echo "INFO: not applicable for a workflow-only repo."

secrets-scan-staged: ## Scan staged files for secrets
	@command -v gitleaks >/dev/null 2>&1 || { \
		echo "ERROR: gitleaks not found. Install it from https://github.com/gitleaks/gitleaks#installing"; \
		echo "Tip: run 'make setup' after installing to verify your dev environment."; \
		exit 1; \
	}
	gitleaks protect --staged --redact

lefthook-bootstrap: ## Download lefthook binary to .bin/
	LEFTHOOK_VERSION="1.7.10" BIN_DIR=".bin" bash ./scripts/bootstrap_lefthook.sh

lefthook-install: ## Install git hooks via lefthook
	lefthook install

hooks: lefthook-bootstrap lefthook-install ## Bootstrap and install all git hooks

setup: hooks ## Install git hooks and verify required tools
	@command -v gitleaks >/dev/null 2>&1 || { \
		echo ""; \
		echo "ACTION REQUIRED: gitleaks is not installed."; \
		echo "Install it from https://github.com/gitleaks/gitleaks#installing then re-run 'make setup'."; \
		echo ""; \
		exit 1; \
	}
	@echo "Dev environment ready."
