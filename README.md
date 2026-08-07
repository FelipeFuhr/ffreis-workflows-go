# ffreis-workflows-go

<!-- ffreis-badges:start -->
[![CI](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/FelipeFuhr/ffreis-badges/main/badges/ffreis-workflows-go/ci.json)](https://github.com/FelipeFuhr/ffreis-workflows-go/actions) [![Latest version](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/FelipeFuhr/ffreis-badges/main/badges/ffreis-workflows-go/version.json)](https://github.com/FelipeFuhr/ffreis-workflows-go/releases) [![License](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/FelipeFuhr/ffreis-badges/main/badges/ffreis-workflows-go/license.json)](https://github.com/FelipeFuhr/ffreis-workflows-go/blob/main/LICENSE)
<!-- ffreis-badges:end -->

Reusable GitHub Actions workflow library for Go projects in the ffreis fleet.

All workflows use `on: workflow_call` and should be consumed from other repositories by pinning to a specific commit SHA.

```yaml
uses: FelipeFuhr/ffreis-workflows-go/.github/workflows/go-build.yml@<sha> # vX.Y.Z
```

## Workflows

| Workflow file | Purpose | Key inputs |
|---|---|---|
| `go-fmt.yml` | `gofmt`/`goimports` format check | `go-version`, `working-directory` |
| `go-lint.yml` | golangci-lint | `go-version`, `working-directory`, `config-path` |
| `go-test.yml` | `go test -race -shuffle=on` + artifact upload | `go-version`, `working-directory`, `test-args` |
| `go-build.yml` | `go build ./...` | `go-version`, `working-directory`, `build-args` |
| `go-coverage.yml` | Coverage report + Codecov upload | `go-version`, `working-directory`, `goprivate`; secrets `CODECOV_TOKEN`, `GIT_AUTH_TOKEN` |
| `go-integration-coverage.yml` | Integration-tagged coverage report + Codecov upload (separate metric from unit coverage; opt-in gate via `coverage-threshold`) | `go-version`, `working-directory`, `build-tag`, `coverage-threshold`, `goprivate`; secrets `CODECOV_TOKEN`, `GIT_AUTH_TOKEN` |
| `go-security.yml` | `govulncheck` CVE scan | `go-version`, `working-directory` |
| `go-semgrep.yml` | Semgrep SAST (no upload) | `working-directory` |
| `go-semgrep-sarif.yml` | Semgrep SAST with SARIF upload | `working-directory`, `upload-sarif` |
| `go-sonar.yml` | SonarCloud scan + coverage | `go-version`, `working-directory`; secret `SONAR_TOKEN` |
| `go-docs.yml` | `go doc` / pkgsite validation | `go-version`, `working-directory` |
| `go-fuzz.yml` | `go test -fuzz` corpus run | `go-version`, `working-directory`, `fuzz-targets`, `fuzz-time` |
| `go-mutation.yml` | go-mutants mutation testing | `go-version`, `working-directory`, `packages`, `threshold`, `goprivate`; secret `GIT_AUTH_TOKEN` |
| `go-mod-tidy-check.yml` | `go mod tidy` idempotency check | `go-version`, `working-directory` |
| `go-osv-scanner.yml` | OSV vulnerability scanner | `working-directory` |
| `go-snyk.yml` | Snyk dependency scan | `working-directory`; secret `SNYK_TOKEN` |
| `go-deepsource-coverage.yml` | DeepSource coverage upload | `working-directory`; secret `DEEPSOURCE_DSN` |
| `go-cross-build-matrix.yml` | Cross-compilation matrix | `go-version`, `working-directory`, `targets` |
| `go-quick-checks.yml` | Bundled fmt+lint+vet for fast PR feedback | `go-version`, `working-directory` |
| `go-container.yml` | Container image build from Go binary | `image`, `containerfile`, `working-directory`, `push` |
