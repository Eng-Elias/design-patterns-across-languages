# HTTP Request Middleware - Chain of Responsibility Pattern

## Problem

Web applications often need to perform a series of checks or operations on incoming HTTP requests before the main request handler processes them. Examples include logging, authentication, authorization, data validation, caching checks, etc. Each check might decide whether the request processing should continue or be halted (e.g., returning a 401 Unauthorized or 403 Forbidden response). Implementing this as a monolithic block makes it hard to modify, reorder, or reuse these processing steps.

## Solution / Implementation

The Chain of Responsibility pattern provides a clean way to structure this processing flow. We define a chain of handler objects (middleware). Each handler contains logic for a specific task and holds a reference to the next handler in the chain.

When a request arrives, it's passed to the first handler. The handler performs its action and then decides whether to call the next handler or stop the chain. This allows for flexible configuration and decoupling of handlers.

- **Python:** Uses an abstract base class `Handler` with concrete middleware classes inheriting from it. Each `handle` method calls the next handler if processing should continue.
- **TypeScript:** Defines a `Middleware` interface and an abstract class `AbstractMiddleware` to manage the `next` handler link. Concrete middleware classes implement the `handle` method.
- **Go:** Uses a `Middleware` interface with `SetNext` and `Handle` methods. Concrete middleware structs implement this interface, embedding the next middleware in the chain. Errors are used to signal stopping the chain.

## Setup

Instructions assume you are in the `behavioral/chain_of_responsibility/http_request_middleware` directory.

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

Instructions assume you are in the `behavioral/chain_of_responsibility/http_request_middleware` directory.

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

Instructions assume you are in the `behavioral/chain_of_responsibility/http_request_middleware` directory.

### Python

```bash
cd python
python -m unittest test_http_request_middleware.py
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
