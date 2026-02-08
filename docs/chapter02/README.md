# Chapter 2: Introducing Behaviour-Driven Development

## Overview

This chapter explores the origins and core concepts of Behaviour-Driven Development (BDD) as introduced by Dan North. It covers the evolution from Test-Driven Development (TDD) to BDD and explains why the shift in terminology and approach matters.

## Contents

### [01. Introducing BDD](./01-introducing-bdd.md)
- The problems with TDD that led to BDD
- Key principles: "should" template, behavior over testing
- Given-When-Then acceptance criteria framework
- Evolution and framework influence

### [02. BDD Principles and Practices - From Business Goals to Code](./02-bdd-principles-and-practices.md)
- The 7-level BDD hierarchy
- Business Goals → Features → Examples → Executable → Specifications → Low-Level Specs → Application Code
- Complete money transfer example across all levels
- Outside-in development approach
- Traceability from business value to implementation
- Practical Go code examples with godog

## Key Concepts

### What is BDD?

BDD is a software development approach that:
- Focuses on **behavior** rather than testing
- Uses **ubiquitous language** understood by all stakeholders
- Creates **executable specifications** using Given-When-Then
- Encourages **outside-in** development from business value

### Why BDD Matters

1. **Clarity**: Shifts focus from "how to test" to "what should the system do"
2. **Communication**: Creates shared understanding between technical and business teams
3. **Documentation**: Specifications serve as living documentation
4. **Design**: Encourages better design through focus on behavior

### Core Practices

1. **Unit Level**: Use "should" sentences to describe behavior
2. **Acceptance Level**: Use Given-When-Then scenarios
3. **Collaboration**: Write scenarios with business stakeholders
4. **Automation**: Make scenarios executable

## Questions Answered

- **Where to start?** → Start with the behavior that delivers the most value
- **What to test?** → Describe the next most important behavior
- **How much to test?** → One behavior (one sentence) at a time
- **What to call the test?** → A sentence describing what the class/system should do
- **Why did a test fail?** → Either a bug, behavior moved elsewhere, or premise changed

## Progression

This chapter sets the foundation for:
- Understanding BDD principles and practices
- Writing effective Given-When-Then scenarios
- Implementing BDD with your team
- Choosing appropriate BDD frameworks

## Further Reading

- [Original Article by Dan North](https://dannorth.net/blog/introducing-bdd/)
- Chapter 3: BDD Practices in Depth (coming next)
- Appendix: BDD Framework Comparison
