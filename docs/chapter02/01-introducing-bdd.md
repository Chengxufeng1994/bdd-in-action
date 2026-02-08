# Chapter 2: Introducing BDD

> Source: [Introducing BDD by Dan North](https://dannorth.net/blog/introducing-bdd/)

## Overview

Behaviour-Driven Development (BDD) emerged from Dan North's observations about widespread confusion in Test-Driven Development (TDD) practices. Rather than focusing on "testing," BDD reframes software development around describing system behavior.

## Core Concepts

### The Problem with TDD

Despite the established practice of Test-Driven Development, many teams struggled with fundamental questions:

- Where to start?
- What to test and what not to test?
- How much to test in one go?
- What to call the tests?
- How to understand why a test fails?

These recurring questions suggested that something was missing from the TDD narrative.

## Key Principles

### 1. Test Method Naming as Sentences

The foundational insight came from the **agiledox** tool, which converts camelCase method names into readable sentences. This enables test methods to serve as documentation that resonates with both technical and business stakeholders.

**Example:**

```
testIsValidSubscriber  →  Test is valid subscriber
```

This simple transformation makes test names readable and meaningful to non-technical stakeholders.

### 2. The "Should" Template

Using the sentence structure **"The class _should_ do something"** keeps developers focused on individual responsibilities.

**Benefits:**

- Keeps focus on behavior, not implementation
- Encourages single responsibility principle
- Makes design issues obvious

**Example:**

```
CustomerValidator should check that customer has valid subscription
CustomerValidator should reject customer with expired subscription
```

When a method name doesn't fit this pattern, it signals that behavior belongs in a different class—encouraging better design through single responsibility principles.

### 3. Behaviour Over Testing Terminology

The linguistic shift from "test" to "behaviour" dissolved numerous coaching questions:

| Question                   | BDD Answer                               |
| -------------------------- | ---------------------------------------- |
| What to test?              | Describe the next behavior               |
| How much to test?          | One sentence worth of behavior           |
| What to call the test?     | A sentence describing the behavior       |
| How to interpret failures? | Bug, moved behavior, or outdated premise |

## Acceptance Criteria Framework

BDD extends beyond unit testing to requirements definition using the **Given-When-Then** template:

### The Template Structure

```gherkin
Given [initial context]
When [event occurs]
Then [ensure outcomes]
```

### ATM Withdrawal Example

**Scenario 1: Account is in credit**

```gherkin
Given the account is in credit
  And the card is valid
  And the dispenser contains cash
When the customer requests cash
Then ensure the account is debited
  And ensure cash is dispensed
  And ensure the card is returned
```

**Scenario 2: Account overdrawn past limit**

```gherkin
Given the account is overdrawn
  And the card is valid
When the customer requests cash
Then ensure a rejection message is displayed
  And ensure cash is not dispensed
  And ensure the card is returned
```

### Benefits of Given-When-Then

1. **Reusable scenario fragments**: Each clause can be mapped to executable code
2. **Clear structure**: Separates setup, action, and verification
3. **Business-readable**: Stakeholders can validate scenarios
4. **Direct mapping to code**: Can be automated using frameworks like JBehave

## Evolution and Impact

### JBehave Framework

JBehave was created as a framework that explicitly avoids testing vocabulary. It provides:

- Support for Given-When-Then scenarios
- Plain text story files
- Direct mapping to Java classes
- Business-readable test specifications

### Framework Influence

The BDD approach has influenced numerous frameworks:

- **RSpec** (Ruby): Pioneered spec-style syntax
- **Cucumber**: Cross-platform Given-When-Then framework
- **SpecFlow** (.NET): BDD framework for .NET
- **Jasmine/Jest** (JavaScript): Behavior-focused testing

### Real-World Adoption

BDD has been adopted across agile teams, establishing a **ubiquitous language** that connects:

- Business analysts
- Developers
- Testers
- Business stakeholders

## Key Takeaways

1. **Language matters**: Shifting from "test" to "behavior" clarifies intent
2. **Readable names**: Test/spec names should be complete sentences
3. **Focus on behavior**: Describe what the system should do, not how it does it
4. **Business alignment**: Use Given-When-Then to create shared understanding
5. **Executable specifications**: Scenarios should be both readable and runnable

## Next Steps

To implement BDD effectively:

1. Start with behavior descriptions, not tests
2. Use the "should" template for unit-level behaviors
3. Use Given-When-Then for acceptance criteria
4. Create a ubiquitous language with your team
5. Make specifications executable with appropriate frameworks

---

## References

- [Original Article: Introducing BDD](https://dannorth.net/blog/introducing-bdd/)
- Dan North's Blog: <https://dannorth.net>
- JBehave Framework: <https://jbehave.org>
