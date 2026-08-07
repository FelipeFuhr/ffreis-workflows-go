# Agent Context

**This repo:** `ffreis-workflows-go` — reusable GitHub Actions workflow library for Go
projects. Covers fmt, lint, test, matrix build, unit + integration coverage, SBOM,
container build, fuzzing, mutation testing, and OSV scanning.

## Non-obvious rules (read before changing anything)

1. **ALL `go-*.yml` workflows must appear in `ci.yml`.** No exceptions —
   unlike container/python repos, there is no live-infra exemption here.

2. **Shell injection prevention is enforced by Semgrep** (`run-shell-injection` rule).
   Some action SHAs trigger false-positive secret detection — suppress with
   `# nosemgrep: <rule-id>` inline comment, not by disabling the rule.

3. **Third-party action SHAs are managed by Renovate.** Do not edit manually.

4. **Fork PR gating for secrets** (e.g., Codecov token):
   ```yaml
   if: github.event_name != 'pull_request' || github.event.pull_request.head.repo.fork == false
   ```

5. **Concurrency is caller-controlled.** Never add `concurrency:` to reusable workflows.

6. **`examples/hello/` tests both Go and container workflows** — it has Go source, a
   Containerfile, and CI config. Don't simplify or split it.

7. **Private cross-repo Go dependency access (`goprivate` input).**
   `go-mod-tidy-check.yml`, `go-lint.yml`, and `go-test.yml` accept an optional
   `goprivate` string input plus an optional `GIT_AUTH_TOKEN` secret. When
   `goprivate` is non-empty, a step sets `GOPRIVATE`/`GONOSUMCHECK` via
   `$GITHUB_ENV` and (if `GIT_AUTH_TOKEN` is set) configures
   `git config --global url."https://x-access-token:${GIT_AUTH_TOKEN}@github.com/".insteadOf
   "https://github.com/"` so `go mod`/`go vet`/`go test`/golangci-lint can
   resolve private module paths. Both default to empty/unset — zero behavior
   change for existing callers. `go-fmt.yml` is intentionally NOT wired: `gofmt`
   never resolves modules, so it has nothing to authenticate. Secret name uses
   `GIT_AUTH_TOKEN` (SCREAMING_SNAKE_CASE), matching the fleet's existing
   `secrets:` naming convention (e.g. `CI_REPO_READ_TOKEN` in
   `ffreis-workflows-general`'s `general-kb-sync.yml`) — GitHub Actions secret
   names cannot contain hyphens.

8. **`go-integration-coverage.yml` is a separate gated metric from
   `go-coverage.yml`**, not a variant of it — same input/step shape
   (`coverage-threshold`, opt-in gate at 0), but its own Codecov flag
   (`integration` vs `unit`) and its own coverage file
   (`coverage-integration.out` vs `coverage.out`) so the two never collide on
   upload. It only runs `go test -tags=<build-tag>` (default tag:
   `integration`) when at least one file in `working-directory` actually
   declares that build tag (`//go:build integration` / `// +build
   integration`) — checked by grep before any test runs. This guard exists
   because untagged test files always run regardless of `-tags`, so a repo
   with zero integration-tagged files would otherwise silently re-report its
   unit coverage as "integration coverage" (a false pass/fail signal) instead
   of skipping the gate. `examples/hello/calculator/calculator_integration_test.go`
   is the fleet's reference example of a tagged file.

## Structure

```
.github/workflows/
  go-*.yml        ← reusable library
  devops-*.yml    ← repo-maintenance (exempt from self-test)
  ci.yml          ← self-test orchestrator
examples/hello/   ← Go project + Containerfile + .golangci.yml
Makefile          ← setup, fmt, fmt-check, lint, test, hooks
```

## Build/test

```bash
make setup              # lefthook + gitleaks check
make fmt                # gofmt examples/hello
make lint               # actionlint + golangci-lint (soft-fail if not installed)
make test               # go test ./... in examples/hello
```

## Cross-repo role

Consumed by Go repos in the fleet. Callers must pass `working-directory` when their
Go code is not at the repo root.

## Public repo — private-repo hygiene

This is a **public** GitHub repository. When writing commit messages, PR titles,
PR descriptions, or any other user-visible text, **never name private repos** —
website content, inventory, infra, Lambda, or data repos that are not publicly
listed. Use generic terms instead: "the fleet inventory", "a private consumer",
"internal infra", "private data repo", etc.

## Keeping this file current

- **If you discover a fact not reflected here:** add it before finishing your task.
- **If something here is wrong or outdated:** correct it in the same commit as the code change.
- **If you rename a file, command, or concept referenced here:** update the reference.
