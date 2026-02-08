# BDD in Action - Go Implementation

This repository contains Go implementations of examples from "BDD in Action, Second Edition" by John Ferguson Smart.

## Project Guidelines for Claude Code

### Project Overview

This is a learning project focused on Behavior-Driven Development (BDD) practices using:

- **Go** as the programming language
- **Godog** for BDD testing (Cucumber for Go)
- **Gherkin** for writing feature specifications

### Development Philosophy

1. **BDD-First Approach**
   - Write Gherkin features before implementation
   - Let acceptance tests drive the design
   - Keep features readable by non-technical stakeholders

2. **Clean Architecture**
   - **Domain Layer**: Business entities and logic (`banking/`)
   - **Test Layer**: Organized into:
     - `features/`: Gherkin feature files
     - `testcontext/`: Test infrastructure (TestContext)
     - `domain/`: Domain models for testing (e.g., InitialAccount)
     - `actions/`: Fluent APIs for test operations
     - `stepdefinitions/`: Gherkin step mappings (organized by feature)

3. **Testing Standards**
   - All features must have corresponding step definitions
   - Use fluent APIs for better readability
   - Keep step definitions simple - directly use domain objects and fluent APIs
   - Maintain clear separation: Given/When/Then
   - Organize step definitions by feature area

### Code Organization

```
chapter-XX/
├── banking/                    # Business domain
│   ├── *.go                   # Domain entities
├── tests/
│   ├── features/              # Gherkin files (.feature)
│   │   ├── interest/          # Interest calculation features
│   │   └── transfers/         # Transfer features
│   └── acceptancetests/       # Test code
│       ├── acceptance_suite_test.go  # Test suite entry point
│       ├── testcontext/       # Test infrastructure
│       │   └── test_context.go       # Manages test state
│       ├── domain/            # Domain models for testing
│       │   └── initial_account.go    # Test data models
│       ├── actions/           # Test operations
│       │   ├── transfer_api.go       # Fluent API for transfers
│       │   └── account_type_parser.go # Helper utilities
│       └── stepdefinitions/   # Step definitions (by feature)
│           ├── transfer_steps.go     # Transfer scenarios
│           ├── interest_steps.go     # Interest scenarios
│           └── helpers.go            # Shared helpers
├── go.mod
├── Makefile                   # Common commands
└── README.md
```

### When Working on This Project

#### Adding New Features

1. **Start with Gherkin**

   ```gherkin
   Feature: Feature name
     As a <role>
     I want <goal>
     So that <benefit>

     Scenario: Scenario name
       Given <precondition>
       When <action>
       Then <expected result>
   ```

2. **Create Step Definitions**
   - Create new file in `stepdefinitions/` (e.g., `payment_steps.go`)
   - Organize by feature area (one file per feature)
   - Keep them thin - directly use domain objects and fluent APIs
   - Use descriptive error messages
   - Register steps in `acceptance_suite_test.go`

3. **Implement Fluent APIs (if needed)**
   - Add to `actions/` for complex operations
   - Use method chaining for readability
   - Follow the TransferApi pattern
   - Maintain single responsibility

4. **Update Domain**
   - Add business logic to `banking/`
   - Keep domain pure (no test dependencies)
   - Domain should be independent of test code

#### Writing Tests

**DO:**

- ✅ Use business language in features
- ✅ Keep scenarios focused and independent
- ✅ Use Scenario Outlines for data-driven tests
- ✅ Organize features by business capability
- ✅ Write descriptive scenario names

**DON'T:**

- ❌ Put business logic in step definitions
- ❌ Use technical jargon in feature files
- ❌ Create dependencies between scenarios
- ❌ Test implementation details

#### Running Tests

```bash
# Run all tests
make test

# Run with verbose output
make test-verbose

# Run specific feature
make test-interest
make test-transfers

# Generate coverage report
make test-coverage

# Format and lint
make lint
```

### Design Patterns in Use

1. **Fluent Interface** (`TransferApi`)
   - Makes test code readable and expressive
   - Uses method chaining for natural flow
   - Example:

   ```go
   transferApi.
       TheAmount(100).
       From(sourceAccount).
       To(destinationAccount)
   ```

2. **Test Context Pattern** (`TestContext`)
   - Located in dedicated `testcontext/` package
   - Maintains test state between steps
   - Provides helper methods (GetAccount, RegisterAccount)
   - Handles cleanup between scenarios
   - Separates test infrastructure from domain models

3. **Separation of Concerns**
   - Step definitions directly use domain objects
   - No intermediate "actions" layer (simplified from Java version)
   - Fluent APIs only where they add readability
   - Each step definition file focuses on one feature area

### Code Style

Follow Go conventions:

- Use `gofmt` and `goimports`
- Exported names start with uppercase
- Clear, descriptive names over comments
- Error handling at every public method

### Common Commands

```bash
# Install dependencies
make deps

# Run tests
make test

# Format code
make fmt

# Static analysis
make vet

# All checks (format, vet, build, test)
make all

# Clean artifacts
make clean
```

### Reference Implementation

This project is based on the Java implementation from:
<https://github.com/bdd-in-action/second-edition>

Key adaptations for Go:

- **Fluent interfaces**: Method chaining with pointer returns
- **Table parsing**: Manual parsing instead of `@DataTableType` annotations
- **Constructors**: Factory functions instead of class constructors
- **Error handling**: Error returns instead of exceptions
- **Step organization**: Separate files per feature area (e.g., `transfer_steps.go`, `interest_steps.go`)
- **Simplified architecture**: Step definitions directly use domain objects, avoiding unnecessary abstraction layers
- **Package structure**: Separate `testcontext` package to avoid circular dependencies

### Architectural Decisions

1. **TestContext in separate package**
   - Reason: Avoids circular dependencies between `acceptancetests` and `stepdefinitions`
   - Location: `acceptancetests/testcontext/`
   - Usage: All step definitions and fluent APIs reference it

2. **Domain models separate from test infrastructure**
   - `testcontext/`: Test infrastructure (TestContext)
   - `domain/`: Business domain models for tests (InitialAccount)
   - Clear separation of concerns

3. **One step definition file per feature**
   - `transfer_steps.go`: Transfer-related scenarios
   - `interest_steps.go`: Interest calculation scenarios
   - Better organization and maintainability
   - Follows Single Responsibility Principle

4. **Minimal abstraction**
   - No intermediate "BankingActions" layer
   - Step definitions directly use domain objects
   - Fluent APIs only where they add clarity (e.g., TransferApi)
   - Simpler than Java version, more idiomatic for Go

### Key Files to Understand

1. **`tests/acceptancetests/acceptance_suite_test.go`**
   - Entry point for Godog tests
   - Scenario initialization with TestContext
   - Registers all step definition files
   - Before/After hooks for cleanup

2. **`tests/acceptancetests/testcontext/test_context.go`**
   - Test infrastructure (separate from domain)
   - Manages test state (Client, InterestCalculator, LastError, etc.)
   - Helper methods: GetAccount, RegisterAccount, Reset
   - Used by all step definitions

3. **`tests/acceptancetests/actions/transfer_api.go`**
   - Example of fluent interface pattern
   - Shows method chaining for readability
   - Demonstrates how to make tests expressive

4. **`tests/acceptancetests/stepdefinitions/transfer_steps.go`**
   - Transfer scenario step definitions
   - Shows how to organize steps by feature
   - Directly uses domain objects and TransferApi
   - Table parsing examples

5. **`tests/acceptancetests/stepdefinitions/interest_steps.go`**
   - Interest calculation step definitions
   - Shows Given/When/Then organization
   - Demonstrates business logic delegation

6. **`tests/acceptancetests/domain/initial_account.go`**
   - Domain model for test data
   - Represents initial account setup
   - Separate from test infrastructure

### Learning Resources

- Project Documentation:
  - `QUICKSTART.md` - Getting started guide
  - `REFACTOR.md` - Refactoring documentation
  - `tests/README.md` - Detailed test documentation

- External:
  - [Godog Documentation](https://github.com/cucumber/godog)
  - [Gherkin Reference](https://cucumber.io/docs/gherkin/reference/)
  - [BDD in Action Book](https://www.manning.com/books/bdd-in-action-second-edition)

### When to Ask for Help

Claude should ask for clarification when:

- Business requirements are unclear
- Multiple valid BDD approaches exist
- Feature scope is ambiguous
- Testing strategy needs discussion

### Project-Specific Conventions

1. **Account Types**
   - Use exact names: "Current", "Savings", "SuperSaver", "Investment"
   - Case-sensitive in features

2. **Money Amounts**
   - Use float64 for flexibility
   - Display with 2 decimal places
   - Allow 0.01 epsilon for comparisons

3. **Error Handling**
   - Store business errors in TestContext.LastError
   - Return system errors immediately
   - Use descriptive error messages

4. **Test Organization**
   - One feature file per business capability
   - Group related scenarios in same feature
   - Use Background for common setup

### Anti-Patterns to Avoid

1. **Don't**: Put domain logic in test code

   ```go
   // ❌ Bad - business logic in step definition
   func (ts *TransferSteps) transfer(...) error {
       if fromAccount.Balance() < amount {
           return errors.New("insufficient funds")
       }
       fromAccount.Withdraw(amount)
       toAccount.Deposit(amount)
       return nil
   }
   ```

2. **Do**: Use domain objects or fluent APIs

   ```go
   // ✅ Good - using fluent API
   func (ts *TransferSteps) transfer(amount float64, fromType, toType string) error {
       fromAccount := ts.ctx.GetAccount(fromType)
       toAccount := ts.ctx.GetAccount(toType)
       return ts.transferApi.TheAmount(amount).From(fromAccount).To(toAccount)
   }
   ```

3. **Don't**: Mix test infrastructure with domain models

   ```go
   // ❌ Bad - TestContext in domain package
   package domain
   type TestContext struct { ... }
   type InitialAccount struct { ... }
   ```

4. **Do**: Separate concerns

   ```go
   // ✅ Good - separate packages
   package testcontext  // Test infrastructure
   type TestContext struct { ... }

   package domain       // Domain models
   type InitialAccount struct { ... }
   ```

5. **Don't**: Create unnecessary abstraction layers

   ```go
   // ❌ Bad - unnecessary wrapper layer
   type BankingActions struct { ... }
   func (ba *BankingActions) DoEverything() { ... }
   ```

6. **Do**: Keep it simple - use domain objects directly

   ```go
   // ✅ Good - direct domain usage
   func (is *InterestSteps) setRate(accountType string, rate float64) error {
       is.ctx.InterestCalculator.SetRates(accountType, rate)
       return nil
   }
   ```

### Debugging Tips

1. **Step not matching?**
   - Check regex pattern in `RegisterSteps`
   - Verify parameter types match
   - Look at Godog output for suggestions

2. **Test failing unexpectedly?**
   - Check TestContext state
   - Verify scenario independence
   - Use `make test-verbose` for details

3. **Build errors?**
   - Run `make fmt` first
   - Check `go.mod` is tidy: `make tidy`
   - Verify imports are correct

### Adding a New Feature Area

When adding a new feature (e.g., "Payments"):

1. **Create the feature file**
   ```bash
   mkdir -p tests/features/payments
   # Create payment_processing.feature
   ```

2. **Create step definitions**
   ```bash
   # Create tests/acceptancetests/stepdefinitions/payment_steps.go
   ```

   ```go
   package stepdefinitions

   type PaymentSteps struct {
       ctx *testcontext.TestContext
   }

   func NewPaymentSteps(ctx *testcontext.TestContext) *PaymentSteps {
       return &PaymentSteps{ctx: ctx}
   }

   func (ps *PaymentSteps) RegisterSteps(sc *godog.ScenarioContext) {
       sc.Step(`^...`, ps.someStep)
   }
   ```

3. **Register in test suite**
   Edit `acceptance_suite_test.go`:
   ```go
   func InitializeScenario(sc *godog.ScenarioContext) {
       testCtx := testcontext.NewTestContext()

       // Existing
       transferSteps := stepdefinitions.NewTransferSteps(testCtx)
       transferSteps.RegisterSteps(sc)

       // Add new
       paymentSteps := stepdefinitions.NewPaymentSteps(testCtx)
       paymentSteps.RegisterSteps(sc)

       // ... hooks ...
   }
   ```

4. **Add fluent API if needed**
   If the operation is complex and benefits from readability:
   ```bash
   # Create tests/acceptancetests/actions/payment_api.go
   ```

5. **Add domain models if needed**
   ```bash
   # Create tests/acceptancetests/domain/payment_details.go
   ```

### Contributing to This Project

When adding new chapters:

1. Create new directory: `chapterXX/`
2. Copy structure from `chapter02/`
3. Update Makefile with new paths
4. Add to main README.md
5. Write features first, then implementation
6. Ensure all tests pass before committing

### Success Criteria

A well-implemented BDD feature should:

- ✅ Be readable by business stakeholders
- ✅ Clearly express business value
- ✅ Have focused, independent scenarios
- ✅ Use ubiquitous language from the domain
- ✅ Pass all acceptance tests
- ✅ Have clean, maintainable step definitions
- ✅ Follow established patterns in the codebase

---

**Remember**: BDD is about collaboration and shared understanding. The code should reflect the conversations between developers, testers, and business stakeholders.
