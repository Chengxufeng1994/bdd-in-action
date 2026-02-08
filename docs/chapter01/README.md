# Chapter 1: Getting Started with BDD

## Overview

This chapter introduces the fundamental workflow for developing features using Behavior-Driven Development (BDD) practices. You'll learn the complete flow from initial feature request through deployment.

## Contents

1. **[Feature Development Flow](01-feature-development-flow.md)**
   - Complete 8-phase workflow
   - Detailed explanations and examples
   - Best practices and anti-patterns
   - Code samples in Go

2. **[Quick Reference](02-quick-reference.md)**
   - Checklists for each phase
   - Templates for common tasks
   - Command reference
   - Decision trees

3. **[BDD Core Workflow](03-bdd-core-workflow.md)** ⭐ NEW
   - From business conversation to executable specifications
   - Three Amigos collaboration
   - Concrete examples and counterexamples
   - Automated acceptance tests as living documentation
   - Complete workflow with real-world examples

## Learning Path

### New to BDD?

**Start here:** [BDD Core Workflow](03-bdd-core-workflow.md) ⭐

This document provides a deep dive into the BDD process:
- How to conduct effective Three Amigos meetings
- Transforming abstract requirements into concrete examples
- Writing executable Gherkin specifications
- Building automated tests as living documentation
- Complete real-world examples with code

Then continue with [Feature Development Flow](01-feature-development-flow.md) to understand:
- The complete 8-phase implementation workflow
- How BDD integrates with your development process
- Best practices and anti-patterns
- Code samples in Go

### Already Familiar with BDD?

Use the [Quick Reference](02-quick-reference.md) guide for:
- Phase checklists
- Template code
- Common commands
- Quick decision trees

### Want to Understand the Core Process?

Refer to [BDD Core Workflow](03-bdd-core-workflow.md) for:
- In-depth explanation of each BDD stage
- Example Mapping techniques
- Gherkin best practices
- Step definition implementation patterns
- Living documentation strategies

## Key Concepts

### The 8-Phase Flow

```
1. Discovery          → Understand the requirement
2. Examples           → Create concrete scenarios
3. Formalize          → Write Gherkin specifications
4. Plan               → Break down implementation
5. TDD                → Test-Driven Development
6. Verify             → Confirm everything works
7. Review             → Quality assurance
8. Deploy             → Ship to production
```

### The BDD Triangle

```
        Business
        (Examples)
           ▲
          / \
         /   \
        /     \
       /       \
      /         \
Development ◄──► Testing
  (Code)        (Scenarios)
```

BDD brings together:

- **Business stakeholders**: Define what should be built through examples
- **Developers**: Build the functionality
- **Testers**: Verify behavior through automated scenarios

### Core Principles

1. **Shared Understanding**: Everyone understands what's being built
2. **Living Documentation**: Scenarios document the system behavior
3. **Outside-In**: Start from user behavior, not technical implementation
4. **Executable Specifications**: Scenarios are automated tests

## Example: Money Transfer Feature

Throughout this chapter, we use a money transfer feature as a running example:

```gherkin
Feature: Money Transfer
  As a bank customer
  I want to transfer money between accounts
  So that I can send money to others

  Scenario: Successful transfer
    Given Chris has $100 in his account
    And Lisa has $50 in her account
    When Chris transfers $30 to Lisa
    Then Chris should have $70
    And Lisa should have $80
```

This example demonstrates:

- **User story format**: As a... I want... So that...
- **Concrete examples**: Real names and amounts
- **Business language**: No technical jargon
- **Verifiable outcomes**: Clear assertions

## Prerequisites

### Required Knowledge

- Basic Go programming
- Understanding of software testing
- Familiarity with Git

### Required Tools

```bash
# Go
go version  # Should be 1.21 or later

# BDD Framework
go get github.com/cucumber/godog/cmd/godog

# Testing
go get github.com/stretchr/testify

# Code Quality
go install golang.org/x/tools/cmd/goimports@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
go install github.com/securego/gosec/v2/cmd/gosec@latest
```

## Getting Started

### 1. Set Up Your Project

```bash
# Create project structure
mkdir -p features/steps
mkdir -p internal/{domain,service,repository}
mkdir -p cmd/app

# Initialize Go module
go mod init github.com/yourusername/your-project

# Install dependencies
go get github.com/cucumber/godog
```

### 2. Create Your First Feature

Create `features/greeting.feature`:

```gherkin
Feature: Greeting
  As a user
  I want to be greeted by name
  So that I feel welcomed

  Scenario: Greet a user
    Given my name is "Chris"
    When I request a greeting
    Then I should see "Hello, Chris!"
```

### 3. Implement Step Definitions

Create `features/steps/greeting_test.go`:

```go
package steps

import (
    "testing"
    "github.com/cucumber/godog"
)

func TestFeatures(t *testing.T) {
    suite := godog.TestSuite{
        ScenarioInitializer: InitializeScenario,
        Options: &godog.Options{
            Format:   "pretty",
            Paths:    []string{"../"},
            TestingT: t,
        },
    }

    if suite.Run() != 0 {
        t.Fatal("non-zero status returned")
    }
}

func InitializeScenario(ctx *godog.ScenarioContext) {
    // Step definitions will go here
}
```

### 4. Run Your First Test

```bash
go test ./features/...
```

You'll see pending steps - that's perfect! Now you're ready to implement them.

## Common Patterns

### Feature File Structure

```gherkin
Feature: [Name]
  [User Story]

  Background:
    [Common setup for all scenarios]

  Scenario: [Happy path]
    Given [context]
    When [action]
    Then [outcome]

  Scenario: [Error case]
    ...

  Scenario Outline: [Data-driven]
    ...
    Examples:
      ...
```

### Step Definition Pattern

```go
func InitializeScenario(ctx *godog.ScenarioContext) {
    // Setup/teardown
    ctx.Before(beforeScenario)
    ctx.After(afterScenario)

    // Given steps (arrange)
    ctx.Step(`pattern`, givenFunction)

    // When steps (act)
    ctx.Step(`pattern`, whenFunction)

    // Then steps (assert)
    ctx.Step(`pattern`, thenFunction)
}
```

### TDD Cycle

```go
// 1. RED - Write failing test
func TestTransfer_InsufficientBalance(t *testing.T) {
    // This will fail initially
    err := transfer.Execute()
    assert.Error(t, err)
}

// 2. GREEN - Make it pass
func (t *Transfer) Execute() error {
    if t.balance < t.amount {
        return errors.New("insufficient balance")
    }
    return nil
}

// 3. REFACTOR - Improve
func (t *Transfer) Execute() error {
    if err := t.validateBalance(); err != nil {
        return fmt.Errorf("transfer failed: %w", err)
    }
    return nil
}
```

## Troubleshooting

### Scenarios Not Running

```bash
# Check feature file syntax
godog --dry-run features/

# Verify file paths
go test -v ./features/...
```

### Step Definitions Not Matching

```bash
# Run with verbose output to see patterns
go test -v ./features/...

# Check regex patterns carefully
ctx.Step(`^exact pattern here$`, function)
```

### Tests Failing

```bash
# Run with race detection
go test -race ./...

# Check for state leaking between tests
# Ensure proper cleanup in ctx.After()
```

## Next Steps

After completing Chapter 1, you should be able to:

- ✅ Understand the 8-phase BDD flow
- ✅ Write basic Gherkin scenarios
- ✅ Implement step definitions in Go
- ✅ Follow the TDD cycle (RED-GREEN-REFACTOR)
- ✅ Run BDD tests

Continue to:

- **Chapter 2**: Writing Effective Gherkin Scenarios
- **Chapter 3**: Advanced Step Definition Patterns
- **Chapter 4**: Domain-Driven Design with BDD
- **Chapter 5**: Integration Testing Strategies

## Additional Resources

### Official Documentation

- [Cucumber/Gherkin Docs](https://cucumber.io/docs/gherkin/)
- [Godog Framework](https://github.com/cucumber/godog)
- [BDD Guide](https://cucumber.io/docs/bdd/)

### Books

- "BDD in Action" by John Ferguson Smart
- "The Cucumber Book" by Matt Wynne & Aslak Hellesøy
- "Specification by Example" by Gojko Adzic

### Community

- [Cucumber Community](https://cucumber.io/community)
- [BDD Slack Channel](https://cucumber.io/support#slack)

## Feedback

Found an issue or have a suggestion? Please open an issue in the repository.

---

**Ready to start?** Open [Feature Development Flow](01-feature-development-flow.md) and begin your BDD journey!
