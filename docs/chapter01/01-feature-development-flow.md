# Chapter 1: Feature Development Flow in BDD

## Overview

This document describes the complete flow when Chris (or any team member) wants to add a new feature using Behavior-Driven Development (BDD) practices.

## The BDD Feature Development Flow

### Phase 1: Discovery & Collaboration

#### 1.1 Identify the Need

- **Who**: Product Owner, Stakeholder, or Developer (Chris)
- **What**: Identify a business need or user requirement
- **Output**: Initial feature idea or user story

**Example:**

```
As a bank customer
I want to transfer money to another account
So that I can send money to my friends and family
```

#### 1.2 Collaborative Discussion (Three Amigos)

- **Participants**:
  - Business Analyst/Product Owner (What needs to be built)
  - Developer (How it will be built)
  - Tester (What could go wrong)

- **Activities**:
  - Discuss acceptance criteria
  - Identify edge cases
  - Clarify business rules
  - Explore examples

- **Output**: Shared understanding of the feature

### Phase 2: Example Discovery

#### 2.1 Write Concrete Examples

Transform abstract requirements into concrete examples using real data.

**Poor (Abstract):**

```
The system should validate the transfer amount
```

**Good (Concrete):**

```gherkin
Scenario: Successful transfer with sufficient balance
  Given Chris has $100 in his checking account
  And Lisa has $50 in her savings account
  When Chris transfers $30 to Lisa's account
  Then Chris should have $70 in his checking account
  And Lisa should have $80 in her savings account
```

#### 2.2 Identify Multiple Scenarios

Cover the main flow and alternative flows:

- **Happy Path**: Normal successful execution
- **Alternative Paths**: Different valid ways to achieve the goal
- **Error Paths**: What happens when things go wrong

**Example Scenarios:**

```gherkin
Feature: Money Transfer

  Scenario: Successful transfer with sufficient balance
    # Happy path

  Scenario: Transfer fails with insufficient balance
    # Error path
    Given Chris has $100 in his checking account
    When Chris attempts to transfer $150 to Lisa's account
    Then the transfer should be rejected
    And Chris should see error "Insufficient balance"
    And Chris should still have $100 in his checking account

  Scenario: Transfer to the same account
    # Edge case
    Given Chris has $100 in his checking account
    When Chris attempts to transfer $50 to his own account
    Then the transfer should be rejected
    And Chris should see error "Cannot transfer to the same account"
```

### Phase 3: Formalize as Gherkin

#### 3.1 Write Feature File

Create a `.feature` file using Gherkin syntax:

**File**: `features/money_transfer.feature`

```gherkin
Feature: Money Transfer
  As a bank customer
  I want to transfer money between accounts
  So that I can send money to others

  Background:
    Given the following accounts exist:
      | account_id | owner | balance |
      | ACC001     | Chris | 1000    |
      | ACC002     | Lisa  | 500     |

  Scenario: Transfer money successfully
    Given Chris is logged in
    When Chris transfers $300 from account "ACC001" to account "ACC002"
    Then the transfer should succeed
    And account "ACC001" should have balance of $700
    And account "ACC002" should have balance of $800
    And Chris should see confirmation "Transfer successful"

  Scenario: Insufficient balance
    Given Chris is logged in
    When Chris transfers $1500 from account "ACC001" to account "ACC002"
    Then the transfer should fail
    And Chris should see error "Insufficient balance"
    And account "ACC001" should have balance of $1000
    And account "ACC002" should have balance of $500

  Scenario Outline: Transfer amount validation
    Given Chris is logged in
    When Chris transfers <amount> from account "ACC001" to account "ACC002"
    Then the transfer should <result>
    And Chris should see <message>

    Examples:
      | amount | result  | message                        |
      | $0     | fail    | "Amount must be greater than 0" |
      | $-100  | fail    | "Amount cannot be negative"     |
      | $0.01  | succeed | "Transfer successful"           |
```

#### 3.2 Review and Refine

- Ensure scenarios are independent
- Avoid implementation details
- Focus on behavior, not technical implementation
- Use business language, not technical jargon

### Phase 4: Implementation Planning

#### 4.1 Create Implementation Plan

Following the development guidelines, create `IMPLEMENTATION_PLAN.md`:

```markdown
# Money Transfer Feature Implementation

## Stage 1: Domain Model

**Goal**: Create core domain entities and value objects
**Success Criteria**:

- Account entity with balance tracking
- Transfer entity with validation rules
- Unit tests passing

**Tests**:

- Account creation and balance updates
- Transfer validation (amount, accounts)
- Business rule enforcement

**Status**: Not Started

## Stage 2: Step Definitions

**Goal**: Implement Gherkin step definitions
**Success Criteria**:

- All Gherkin steps have corresponding Go functions
- Steps can execute scenarios
- Scenarios fail (RED phase)

**Tests**:

- Step definition integration tests
- Scenario execution framework

**Status**: Not Started

## Stage 3: Business Logic

**Goal**: Implement transfer service
**Success Criteria**:

- Transfer service completes transfers
- All business rules enforced
- All scenarios pass (GREEN phase)

**Tests**:

- Service integration tests
- All BDD scenarios passing

**Status**: Not Started

## Stage 4: Repository Integration

**Goal**: Integrate with data persistence
**Success Criteria**:

- Repository implementation
- Transaction management
- Data consistency tests passing

**Tests**:

- Repository integration tests
- Transaction rollback scenarios

**Status**: Not Started

## Stage 5: API Layer

**Goal**: Expose transfer functionality via API
**Success Criteria**:

- REST API endpoint
- Request/response validation
- API tests passing

**Tests**:

- API endpoint tests
- End-to-end scenarios

**Status**: Not Started
```

#### 4.2 Break Down into Tasks

Identify specific implementation tasks for each stage.

### Phase 5: Test-Driven Development (TDD)

#### 5.1 RED - Write Failing Tests

**Step Definition Test (RED):**

```go
// features/steps/money_transfer_test.go
func TestMoneyTransfer(t *testing.T) {
    suite := godog.TestSuite{
        ScenarioInitializer: InitializeScenario,
        Options: &godog.Options{
            Format:   "pretty",
            Paths:    []string{"../money_transfer.feature"},
            TestingT: t,
        },
    }

    if suite.Run() != 0 {
        t.Fatal("non-zero status returned, failed to run feature tests")
    }
}

// Step definitions (will fail initially)
func InitializeScenario(ctx *godog.ScenarioContext) {
    ctx.Step(`^Chris is logged in$`, chrisIsLoggedIn)
    ctx.Step(`^Chris transfers \$(\d+) from account "([^"]*)" to account "([^"]*)"$`,
        chrisTransfersMoney)
    ctx.Step(`^the transfer should succeed$`, transferShouldSucceed)
    // ... more step definitions
}
```

**Unit Test (RED):**

```go
// internal/domain/transfer_test.go
func TestTransfer_ValidateAmount(t *testing.T) {
    tests := []struct {
        name    string
        amount  decimal.Decimal
        wantErr bool
        errMsg  string
    }{
        {
            name:    "negative amount",
            amount:  decimal.NewFromFloat(-100),
            wantErr: true,
            errMsg:  "amount cannot be negative",
        },
        {
            name:    "zero amount",
            amount:  decimal.Zero,
            wantErr: true,
            errMsg:  "amount must be greater than 0",
        },
        {
            name:    "valid amount",
            amount:  decimal.NewFromFloat(100),
            wantErr: false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            transfer := &Transfer{Amount: tt.amount}
            err := transfer.Validate()

            if tt.wantErr && err == nil {
                t.Error("expected error but got nil")
            }
            if !tt.wantErr && err != nil {
                t.Errorf("unexpected error: %v", err)
            }
            if tt.wantErr && err != nil && err.Error() != tt.errMsg {
                t.Errorf("expected error %q, got %q", tt.errMsg, err.Error())
            }
        })
    }
}
```

#### 5.2 GREEN - Implement Minimum Code

**Implementation:**

```go
// internal/domain/transfer.go
type Transfer struct {
    ID            string
    FromAccountID string
    ToAccountID   string
    Amount        decimal.Decimal
    Status        TransferStatus
    CreatedAt     time.Time
}

func (t *Transfer) Validate() error {
    if t.Amount.IsNegative() {
        return errors.New("amount cannot be negative")
    }
    if t.Amount.IsZero() {
        return errors.New("amount must be greater than 0")
    }
    if t.FromAccountID == t.ToAccountID {
        return errors.New("cannot transfer to the same account")
    }
    return nil
}

// internal/service/transfer_service.go
type TransferService struct {
    accountRepo AccountRepository
    transferRepo TransferRepository
}

func (s *TransferService) ExecuteTransfer(
    ctx context.Context,
    from, to string,
    amount decimal.Decimal,
) (*Transfer, error) {
    // Create transfer object
    transfer := &Transfer{
        ID:            uuid.New().String(),
        FromAccountID: from,
        ToAccountID:   to,
        Amount:        amount,
        Status:        TransferStatusPending,
        CreatedAt:     time.Now(),
    }

    // Validate transfer
    if err := transfer.Validate(); err != nil {
        return nil, fmt.Errorf("validation failed: %w", err)
    }

    // Load accounts
    fromAccount, err := s.accountRepo.FindByID(ctx, from)
    if err != nil {
        return nil, fmt.Errorf("source account not found: %w", err)
    }

    toAccount, err := s.accountRepo.FindByID(ctx, to)
    if err != nil {
        return nil, fmt.Errorf("destination account not found: %w", err)
    }

    // Check balance
    if fromAccount.Balance.LessThan(amount) {
        return nil, errors.New("insufficient balance")
    }

    // Execute transfer
    fromAccount.Balance = fromAccount.Balance.Sub(amount)
    toAccount.Balance = toAccount.Balance.Add(amount)
    transfer.Status = TransferStatusCompleted

    // Persist changes (in transaction)
    if err := s.accountRepo.Update(ctx, fromAccount); err != nil {
        return nil, fmt.Errorf("failed to update source account: %w", err)
    }
    if err := s.accountRepo.Update(ctx, toAccount); err != nil {
        return nil, fmt.Errorf("failed to update destination account: %w", err)
    }
    if err := s.transferRepo.Save(ctx, transfer); err != nil {
        return nil, fmt.Errorf("failed to save transfer: %w", err)
    }

    return transfer, nil
}
```

#### 5.3 REFACTOR - Clean Up

```go
// Extract methods for clarity
func (s *TransferService) validateAndLoadAccounts(
    ctx context.Context,
    transfer *Transfer,
) (*Account, *Account, error) {
    if err := transfer.Validate(); err != nil {
        return nil, nil, fmt.Errorf("validation failed: %w", err)
    }

    from, err := s.accountRepo.FindByID(ctx, transfer.FromAccountID)
    if err != nil {
        return nil, nil, fmt.Errorf("source account not found: %w", err)
    }

    to, err := s.accountRepo.FindByID(ctx, transfer.ToAccountID)
    if err != nil {
        return nil, nil, fmt.Errorf("destination account not found: %w", err)
    }

    return from, to, nil
}

func (s *TransferService) checkSufficientBalance(
    account *Account,
    amount decimal.Decimal,
) error {
    if account.Balance.LessThan(amount) {
        return ErrInsufficientBalance
    }
    return nil
}
```

### Phase 6: Verification

#### 6.1 Run All Tests

```bash
# Unit tests
go test -v -race -cover ./...

# BDD scenarios
go test -v ./features/...

# Coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

#### 6.2 Verify Scenarios

All Gherkin scenarios should pass:

```
Feature: Money Transfer
  ✓ Transfer money successfully
  ✓ Insufficient balance
  ✓ Transfer amount validation
```

### Phase 7: Review & Integration

#### 7.1 Code Review

Use the code-reviewer agent:

- Security analysis
- Performance review
- Code quality check
- Best practices compliance

#### 7.2 Documentation

- Update API documentation
- Update user guides
- Document business rules
- Add inline comments for complex logic

#### 7.3 Commit Changes

```bash
# Stage changes
git add .

# Commit with clear message
git commit -m "feat: implement money transfer feature

- Add Transfer domain entity with validation
- Implement TransferService for business logic
- Add repository layer for persistence
- Implement BDD scenarios for transfer flows
- Add comprehensive unit and integration tests

Closes #123"
```

### Phase 8: Deployment

#### 8.1 Create Pull Request

- Comprehensive PR description
- Link to feature specification
- Test coverage report
- Manual testing checklist

#### 8.2 Deployment Pipeline

- Automated tests run
- Code quality checks
- Security scanning
- Staging deployment
- Production deployment

## Best Practices Summary

### DO

- ✅ Start with conversations and examples
- ✅ Write scenarios before code
- ✅ Use concrete examples with real data
- ✅ Focus on behavior, not implementation
- ✅ Keep scenarios independent
- ✅ Follow TDD: RED → GREEN → REFACTOR
- ✅ Commit working code incrementally
- ✅ Run tests frequently

### DON'T

- ❌ Skip the discovery phase
- ❌ Write implementation before scenarios
- ❌ Use technical language in scenarios
- ❌ Make scenarios dependent on each other
- ❌ Disable or skip tests
- ❌ Commit broken code
- ❌ Use `--no-verify` to bypass checks

## Tools & Commands

### Running BDD Tests

```bash
# Run all feature tests
go test ./features/...

# Run specific feature
go test ./features/... -godog.tags="@transfer"

# Run with verbose output
go test -v ./features/...
```

### Test Coverage

```bash
# Generate coverage for all tests
go test -cover ./...

# Detailed coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Code Quality

```bash
# Format code
gofmt -w .
goimports -w .

# Static analysis
go vet ./...
staticcheck ./...

# Security scan
gosec ./...
```

## Reference Materials

- [Cucumber/Gherkin Syntax](https://cucumber.io/docs/gherkin/reference/)
- [Godog BDD Framework](https://github.com/cucumber/godog)
- [BDD Best Practices](https://cucumber.io/docs/bdd/)
- [Example Mapping](https://cucumber.io/blog/bdd/example-mapping-introduction/)

## Workflow Diagram

```
1. Discovery          2. Examples         3. Formalize
   (Conversation)        (Concrete)          (Gherkin)
        ↓                    ↓                    ↓
   Requirements    →   Scenarios       →   .feature files

4. Plan              5. TDD              6. Verify
   (Stages)             (R-G-R)             (Tests)
        ↓                    ↓                    ↓
   IMPLEMENTATION_  →   Write Tests     →   All Passing
   PLAN.md              Implement
                        Refactor

7. Review            8. Deploy
   (Quality)            (Ship)
        ↓                    ↓
   Code Review      →   Production
   Documentation
```

## Next Steps

- **Chapter 2**: Writing Effective Gherkin Scenarios
- **Chapter 3**: Implementing Step Definitions in Go
- **Chapter 4**: Domain-Driven Design with BDD
- **Chapter 5**: Integration Testing Strategies
