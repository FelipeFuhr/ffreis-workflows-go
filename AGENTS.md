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
   `go-mod-tidy-check.yml`, `go-lint.yml`, `go-test.yml`, `go-sonar.yml`, and
   `go-cross-build-matrix.yml` accept an optional `goprivate` string input plus
   an optional `GIT_AUTH_TOKEN` secret. When
   `goprivate` is non-empty, a step sets `GOPRIVATE`/`GONOSUMCHECK` via
   `$GITHUB_ENV` and (if `GIT_AUTH_TOKEN` is set) configures
   `git config --global url."https://x-access-token:${GIT_AUTH_TOKEN}@github.com/".insteadOf
   "https://github.com/"` so `go mod`/`go vet`/`go test`/`go build`/golangci-lint/
   the SonarCloud scanner's own `go test` coverage step can resolve private
   module paths. Both default to empty/unset — zero behavior
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

   Its `dynamodb-local` boolean input (default `false`, additive) starts a
   real `amazon/dynamodb-local:3.3.1` container via explicit `docker run -d`/
   `docker stop` steps — not a `services:` block, which GitHub Actions can't
   gate on an input — waits for it with a real curl-based readiness loop
   (any completed HTTP transaction on :8000, not a fixed sleep), then proves
   the DynamoDB API itself responds via `aws dynamodb list-tables
   --endpoint-url` (stronger than "port is open") before exporting
   `DYNAMODB_ENDPOINT=http://localhost:8000` for the test step (matching the
   env var + default every current consumer's own `ddbEndpoint()`-style
   helper already reads). Without this,
   an integration test written to "skip when nothing is reachable" (a
   deliberate fleet convention so `go test ./...` stays green without a
   container runtime) silently skips forever in CI with the job still
   reporting green — confirmed for real in a caller's own log, not just by
   reading the YAML. **Requires a runner with Docker.** The self-hosted
   `local` default does not have a container runtime provisioned (same
   constraint `go-container.yml` already documents for its own `runner`
   default) — any caller setting `dynamodb-local: true` must also pass
   `runner: '["ubuntu-latest"]'`.

9. **`go-mutation.yml`'s `packages` input must be plain directories, never
   `...`-suffixed.** gremlins' `unleash [path]` takes exactly one plain
   directory path — not a go-list `...` pattern, and not several
   space-separated paths in one invocation (more than one array element in
   a single call fails loudly with "accepts at most 1 arg(s)"). A single
   `...`-suffixed path (the default, `./internal`, and the shape most
   callers' READMEs still show, e.g. `./internal/...`) used to fail
   silently instead: gremlins doesn't understand the glob, prints "No
   results to report." and still exits 0 — a vacuous pass, zero mutants
   tested, gate still green. The "Run mutation testing" step now loops
   over each space-separated entry in `packages`, stripping a trailing
   `/...` before invoking gremlins once per package, and propagates the
   worst exit code plus the minimum efficacy score across the run. Mirrors
   the identical fix already applied in `ffreis-platform-configctl`'s own
   Makefile `mutation:` target — verify any change here against a real
   package with real mutants (`examples/hello/calculator` has some), not
   just a clean exit code, since exit 0 is exactly what the bug also
   produced.

10. **`govulncheck-go-version` is decoupled from `go-version` on purpose.**
    `golang/govulncheck-action` always runs `go install
    golang.org/x/vuln/cmd/govulncheck@latest` — it has no version-pin input of
    its own. When x/vuln raises its own Go floor (v1.8.0 needs Go >= 1.26),
    every caller still on `go-version: "1.25.x"` fails at the install step,
    before a single package is scanned. `go-security.yml` and `go-test.yml`
    (its embedded govulncheck step) both take a separate
    `govulncheck-go-version` input, defaulted to a version that satisfies the
    current x/vuln floor, so a future floor bump is fixed by bumping this one
    default rather than every consumer's `go-version`. The scan itself still
    honours the target module's own `go` directive in `go.mod` — this input
    only controls the toolchain used to install/run the govulncheck binary.

11. **`coverage-exclude` filters the PROFILE, never the test run.**
    `go-coverage.yml` and `go-integration-coverage.yml` take an optional ERE
    matched against the file paths in the coverage profile. Matching entries are
    dropped before the threshold is measured **and** before the Codecov upload,
    so the gate, the badge and the repo's own `make coverage` all report one
    number. It exists for generated code — protobuf, mocks, sqlc, ent — where a
    surviving statement is a fact about the generator, not the repo: adding one
    committed `*.pb.go` package took a real caller from 86.6% to 57.8% and
    failed a 90% gate on code no human wrote.

    `go test ./...` is deliberately left alone. Narrowing the test command
    instead would stop compiling the generated package, and a build break in
    generated code would then pass CI.

    **A pattern that matches nothing FAILS the job, and so does one that
    matches everything.** A filter silently matching zero entries is the exact
    failure this fleet keeps rediscovering — the gate looks configured, reports
    the unfiltered number, and nobody notices. The step prints the observed
    profile paths on that failure so the fix is one read, not a guess.

    The `coverage-exclude-generated` job in `ci.yml` is the self-test, and it is
    a real one: `examples/hello/generated/` is an untested generated-code
    fixture, the threshold (60) sits between the score with it counted (37.5%)
    and without (67.7%), so **removing the input turns that job red**. Do not
    add tests for that fixture — a covered fixture makes the self-test pass for
    the wrong reason.

    Sonar measures coverage independently, so a repo using this input must also
    state the exclusion in `sonar-project.properties`
    (`sonar.coverage.exclusions`) or the two will disagree about the same code.

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
