# Online Banking Application with Godog

This chapter demonstrates Behavior-Driven Development (BDD) using Godog and Gherkin for a banking application.

## Project Structure

```
online-banking/
├── banking/                    # Domain models
│   ├── account_type.go        # Account type enum
│   ├── bank_account.go        # Bank account entity
│   ├── client.go              # Client entity
│   ├── interest_calculator.go # Interest calculation logic
│   └── transaction.go         # Transaction entity
├── tests/
│   ├── acceptancetests/       # BDD acceptance tests
│   │   ├── actions/           # Test actions (fluent API)
│   │   │   ├── account_type_parser.go
│   │   │   └── transfer_api.go
│   │   ├── stepdefinitions/   # Cucumber step definitions
│   │   │   ├── interest_steps.go
│   │   │   └── transfer_steps.go
│   │   ├── testcontext/       # Test state management
│   │   │   └── test_context.go
│   │   └── acceptance_suite_test.go
│   └── features/              # Gherkin feature files
│       ├── interest/
│       │   └── earning_interest.feature
│       └── transfers/
│           ├── transferring_between_accounts.feature
│           ├── transferring_to_savings.feature
│           └── transferring_with_insufficient_funds.feature
└── Makefile
```

## Running Tests

### Run all acceptance tests

```bash
make test
```

### Run with verbose output

```bash
make test-verbose
```

### Generate Cucumber coverage report

```bash
make test-cucumber-report
```

This generates a JSON report at `tests/acceptancetests/cucumber-report.json` containing:
- All features and scenarios executed
- Step definitions with file locations (file:line)
- Pass/fail status for each step
- Execution duration in nanoseconds

### Run specific feature tests

```bash
# Interest features only
make test-interest

# Transfer features only
make test-transfers
```

## Cucumber Coverage

The Cucumber JSON report provides detailed information about BDD test execution:

- **Features**: Lists all `.feature` files executed with URIs
- **Scenarios**: Shows which scenarios passed/failed with line numbers
- **Steps**: Maps Gherkin steps to Go step definition implementations
- **Coverage**: Identifies which step definitions are used and where they're defined (file:line)
- **Performance**: Execution duration for each step in nanoseconds

The JSON format follows the standard Cucumber JSON schema and can be used for:
- Coverage analysis and reporting
- CI/CD pipeline integration
- Custom visualization tools
- Jenkins Cucumber Reports Plugin

## Architecture Patterns

### Fluent API

Transfer operations use a fluent interface for readability:

```go
transferApi.TheAmount(200.0).From(currentAccount).To(savingsAccount)
```

### Test Context

Centralized test state management in `testcontext` package:

```go
type TestContext struct {
    Client              *banking.Client
    InterestCalculator  *banking.InterestCalculator
    LastError           error
    LastTransaction     banking.Transaction
    Accounts            map[banking.AccountType]*banking.BankAccount
}
```

### Step Definitions by Feature

Step definitions are organized by feature area:
- `transfer_steps.go`: Account creation, balance verification, transfers
- `interest_steps.go`: Interest rate configuration, interest calculation

## Development Commands

```bash
# Format code
make fmt

# Run linter
make vet

# Run all checks
make lint

# Build packages
make build

# Clean artifacts
make clean

# Update dependencies
make deps

# Run everything
make all
```

## Dependencies

- [Godog](https://github.com/cucumber/godog) - Cucumber BDD framework for Go
- Go 1.24+ (for testing.T support)

## References

- Based on "BDD in Action" (Second Edition) by John Ferguson Smart
- Original Java implementation: https://github.com/bdd-in-action/second-edition/tree/master/online-banking
