---
name: golang-code-review
description: Reviews Go (Golang) code and provides feedback on correctness, readability, performance, and best practices.
argument-hint: go_code_or_diff
---

# Skill: Golang Code Review

## Purpose
This skill analyzes Go (Golang) source code and generates a structured code review.  
It focuses on identifying bugs, improving readability, ensuring idiomatic Go practices, and suggesting performance or maintainability improvements.

## When to Use
Use this skill when:
- The user provides Go source code.
- The user asks for a **code review**, **feedback**, or **improvements** for Go code.
- The user provides a **git diff**, **pull request**, or **code snippet** written in Go.

Do NOT use this skill when:
- The code is not written in Go.
- The user only asks for explanation rather than review.
- The request is about project setup or architecture without code.

## Input
Expected input:

- A Go source file
- A code snippet
- A git diff or pull request patch

Example input:

```go
func Add(a int, b int) int {
    result := a + b
    return result
}