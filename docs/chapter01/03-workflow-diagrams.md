# BDD Workflow Diagrams

## Complete Feature Development Flow

```
┌─────────────────────────────────────────────────────────────────┐
│                    PHASE 1: DISCOVERY                           │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  Business Need  →  User Story  →  Three Amigos  →  Examples    │
│                                                                  │
│  Output: Shared understanding of feature requirements           │
└─────────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────────┐
│                    PHASE 2: EXAMPLES                            │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  Concrete     →  Happy Path    →  Error Cases  →  Edge Cases   │
│  Examples        Scenarios        Scenarios       Scenarios     │
│                                                                  │
│  Output: Multiple concrete scenarios with real data             │
└─────────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────────┐
│                    PHASE 3: FORMALIZE                           │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  Create .feature  →  Write Gherkin  →  Review  →  Refine       │
│  File                Scenarios          Syntax     Scenarios    │
│                                                                  │
│  Output: Executable specification in .feature files             │
└─────────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────────┐
│                    PHASE 4: PLAN                                │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  Break Down  →  Define Stages  →  Success     →  Identify      │
│  Work           (3-5 stages)      Criteria        Tests         │
│                                                                  │
│  Output: IMPLEMENTATION_PLAN.md with clear stages               │
└─────────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────────┐
│                    PHASE 5: TDD                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  ┌────────────────────────────────────┐                        │
│  │  RED → GREEN → REFACTOR → COMMIT   │  ← Repeat for          │
│  │   ↑                         ↓      │     each test          │
│  │   └─────────────────────────┘      │                        │
│  └────────────────────────────────────┘                        │
│                                                                  │
│  Output: Tested, working code with all tests passing            │
└─────────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────────┐
│                    PHASE 6: VERIFY                              │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  Unit Tests  →  BDD Scenarios  →  Coverage  →  All Green?      │
│  Pass          Pass               Report       Yes/No           │
│                                                                  │
│  Output: Verified feature with >80% coverage                    │
└─────────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────────┐
│                    PHASE 7: REVIEW                              │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  Code      →  Security   →  Quality   →  Documentation         │
│  Review       Review        Check        Update                 │
│                                                                  │
│  Output: Production-ready code meeting quality standards        │
└─────────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────────┐
│                    PHASE 8: DEPLOY                              │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  Commit  →  Pull Request  →  CI/CD  →  Staging  →  Production │
│                                                                  │
│  Output: Feature live in production                             │
└─────────────────────────────────────────────────────────────────┘
```

## TDD Cycle (Detailed)

```
                    Start New Feature
                           │
                           ↓
┌──────────────────────────────────────────────────────────────┐
│                        RED PHASE                              │
├──────────────────────────────────────────────────────────────┤
│                                                               │
│  1. Write a failing test                                     │
│     • Test should fail for the RIGHT reason                  │
│     • Test should be specific and focused                    │
│     • Use table-driven tests in Go                           │
│                                                               │
│  Example:                                                     │
│    func TestTransfer_InsufficientBalance(t *testing.T) {    │
│        err := transfer.Execute()                             │
│        assert.Error(t, err) // FAILS - not implemented       │
│    }                                                          │
│                                                               │
└──────────────────────────────────────────────────────────────┘
                           │
                           ↓
                     Run Tests
                           │
                           ↓
                    Tests Fail? ────No──→ Test is invalid!
                           │              Write better test
                          Yes
                           ↓
┌──────────────────────────────────────────────────────────────┐
│                       GREEN PHASE                             │
├──────────────────────────────────────────────────────────────┤
│                                                               │
│  2. Write MINIMUM code to make test pass                     │
│     • Don't overthink it                                     │
│     • Simplest solution that works                           │
│     • No extra features                                      │
│                                                               │
│  Example:                                                     │
│    func (t *Transfer) Execute() error {                      │
│        if t.balance < t.amount {                             │
│            return errors.New("insufficient balance")         │
│        }                                                      │
│        return nil                                            │
│    }                                                          │
│                                                               │
└──────────────────────────────────────────────────────────────┘
                           │
                           ↓
                     Run Tests
                           │
                           ↓
                    Tests Pass? ────No──→ Fix implementation
                           │              Don't change test!
                          Yes
                           ↓
┌──────────────────────────────────────────────────────────────┐
│                     REFACTOR PHASE                            │
├──────────────────────────────────────────────────────────────┤
│                                                               │
│  3. Clean up code (while tests stay green)                   │
│     • Extract methods                                        │
│     • Rename for clarity                                     │
│     • Remove duplication                                     │
│     • Improve structure                                      │
│                                                               │
│  Example:                                                     │
│    func (t *Transfer) Execute() error {                      │
│        if err := t.validateBalance(); err != nil {           │
│            return fmt.Errorf("transfer failed: %w", err)     │
│        }                                                      │
│        return t.processTransfer()                            │
│    }                                                          │
│                                                               │
└──────────────────────────────────────────────────────────────┘
                           │
                           ↓
                     Run Tests
                           │
                           ↓
                    Still Pass? ────No──→ Undo refactoring
                           │              Tests must stay green!
                          Yes
                           ↓
┌──────────────────────────────────────────────────────────────┐
│                      COMMIT PHASE                             │
├──────────────────────────────────────────────────────────────┤
│                                                               │
│  4. Commit working code                                      │
│     • All tests passing                                      │
│     • Code formatted                                         │
│     • Clear commit message                                   │
│                                                               │
│  git add .                                                    │
│  git commit -m "feat: add balance validation"                │
│                                                               │
└──────────────────────────────────────────────────────────────┘
                           │
                           ↓
                   More tests needed? ──Yes──→ Return to RED
                           │
                          No
                           ↓
                    Feature Complete
```

## Three Amigos Collaboration

```
┌─────────────────┐         ┌─────────────────┐         ┌─────────────────┐
│   BUSINESS      │         │   DEVELOPMENT   │         │    TESTING      │
│   ANALYST       │         │   ENGINEER      │         │    ENGINEER     │
│                 │         │                 │         │                 │
│  "What do we    │         │  "How can we    │         │  "What could    │
│   need to       │         │   build this?"  │         │   go wrong?"    │
│   build?"       │         │                 │         │                 │
└────────┬────────┘         └────────┬────────┘         └────────┬────────┘
         │                           │                           │
         │                           │                           │
         │         ┌─────────────────▼─────────────────┐        │
         └────────►│    SHARED UNDERSTANDING          │◄────────┘
                   │                                   │
                   │  • Clear requirements             │
                   │  • Concrete examples              │
                   │  • Edge cases identified          │
                   │  • Acceptance criteria defined    │
                   │                                   │
                   └──────────────┬────────────────────┘
                                  │
                                  ↓
                   ┌──────────────────────────────────┐
                   │     EXECUTABLE SPECIFICATION     │
                   │        (.feature files)          │
                   └──────────────────────────────────┘
```

## BDD Testing Pyramid

```
                          ┌───────────┐
                         /             \
                        /    E2E Tests  \          ← Few
                       /    (Scenarios)  \           Slow
                      /___________________\          Brittle
                     /                     \
                    /   Integration Tests   \      ← Some
                   /    (Step Definitions)   \       Medium
                  /___________________________ \      Stable
                 /                             \
                /        Unit Tests             \   ← Many
               /      (Domain Logic TDD)         \    Fast
              /_____________________________________\  Reliable

    Unit Tests (70%)       → Test individual functions/methods
    Integration Tests (20%) → Test components working together
    E2E Tests (10%)        → Test complete user journeys
```

## Scenario State Flow

```
┌──────────────────────────────────────────────────────────────┐
│                    SCENARIO EXECUTION                         │
└──────────────────────────────────────────────────────────────┘

    Background (Shared Setup)
           │
           ↓
    ┌─────────────┐
    │   GIVEN     │  ← Arrange (Set up state)
    │   (Setup)   │
    └──────┬──────┘
           │
           ↓
    ┌─────────────┐
    │   WHEN      │  ← Act (Perform action)
    │   (Action)  │
    └──────┬──────┘
           │
           ↓
    ┌─────────────┐
    │   THEN      │  ← Assert (Verify outcome)
    │   (Assert)  │
    └──────┬──────┘
           │
           ↓
    ┌─────────────┐
    │   AND       │  ← Additional verifications
    │   (More     │
    │   Asserts)  │
    └──────┬──────┘
           │
           ↓
    Scenario Complete
           │
           ↓
    Clean Up (ctx.After)
```

## Feature Lifecycle

```
1. IDEA PHASE
   ├─ Business need identified
   ├─ Stakeholder discussion
   └─ User story created
          │
          ↓
2. DISCOVERY PHASE
   ├─ Three Amigos meeting
   ├─ Example mapping
   ├─ Edge cases explored
   └─ Concrete examples written
          │
          ↓
3. SPECIFICATION PHASE
   ├─ .feature file created
   ├─ Gherkin scenarios written
   ├─ Scenarios reviewed
   └─ Acceptance criteria formalized
          │
          ↓
4. DEVELOPMENT PHASE
   ├─ Implementation plan created
   ├─ Step definitions written (RED)
   ├─ Domain logic implemented (GREEN)
   ├─ Code refactored (REFACTOR)
   └─ Tests passing
          │
          ↓
5. VERIFICATION PHASE
   ├─ Unit tests passing
   ├─ BDD scenarios passing
   ├─ Coverage verified
   └─ Quality checks passed
          │
          ↓
6. DELIVERY PHASE
   ├─ Code reviewed
   ├─ Pull request created
   ├─ CI/CD pipeline passes
   ├─ Deployed to staging
   └─ Deployed to production
          │
          ↓
7. MAINTENANCE PHASE
   ├─ Scenarios become documentation
   ├─ Tests run on every change
   ├─ Living specification maintained
   └─ Regression protection
```

## Decision Flow: When to Write Scenarios

```
                    New Requirement
                           │
                           ↓
              Is behavior user-facing? ────No──→ Write unit tests only
                           │
                          Yes
                           ↓
              Is it a critical flow? ───────No──→ Consider: maybe unit tests
                           │                      + integration tests are enough
                          Yes
                           ↓
              Do stakeholders need ─────No──→ Write for developers
              to understand it?                as living documentation
                           │
                          Yes
                           ↓
              Write BDD scenarios
                           │
                           ↓
              ┌────────────┴─────────────┐
              │                          │
      Happy Path Scenarios      Error Scenarios
              │                          │
              ↓                          ↓
      Normal user flows          What can go wrong?
      Expected outcomes          Edge cases
                           │
                           ↓
              Scenario Outlines for data variations
```

## Error Recovery Flow

```
    Implementation Attempt
           │
           ↓
    ┌─────────────┐
    │  Attempt 1  │ ──Failed──→ Document what failed
    └─────────────┘             Review error messages
           │
      Succeeded
           │
           ↓
    ┌─────────────┐
    │  Attempt 2  │ ──Failed──→ Research alternatives
    └─────────────┘             Find similar implementations
           │
      Succeeded
           │
           ↓
    ┌─────────────┐
    │  Attempt 3  │ ──Failed──→ STOP!
    └─────────────┘             │
           │                     ↓
      Succeeded            Question fundamentals:
           │               • Wrong abstraction level?
           ↓               • Can this be split?
      Continue            • Is there simpler approach?
                                │
                                ↓
                          Try completely different angle
                                │
                                ↓
                          Ask for help if still stuck
```

## Code Quality Gates

```
┌──────────────────────────────────────────────────────────────┐
│                    BEFORE COMMIT                              │
├──────────────────────────────────────────────────────────────┤
│                                                               │
│  ✓ All tests pass              go test ./...                 │
│  ✓ No race conditions          go test -race ./...           │
│  ✓ Code formatted              gofmt -w .                    │
│  ✓ Imports organized           goimports -w .                │
│  ✓ No vet warnings             go vet ./...                  │
│  ✓ No staticcheck issues       staticcheck ./...             │
│  ✓ Coverage maintained         go test -cover ./...          │
│                                                               │
└──────────────────────────────────────────────────────────────┘
                           │
                           ↓
┌──────────────────────────────────────────────────────────────┐
│                  BEFORE PULL REQUEST                          │
├──────────────────────────────────────────────────────────────┤
│                                                               │
│  ✓ All BDD scenarios pass      go test ./features/...        │
│  ✓ No security issues          gosec ./...                   │
│  ✓ Code reviewed               Use code-reviewer agent       │
│  ✓ Documentation updated       Update README, docs           │
│  ✓ Commit message clear        Conventional format           │
│                                                               │
└──────────────────────────────────────────────────────────────┘
                           │
                           ↓
┌──────────────────────────────────────────────────────────────┐
│                    BEFORE DEPLOYMENT                          │
├──────────────────────────────────────────────────────────────┤
│                                                               │
│  ✓ PR approved                 Team review complete          │
│  ✓ CI/CD pipeline green        All automated checks pass     │
│  ✓ Staging tested              Manual verification done      │
│  ✓ Rollback plan ready         Know how to revert            │
│  ✓ Monitoring configured       Can detect issues             │
│                                                               │
└──────────────────────────────────────────────────────────────┘
```

## Parallel Development Flow

```
┌─────────────────────────────────────────────────────────────┐
│                  MULTIPLE FEATURES IN PARALLEL               │
└─────────────────────────────────────────────────────────────┘

    Main Branch
         │
         ├──→ Feature A Branch
         │         │
         │         ├─ Write scenarios
         │         ├─ Implement (TDD)
         │         ├─ Tests pass
         │         └─ PR → Merge
         │              │
         ├──→ Feature B Branch
         │         │
         │         ├─ Write scenarios
         │         ├─ Implement (TDD)
         │         ├─ Tests pass
         │         └─ PR → Merge
         │              │
         └──→ Feature C Branch
                   │
                   ├─ Write scenarios
                   ├─ Implement (TDD)
                   ├─ Tests pass
                   └─ PR → Merge
                        │
                        ↓
                  All Features
                   Integrated
                        │
                        ↓
                   Production
```

## Summary

These diagrams illustrate:

1. **Linear Flow**: The step-by-step progression from idea to production
2. **Cyclic Flow**: The TDD cycle that repeats for each test case
3. **Collaborative Flow**: How different roles work together
4. **Hierarchical Flow**: The testing pyramid showing test distribution
5. **State Flow**: How scenarios execute through Given-When-Then
6. **Lifecycle Flow**: The complete journey of a feature
7. **Decision Flow**: When to use different testing approaches
8. **Recovery Flow**: How to handle implementation challenges
9. **Quality Gates**: Checkpoints before advancing
10. **Parallel Flow**: Managing multiple features simultaneously

Use these diagrams as visual guides while following the detailed instructions in the main documentation.
