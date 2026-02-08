# Quick Reference: BDD Feature Development

## The 8-Phase Flow

```
Discovery → Examples → Formalize → Plan → TDD → Verify → Review → Deploy
```

## Phase Checklist

### 1. Discovery

- [ ] Identify business need
- [ ] Write user story (As a... I want... So that...)
- [ ] Conduct Three Amigos meeting
- [ ] Define acceptance criteria

### 2. Examples

- [ ] Write concrete examples with real data
- [ ] Cover happy path
- [ ] Identify alternative paths
- [ ] Document error scenarios
- [ ] Explore edge cases

### 3. Formalize

- [ ] Create `.feature` file
- [ ] Write scenarios in Gherkin
- [ ] Use Background for common setup
- [ ] Use Scenario Outline for similar cases
- [ ] Review and refine

### 4. Plan

- [ ] Create `IMPLEMENTATION_PLAN.md`
- [ ] Break into 3-5 stages
- [ ] Define success criteria for each stage
- [ ] List specific tests
- [ ] Identify dependencies

### 5. TDD Cycle

- [ ] **RED**: Write failing test
- [ ] **GREEN**: Implement minimum code to pass
- [ ] **REFACTOR**: Clean up while keeping tests green
- [ ] Commit after each complete cycle
- [ ] Repeat for each test case

### 6. Verify

- [ ] Run unit tests: `go test -v -race -cover ./...`
- [ ] Run BDD scenarios: `go test -v ./features/...`
- [ ] Check coverage: `go test -coverprofile=coverage.out ./...`
- [ ] Verify all scenarios pass
- [ ] Confirm 80%+ code coverage

### 7. Review

- [ ] Run code-reviewer agent
- [ ] Address CRITICAL and HIGH issues
- [ ] Fix security vulnerabilities
- [ ] Update documentation
- [ ] Verify code formatting: `gofmt -w .`
- [ ] Run static analysis: `go vet ./...`

### 8. Deploy

- [ ] Create detailed commit message
- [ ] Create pull request
- [ ] Link to feature specification
- [ ] Include test coverage report
- [ ] Wait for CI/CD pipeline
- [ ] Deploy to staging
- [ ] Deploy to production

## Gherkin Template

```gherkin
Feature: [Feature Name]
  As a [role]
  I want [feature]
  So that [benefit]

  Background:
    Given [common setup]

  Scenario: [Happy path scenario name]
    Given [initial context]
    When [action]
    Then [expected outcome]
    And [additional verification]

  Scenario: [Error scenario name]
    Given [initial context]
    When [error-inducing action]
    Then [error outcome]
    And [state unchanged]

  Scenario Outline: [Parameterized scenario name]
    Given [context]
    When [action with <parameter>]
    Then [outcome with <parameter>]

    Examples:
      | parameter | expected |
      | value1    | result1  |
      | value2    | result2  |
```

## Go Test Template

```go
// Table-driven unit test
func TestFeature_Behavior(t *testing.T) {
    tests := []struct {
        name    string
        input   InputType
        want    ExpectedType
        wantErr bool
    }{
        {
            name:    "scenario description",
            input:   /* test input */,
            want:    /* expected result */,
            wantErr: false,
        },
        // More test cases...
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := FunctionUnderTest(tt.input)

            if (err != nil) != tt.wantErr {
                t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("got %v, want %v", got, tt.want)
            }
        })
    }
}
```

## Step Definition Template

```go
func InitializeScenario(ctx *godog.ScenarioContext) {
    // Given steps (setup)
    ctx.Step(`^user has \$(\d+) in account$`, userHasBalance)

    // When steps (actions)
    ctx.Step(`^user transfers \$(\d+) to account "([^"]*)"$`,
        userTransfersMoney)

    // Then steps (assertions)
    ctx.Step(`^the transfer should (succeed|fail)$`,
        transferOutcome)
    ctx.Step(`^user should see (message|error) "([^"]*)"$`,
        userSeesMessage)
}
```

## Common Commands

```bash
# Development
go test ./...                    # Run all tests
go test -v ./...                 # Verbose output
go test -race ./...              # Race detection
go test -cover ./...             # Coverage

# BDD Specific
go test ./features/...           # Run feature tests
go test -godog.tags="@wip"       # Run tagged scenarios

# Code Quality
gofmt -w .                       # Format code
goimports -w .                   # Organize imports
go vet ./...                     # Static analysis
staticcheck ./...                # Extended checks
gosec ./...                      # Security scan

# Coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Implementation Plan Template

```markdown
# [Feature Name] Implementation

## Stage 1: [Stage Name]

**Goal**: [Specific deliverable]
**Success Criteria**: [Testable outcomes]
**Tests**: [Specific test cases]
**Status**: Not Started | In Progress | Complete

## Stage 2: [Next Stage]

...
```

## Git Commit Message Format

```
<type>: <short description>

<detailed description>

<breaking changes>
<ticket references>
```

**Types**: feat, fix, refactor, docs, test, chore, perf, ci

## BDD Anti-Patterns to Avoid

❌ **Implementation details in scenarios**

```gherkin
# BAD
When the TransferService.execute() method is called

# GOOD
When Chris transfers $100 to Lisa
```

❌ **Technical language instead of business language**

```gherkin
# BAD
Then the database should have 2 records

# GOOD
Then both accounts should reflect the transfer
```

❌ **Testing through the UI**

```gherkin
# BAD
When I click the "Transfer" button
And I fill in "amount" with "100"

# GOOD
When Chris transfers $100 to Lisa
```

❌ **Dependent scenarios**

```gherkin
# BAD - depends on previous scenario
Scenario: Create account
  Given...

Scenario: Transfer money  # Assumes account exists
  When...

# GOOD - independent
Scenario: Transfer money
  Given account exists  # Explicit setup
  When...
```

## Decision Tree

```
Need to add feature?
  ↓
Do I understand the requirement?
  No  → Conduct Discovery meeting
  Yes ↓

Do I have concrete examples?
  No  → Write example scenarios
  Yes ↓

Are examples formalized?
  No  → Write .feature file
  Yes ↓

Do I have implementation plan?
  No  → Create IMPLEMENTATION_PLAN.md
  Yes ↓

Have tests been written?
  No  → Write failing tests (RED)
  Yes ↓

Do tests pass?
  No  → Implement code (GREEN)
  Yes ↓

Can code be improved?
  Yes → Refactor and repeat
  No  ↓

All scenarios passing?
  No  → Continue TDD cycle
  Yes ↓

Code reviewed?
  No  → Run code-reviewer agent
  Yes ↓

Ready to deploy? → Create PR and deploy
```

## When to Stop and Ask

Stop after 3 failed attempts if:

- Tests keep failing
- Unclear how to implement
- Architecture feels wrong
- Too much complexity

Then:

1. Document what failed
2. Research alternatives
3. Question fundamentals
4. Try different approach

## Quality Gates

Before committing:

- [ ] All tests pass
- [ ] Code formatted
- [ ] No linter warnings
- [ ] No security issues
- [ ] Coverage maintained
- [ ] Documentation updated

Before creating PR:

- [ ] All scenarios pass
- [ ] Code reviewed
- [ ] Integration tests pass
- [ ] Manual testing complete

Before deploying:

- [ ] PR approved
- [ ] CI/CD pipeline green
- [ ] Staging tested
- [ ] Rollback plan ready

## Resources

- **Gherkin Reference**: https://cucumber.io/docs/gherkin/reference/
- **Godog**: https://github.com/cucumber/godog
- **Example Mapping**: https://cucumber.io/blog/bdd/example-mapping-introduction/
- **BDD Best Practices**: https://cucumber.io/docs/bdd/
