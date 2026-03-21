.DEFAULT_GOAL := help
SHELL         := /usr/bin/env bash

.PHONY: help fmt lint test validate plan

help: ## Show available targets
	@awk 'BEGIN {FS = ":.*##"} /^[a-zA-Z_-]+:.*##/ {printf "\033[36m%-18s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

fmt: ## Not applicable — workflow-only repo
	@echo "INFO: No source code to format in this workflow-only repo."

lint: ## Validate workflow YAML syntax
	@command -v actionlint >/dev/null 2>&1 && actionlint || \
	  python3 -c "import sys, glob, yaml; [yaml.safe_load(open(f)) for f in glob.glob('.github/workflows/*.yml')]; print('YAML valid')"

test: ## Not applicable — workflows are tested by referencing repos
	@echo "INFO: Workflows are validated by the repos that use them."

validate: lint ## Alias for lint

plan: ## Not applicable
	@echo "INFO: not applicable for a workflow-only repo."
