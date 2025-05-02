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

### Python

```bash
# No additional dependencies required
```

### TypeScript

```bash
cd typescript
npm install
```

### Go

```bash
cd go
go mod tidy
```

## How to Run

### Python

```bash
python python/main.py
```

### TypeScript

```bash
cd typescript
npx ts-node main.ts
```

### Go

```bash
cd go
go run main.go
```

## How to Test

### Python

```bash
python -m unittest python/test_query_language.py
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
