# Query Language - Interpreter Pattern

## Problem

When building data-driven applications, users often need the ability to filter and search through collections of data using a simple, flexible language. However, directly exposing complex query parameters or requiring SQL-like knowledge can be overwhelming for end users. How can we provide a simple, domain-specific language that allows non-technical users to construct powerful queries?

## Solution / Implementation

The Interpreter pattern is applied to create a simple query language that interprets expressions like:

- `name = John AND (age > 30 OR department = Engineering)`
- `status = active AND priority > 3`

The implementation parses these expressions into an abstract syntax tree of expression objects, which can then evaluate whether a specific data item matches the criteria.

- **Python:** Uses abstract base classes to define the Expression interface and creates concrete expression classes for various operations (equals, greater than, AND, OR, etc.).

- **TypeScript:** Implements the Expression interface with concrete classes that handle different query operations, leveraging TypeScript's interface system.

- **Go:** Uses interfaces to define the Expression behavior, with struct implementations for different expression types, following Go's composition-based approach.

## Setup

Instructions assume you are in the `behavioral/interpreter/query_language` directory.

### Python

```bash
# No specific setup required, uses standard libraries (abc, unittest).
```

### TypeScript

```bash
# Install Node.js/npm if you haven't already.
cd typescript
npm install
```

### Go

```bash
# Ensure Go is installed.
# The go.mod file defines the module.
```

## How to Run

Instructions assume you are in the `behavioral/interpreter/query_language` directory.

### Python

```bash
cd python
python main.py
```

### TypeScript

```bash
cd typescript
npm start
```

### Go

```bash
cd go
go run main.go
```

## How to Test

Instructions assume you are in the `behavioral/interpreter/query_language` directory.

### Python

```bash
cd python
python -m unittest test_query_language.py
```

### TypeScript

```bash
cd typescript
npm test
```

### Go

```bash
cd go
go test -v ./...
```
