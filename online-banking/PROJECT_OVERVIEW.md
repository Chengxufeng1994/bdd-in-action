# Online Banking Application - Complete Project Overview

Generated: 2026-02-08

## Table of Contents

1. [Project Structure](#project-structure)
2. [Domain Models](#domain-models)
3. [Test Infrastructure](#test-infrastructure)
4. [Feature Files](#feature-files)
5. [Configuration Files](#configuration-files)
6. [Documentation](#documentation)

---

## Project Structure

```
online-banking/
├── banking/                          # Domain layer
│   ├── account_type.go
│   ├── bank_account.go
│   ├── client.go
│   ├── interest_calculator.go
│   └── transaction.go
├── tests/
│   ├── acceptancetests/             # BDD test infrastructure
│   │   ├── actions/                 # Test actions (fluent API)
│   │   │   ├── account_type_parser.go
│   │   │   └── transfer_api.go
│   │   ├── domain/                  # Test domain helpers
│   │   │   └── initial_account.go
│   │   ├── stepdefinitions/         # Cucumber step definitions
│   │   │   ├── helpers.go
│   │   │   ├── interest_steps.go
│   │   │   └── transfer_steps.go
│   │   ├── testcontext/             # Test state management
│   │   │   └── test_context.go
│   │   └── acceptance_suite_test.go # Test suite configuration
│   ├── features/                    # Gherkin feature files
│   │   ├── interest/
│   │   │   └── earning_interest.feature
│   │   └── transfers/
│   │       ├── transferring_between_accounts.feature
│   │       └── transferring_to_savings.feature
│   └── README.md
├── go.mod                           # Go module definition
├── go.sum                           # Dependency checksums
├── Makefile                         # Build and test automation
├── README.md                        # Main documentation
├── QUICKSTART.md                    # Quick start guide
└── REFACTOR.md                      # Refactoring documentation

Total Files: 23
```

---

## Domain Models

### 1. Account Type (`banking/account_type.go`)

```go
package banking

type AccountType string

const (
	AccountTypeCurrent    AccountType = "Current"
	AccountTypeSavings    AccountType = "Savings"
	AccountTypeInvestment AccountType = "Investment"
	AccountTypeSuperSaver AccountType = "SuperSaver"
)

func (at AccountType) String() string {
	return string(at)
}
```

**Purpose**: Enum-like type for different account types
**Design**: Uses const with string type for readability

---

### 2. Bank Account (`banking/bank_account.go`)

```go
package banking

type BankAccount struct {
	accountType AccountType
	balance     float64
}

func NewBankAccount(accountType AccountType) *BankAccount {
	return &BankAccount{
		accountType: accountType,
		balance:     0,
	}
}

func BankAccountOfType(accountType AccountType) *BankAccount {
	return NewBankAccount(accountType)
}

func (ba *BankAccount) WithBalance(balance float64) *BankAccount {
	ba.balance = balance
	return ba
}

func (ba *BankAccount) AccountType() AccountType {
	return ba.accountType
}

func (ba *BankAccount) Balance() float64 {
	return ba.balance
}

func (ba *BankAccount) Deposit(amount float64) {
	ba.balance += amount
}

func (ba *BankAccount) Withdraw(amount float64) error {
	if ba.balance < amount {
		return ErrInsufficientFunds
	}
	ba.balance -= amount
	return nil
}

func (ba *BankAccount) RecordTransaction(transaction Transaction) {
	ba.balance += transaction.Amount()
}

var ErrInsufficientFunds = &InsufficientFundsError{}

type InsufficientFundsError struct{}

func (e *InsufficientFundsError) Error() string {
	return "insufficient funds"
}
```

**Purpose**: Core bank account entity
**Design Patterns**:
- Fluent interface with `WithBalance()`
- Factory methods: `NewBankAccount()`, `BankAccountOfType()`
- Encapsulation: private fields with public methods
- Custom error type for domain-specific errors

---

### 3. Client (`banking/client.go`)

```go
package banking

type Client struct {
	name     string
	accounts map[AccountType]*BankAccount
}

func NewClient(name string) *Client {
	return &Client{
		name:     name,
		accounts: make(map[AccountType]*BankAccount),
	}
}

func (c *Client) Name() string {
	return c.name
}

func (c *Client) Opens(account *BankAccount) {
	c.accounts[account.AccountType()] = account
}

func (c *Client) Get(accountType AccountType) *BankAccount {
	return c.accounts[accountType]
}

func (c *Client) Accounts() []*BankAccount {
	accounts := make([]*BankAccount, 0, len(c.accounts))
	for _, account := range c.accounts {
		accounts = append(accounts, account)
	}
	return accounts
}
```

**Purpose**: Represents a bank client with multiple accounts
**Design**:
- Manages collection of accounts by type
- Defensive copy in `Accounts()` method

---

### 4. Interest Calculator (`banking/interest_calculator.go`)

```go
package banking

import "math"

type InterestCalculator struct {
	interestRates map[AccountType]float64
}

func NewInterestCalculator() *InterestCalculator {
	return &InterestCalculator{
		interestRates: make(map[AccountType]float64),
	}
}

func (ic *InterestCalculator) SetRates(accountType AccountType, rate float64) {
	ic.interestRates[accountType] = rate
}

func (ic *InterestCalculator) CalculateMonthlyInterestOn(account *BankAccount) Transaction {
	rate := ic.interestRates[account.AccountType()]
	monthlyInterest := account.Balance() * rate / 100 / 12

	// Round to 2 decimal places
	monthlyInterest = math.Round(monthlyInterest*100) / 100

	transaction := NewTransaction(
		ic.getCurrentTime(),
		"INTEREST",
		monthlyInterest,
	)

	account.RecordTransaction(transaction)

	return transaction
}

func (ic *InterestCalculator) getCurrentTime() interface{} {
	return nil // Simplified for this example
}
```

**Purpose**: Calculate and apply interest to accounts
**Design**:
- Configurable interest rates per account type
- Rounds to 2 decimal places for currency
- Returns transaction for audit trail

---

### 5. Transaction (`banking/transaction.go`)

```go
package banking

type Transaction struct {
	timestamp   interface{}
	description string
	amount      float64
}

func NewTransaction(timestamp interface{}, description string, amount float64) Transaction {
	return Transaction{
		timestamp:   timestamp,
		description: description,
		amount:      amount,
	}
}

func (t Transaction) Description() string {
	return t.description
}

func (t Transaction) Amount() float64 {
	return t.amount
}
```

**Purpose**: Immutable record of financial transactions
**Design**: Value object pattern (returned by value, not pointer)

---

## Test Infrastructure

### 1. Test Context (`tests/acceptancetests/testcontext/test_context.go`)

```go
package testcontext

import "github.com/Chegnxufeng1994/bdd-in-action/online-banking/banking"

type TestContext struct {
	Client              *banking.Client
	InterestCalculator  *banking.InterestCalculator
	LastError           error
	LastTransaction     banking.Transaction
	Accounts            map[banking.AccountType]*banking.BankAccount
}

func NewTestContext() *TestContext {
	return &TestContext{
		Client:             banking.NewClient("Tess"),
		InterestCalculator: banking.NewInterestCalculator(),
		Accounts:           make(map[banking.AccountType]*banking.BankAccount),
	}
}

func (tc *TestContext) GetAccount(accountType banking.AccountType) *banking.BankAccount {
	return tc.Accounts[accountType]
}

func (tc *TestContext) RegisterAccount(accountType banking.AccountType, account *banking.BankAccount) {
	tc.Accounts[accountType] = account
}

func (tc *TestContext) Reset() {
	tc.Client = banking.NewClient("Tess")
	tc.InterestCalculator = banking.NewInterestCalculator()
	tc.LastError = nil
	tc.Accounts = make(map[banking.AccountType]*banking.BankAccount)
}
```

**Purpose**: Manages test state across scenarios
**Location**: Separate package to avoid circular dependencies
**Design**:
- Centralized state management
- Reset method for scenario isolation
- Account registry for step definitions

---

### 2. Transfer API (`tests/acceptancetests/actions/transfer_api.go`)

```go
package actions

import (
	"github.com/Chegnxufeng1994/bdd-in-action/online-banking/banking"
	"github.com/Chegnxufeng1994/bdd-in-action/online-banking/tests/acceptancetests/testcontext"
)

type TransferApi struct {
	ctx        *testcontext.TestContext
	amount     float64
	fromAccount *banking.BankAccount
}

func NewTransferApi(ctx *testcontext.TestContext) *TransferApi {
	return &TransferApi{
		ctx: ctx,
	}
}

func (ta *TransferApi) TheAmount(amount float64) *TransferApi {
	ta.amount = amount
	return ta
}

func (ta *TransferApi) From(account *banking.BankAccount) *TransferApi {
	ta.fromAccount = account
	return ta
}

func (ta *TransferApi) To(account *banking.BankAccount) error {
	if err := ta.fromAccount.Withdraw(ta.amount); err != nil {
		ta.ctx.LastError = err
		return err
	}
	account.Deposit(ta.amount)
	return nil
}
```

**Purpose**: Fluent interface for transfer operations
**Pattern**: Builder pattern for readable test code
**Usage**: `transferApi.TheAmount(200).From(current).To(savings)`

---

### 3. Account Type Parser (`tests/acceptancetests/actions/account_type_parser.go`)

```go
package actions

import (
	"fmt"
	"strings"

	"github.com/Chegnxufeng1994/bdd-in-action/online-banking/banking"
)

func ParseAccountType(accountTypeStr string) (banking.AccountType, error) {
	normalized := strings.TrimSpace(accountTypeStr)

	switch normalized {
	case "Current", "current":
		return banking.AccountTypeCurrent, nil
	case "Savings", "savings":
		return banking.AccountTypeSavings, nil
	case "Investment", "investment":
		return banking.AccountTypeInvestment, nil
	case "SuperSaver", "super saver", "Super Saver":
		return banking.AccountTypeSuperSaver, nil
	default:
		return "", fmt.Errorf("unknown account type: %s", accountTypeStr)
	}
}
```

**Purpose**: Convert Gherkin strings to AccountType enum
**Design**: Handles multiple string formats for flexibility

---

### 4. Step Definitions - Transfer (`tests/acceptancetests/stepdefinitions/transfer_steps.go`)

```go
package stepdefinitions

import (
	"fmt"

	"github.com/Chegnxufeng1994/bdd-in-action/online-banking/banking"
	"github.com/Chegnxufeng1994/bdd-in-action/online-banking/tests/acceptancetests/actions"
	"github.com/Chegnxufeng1994/bdd-in-action/online-banking/tests/acceptancetests/testcontext"
	"github.com/cucumber/godog"
)

type TransferSteps struct {
	ctx         *testcontext.TestContext
	transferApi *actions.TransferApi
}

func NewTransferSteps(ctx *testcontext.TestContext) *TransferSteps {
	return &TransferSteps{
		ctx:         ctx,
		transferApi: actions.NewTransferApi(ctx),
	}
}

func (ts *TransferSteps) RegisterSteps(sc *godog.ScenarioContext) {
	sc.Step(`^(\w+) has a (\w+) account with \$(\d+)$`, ts.clientHasAccountWithBalance)
	sc.Step(`^she has \$(\d+(?:\.\d+)?) in her (\w+) account$`, ts.hasBalanceInAccount)
	sc.Step(`^she transfers \$(\d+(?:\.\d+)?) from her (\w+) account to her (\w+) account$`, ts.transfersBetweenAccounts)
	sc.Step(`^the transfer should fail due to insufficient funds$`, ts.transferShouldFailDueToInsufficientFunds)
}

func (ts *TransferSteps) clientHasAccountWithBalance(clientName, accountTypeStr string, balance int) error {
	accountType, err := actions.ParseAccountType(accountTypeStr)
	if err != nil {
		return err
	}

	account := banking.BankAccountOfType(accountType).WithBalance(float64(balance))
	ts.ctx.RegisterAccount(accountType, account)
	ts.ctx.Client.Opens(account)

	return nil
}

func (ts *TransferSteps) hasBalanceInAccount(expectedBalance float64, accountTypeStr string) error {
	accountType, err := actions.ParseAccountType(accountTypeStr)
	if err != nil {
		return err
	}

	account := ts.ctx.GetAccount(accountType)
	if account == nil {
		return fmt.Errorf("account %s not found", accountType)
	}

	actualBalance := account.Balance()
	if actualBalance != expectedBalance {
		return fmt.Errorf("expected balance %.2f, got %.2f", expectedBalance, actualBalance)
	}

	return nil
}

func (ts *TransferSteps) transfersBetweenAccounts(amount float64, fromAccountStr, toAccountStr string) error {
	fromAccountType, err := actions.ParseAccountType(fromAccountStr)
	if err != nil {
		return err
	}

	toAccountType, err := actions.ParseAccountType(toAccountStr)
	if err != nil {
		return err
	}

	fromAccount := ts.ctx.GetAccount(fromAccountType)
	toAccount := ts.ctx.GetAccount(toAccountType)

	return ts.transferApi.TheAmount(amount).From(fromAccount).To(toAccount)
}

func (ts *TransferSteps) transferShouldFailDueToInsufficientFunds() error {
	if ts.ctx.LastError == nil {
		return fmt.Errorf("expected transfer to fail with insufficient funds error, but it succeeded")
	}

	if _, ok := ts.ctx.LastError.(*banking.InsufficientFundsError); !ok {
		return fmt.Errorf("expected InsufficientFundsError, got %T: %v", ts.ctx.LastError, ts.ctx.LastError)
	}

	return nil
}
```

**Purpose**: Maps Gherkin steps to Go code for transfers
**Pattern**: Step definition object per feature area

---

### 5. Step Definitions - Interest (`tests/acceptancetests/stepdefinitions/interest_steps.go`)

```go
package stepdefinitions

import (
	"fmt"
	"math"

	"github.com/Chegnxufeng1994/bdd-in-action/online-banking/tests/acceptancetests/actions"
	"github.com/Chegnxufeng1994/bdd-in-action/online-banking/tests/acceptancetests/testcontext"
	"github.com/cucumber/godog"
)

type InterestSteps struct {
	ctx *testcontext.TestContext
}

func NewInterestSteps(ctx *testcontext.TestContext) *InterestSteps {
	return &InterestSteps{
		ctx: ctx,
	}
}

func (is *InterestSteps) RegisterSteps(sc *godog.ScenarioContext) {
	sc.Step(`^the interest rate for (\w+) accounts is (\d+(?:\.\d+)?)$`, is.interestRateIs)
	sc.Step(`^the monthly interest is calculated$`, is.monthlyInterestIsCalculated)
	sc.Step(`^she should have earned \$(\d+(?:\.\d+)?)$`, is.shouldHaveEarned)
}

func (is *InterestSteps) interestRateIs(accountTypeStr string, rate float64) error {
	accountType, err := actions.ParseAccountType(accountTypeStr)
	if err != nil {
		return err
	}

	is.ctx.InterestCalculator.SetRates(accountType, rate)
	return nil
}

func (is *InterestSteps) monthlyInterestIsCalculated() error {
	for _, account := range is.ctx.Client.Accounts() {
		transaction := is.ctx.InterestCalculator.CalculateMonthlyInterestOn(account)
		is.ctx.LastTransaction = transaction
	}
	return nil
}

func (is *InterestSteps) shouldHaveEarned(expectedInterest float64) error {
	actualInterest := is.ctx.LastTransaction.Amount()

	if math.Abs(actualInterest-expectedInterest) > 0.01 {
		return fmt.Errorf("expected interest %.2f, got %.2f", expectedInterest, actualInterest)
	}

	return nil
}
```

**Purpose**: Maps Gherkin steps to Go code for interest calculations
**Note**: Uses 0.01 tolerance for floating-point comparison

---

### 6. Test Suite Configuration (`tests/acceptancetests/acceptance_suite_test.go`)

```go
package acceptancetests

import (
	"context"
	"flag"
	"os"
	"testing"

	"github.com/Chegnxufeng1994/bdd-in-action/online-banking/tests/acceptancetests/stepdefinitions"
	"github.com/Chegnxufeng1994/bdd-in-action/online-banking/tests/acceptancetests/testcontext"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

var opts = godog.Options{
	Output: colors.Colored(os.Stdout),
	Format: "pretty",
}

func init() {
	godog.BindCommandLineFlags("godog.", &opts)
}

func TestMain(m *testing.M) {
	flag.Parse()
	opts.Paths = []string{"../features"}

	os.Exit(m.Run())
}

func TestFeatures(t *testing.T) {
	opts.TestingT = t

	status := godog.TestSuite{
		Name:                "Banking Features",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()

	if status == 2 {
		t.SkipNow()
	}

	if status != 0 {
		t.Fatalf("zero status code expected, %d received", status)
	}
}

func TestFeaturesWithCucumberReport(t *testing.T) {
	// Create output file for cucumber JSON report
	outputFile, err := os.Create("cucumber-report.json")
	if err != nil {
		t.Fatalf("failed to create cucumber report file: %v", err)
	}
	defer outputFile.Close()

	// Configure options for cucumber JSON format
	cucumberOpts := godog.Options{
		Format:   "cucumber",
		Output:   outputFile,
		Paths:    []string{"../features"},
		TestingT: t,
	}

	status := godog.TestSuite{
		Name:                "Banking Features",
		ScenarioInitializer: InitializeScenario,
		Options:             &cucumberOpts,
	}.Run()

	if status == 2 {
		t.SkipNow()
	}

	if status != 0 {
		t.Fatalf("zero status code expected, %d received", status)
	}
}

func InitializeScenario(sc *godog.ScenarioContext) {
	testCtx := testcontext.NewTestContext()

	// Register transfer step definitions
	transferSteps := stepdefinitions.NewTransferSteps(testCtx)
	transferSteps.RegisterSteps(sc)

	// Register interest step definitions
	interestSteps := stepdefinitions.NewInterestSteps(testCtx)
	interestSteps.RegisterSteps(sc)

	// Reset context before each scenario
	sc.Before(func(ctx context.Context, scenario *godog.Scenario) (context.Context, error) {
		return ctx, nil
	})

	// Cleanup after each scenario
	sc.After(func(ctx context.Context, scenario *godog.Scenario, err error) (context.Context, error) {
		testCtx.Reset()
		return ctx, nil
	})
}
```

**Purpose**: Godog test suite configuration
**Features**:
- Two test functions: one for console output, one for JSON report
- Scenario hooks for setup/cleanup
- TestMain for flag parsing

---

## Feature Files

### 1. Earning Interest (`tests/features/interest/earning_interest.feature`)

```gherkin
Feature: Earning Interest

  Scenario Outline: Earning interest
    Given Tess has a <Account Type> account with $10000
    And the interest rate for <Account Type> accounts is <Interest Rate>
    When the monthly interest is calculated
    Then she should have earned $<Expected Interest>
    Then she should have $<Expected Balance> in her <Account Type> account

    Examples:
      | Account Type | Interest Rate | Expected Interest | Expected Balance |
      | Current      | 1.0           | 8.33              | 10008.33         |
      | Savings      | 3.0           | 25.0              | 10025.0          |
```

**Purpose**: Test interest calculation for different account types
**Pattern**: Scenario Outline with Examples table

---

### 2. Transferring Between Accounts (`tests/features/transfers/transferring_between_accounts.feature`)

```gherkin
Feature: Transferring money between accounts within the bank

  Scenario: Transferring money to a savings account
    Given Tess has a Current account with $1000
    And Tess has a Savings account with $2000
    When she transfers $100.0 from her Current account to her Savings account
    Then she has $900.0 in her Current account
    And she has $2100.0 in her Savings account
```

**Purpose**: Test money transfers between accounts
**Pattern**: Simple scenario with Given-When-Then

---

### 3. Transferring with Insufficient Funds (`tests/features/transfers/transferring_to_savings.feature`)

```gherkin
Feature: Transferring with insufficient funds

  Scenario: Transferring with insufficient funds
    Given Tess has a Current account with $100
    And Tess has a Savings account with $0
    When she transfers $200.0 from her Current account to her Savings account
    Then the transfer should fail due to insufficient funds
```

**Purpose**: Test error handling for insufficient funds
**Pattern**: Negative test scenario

---

## Configuration Files

### 1. Go Module (`go.mod`)

```go
module github.com/Chegnxufeng1994/bdd-in-action/chapter02

go 1.24.0

require (
	github.com/cucumber/godog v0.15.0
	github.com/cucumber/messages/go/v24 v24.1.0
)

require (
	github.com/cucumber/gherkin/go/v28 v28.0.0 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-memdb v1.3.4 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)
```

**Dependencies**:
- `godog v0.15.0`: BDD testing framework
- `cucumber/messages`: Cucumber protocol messages

---

### 2. Makefile

```makefile
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

test-interest: ## Run only interest feature tests
	@echo "Running interest feature tests..."
	@go test -v ./tests/acceptancetests -run "TestFeatures.*interest"

test-transfers: ## Run only transfer feature tests
	@echo "Running transfer feature tests..."
	@go test -v ./tests/acceptancetests -run "TestFeatures.*transfer"

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
```

**Key Targets**:
- `make test`: Run BDD tests
- `make test-cucumber-report`: Generate coverage JSON
- `make lint`: Format and vet code
- `make clean`: Remove artifacts

---

## Documentation

### README.md
Main project documentation covering:
- Project structure
- Running tests
- Cucumber coverage
- Architecture patterns
- Development commands

### QUICKSTART.md
Quick reference for:
- Installation
- Running tests
- Common commands

### REFACTOR.md
Documents the refactoring from Java to Go:
- Original Java patterns
- Go adaptations
- Design decisions

### tests/README.md
Test-specific documentation:
- Test organization
- Writing new features
- Step definition patterns

---

## Statistics

- **Total Files**: 23
- **Go Source Files**: 15
- **Feature Files**: 3
- **Scenarios**: 6
- **Steps**: 29
- **Lines of Code**: ~800 (excluding tests)
- **Test Code**: ~400 lines

---

## Key Design Patterns

1. **Fluent Interface**: `TransferApi` for readable test code
2. **Builder Pattern**: `WithBalance()` for object construction
3. **Factory Pattern**: `NewBankAccount()`, `BankAccountOfType()`
4. **Repository Pattern**: `Client` manages account collection
5. **Value Object**: `Transaction` (immutable)
6. **Test Context**: Centralized test state management
7. **Step Definition Organization**: By feature area

---

## Testing Strategy

- **BDD with Godog**: Gherkin features drive development
- **Scenario Isolation**: Each scenario gets fresh context
- **Fluent Test API**: Readable, business-focused test code
- **Coverage Reporting**: JSON format for CI/CD integration
- **Separation of Concerns**: Test infrastructure separate from domain

---

## Build and Test Commands

```bash
# Development
make build              # Build all packages
make test              # Run all acceptance tests
make test-verbose      # Run with detailed output
make lint              # Format and vet code

# Coverage
make test-cucumber-report   # Generate JSON coverage report

# Specific Features
make test-interest     # Run interest tests only
make test-transfers    # Run transfer tests only

# Maintenance
make clean             # Remove artifacts
make deps              # Update dependencies
make all               # Lint + Build + Test
```

---

**Document Generated**: 2026-02-08
**Project**: BDD in Action - Chapter 2
**Framework**: Godog (Cucumber for Go)
**Go Version**: 1.24+
