# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Go implementations of examples from "BDD in Action, Second Edition" by John Ferguson Smart. Each chapter is an **independent Go module** with its own `go.mod`, `Makefile`, and test suite.

| Directory | Chapter | Domain |
|-----------|---------|--------|
| `online-banking/` | Chapter 2 | Bank accounts, transfers, interest |
| `train-timetables/` | Chapter 3 | Train schedules, itinerary search |

## Commands

All commands run from within the module directory (e.g., `cd online-banking`):

```bash
make test              # Run all acceptance tests
make test-verbose      # Run with verbose output
make test-cucumber-report  # Generate cucumber-report.json
make lint              # Format (gofmt + goimports) and go vet
make all               # lint + build + test
```

Run a specific feature tag (from within module dir):
```bash
go test -v ./tests/acceptancetests -run "TestFeatures.*interest"
go test -v ./tests/acceptancetests -run "TestFeatures.*transfer"
```

## Architecture

Each module follows this layout:

```
<module>/
├── <domain>/           # Pure business logic (no test dependencies)
├── tests/
│   ├── features/       # Gherkin .feature files (organized by capability)
│   └── acceptancetests/
│       ├── acceptance_suite_test.go  # Entry point: InitializeScenario()
│       ├── testcontext/              # TestContext struct — shared state between steps
│       ├── stepdefinitions/          # One file per feature area (e.g., transfer_steps.go)
│       ├── actions/                  # Fluent APIs for complex multi-step operations
│       └── domain/                   # Test-only domain models (e.g., InitialAccount)
```

### Key Patterns

**TestContext** (`testcontext/test_context.go`): Holds all mutable state across BDD steps within a scenario. Reset via `sc.After` hook. Step definitions receive it via constructor injection.

**Step definition registration**: Each `*Steps` struct has a `RegisterSteps(sc *godog.ScenarioContext)` method. Add new step structs to `InitializeScenario()` in `acceptance_suite_test.go`.

**Fluent API** (`actions/`): Used when operation chains improve readability. Example:
```go
transferApi.TheAmount(100).From(currentAccount).To(savingsAccount)
```
Only create fluent APIs where method chaining genuinely aids readability — direct domain calls are preferred otherwise.

### BDD Workflow

1. Write Gherkin in `tests/features/<capability>/`
2. Create `stepdefinitions/<feature>_steps.go` with a struct, constructor, and `RegisterSteps`
3. Register in `acceptance_suite_test.go`
4. Implement domain logic in `<domain>/`
5. Errors go in `testcontext.LastError`; system errors fail immediately

### Adding a New Chapter Module

1. Create directory: `<chapter-name>/`
2. Copy `online-banking/` structure as template
3. Update module name in `go.mod`

## Conventions

- **Account types** (online-banking): exact strings `"Current"`, `"Savings"`, `"SuperSaver"`, `"Investment"` — case-sensitive in features
- **Money**: `float64`; allow 0.01 epsilon for comparisons
- **Gherkin language**: `# language: zh-TW` supported (see train-timetables features)
