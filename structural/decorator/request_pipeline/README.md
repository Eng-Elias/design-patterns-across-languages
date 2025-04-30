# Request Processing Pipeline - Decorator Pattern

## Problem

You are building a web application that needs to process HTTP requests through multiple middleware components. Each middleware needs to perform specific tasks like authentication, logging, rate limiting, or request validation. Creating a separate class for every possible combination of middleware would lead to a combinatorial explosion of classes, making the system rigid and hard to extend with new middleware components.

## Solution / Implementation

The **Decorator** pattern allows us to dynamically add responsibilities to an object. In this case, we'll use it to create a flexible middleware pipeline where each middleware can wrap and modify the request/response handling.

1. **Component (`RequestHandler`):**

   - Defines the common interface for all request handlers, including a method like `handle(request)`.

2. **Concrete Component (`BaseHandler`):**

   - Represents the base request handler that processes the core request.
   - Implements the RequestHandler interface.

3. **Decorator (`Middleware`):**

   - Implements the RequestHandler interface and holds a reference to another RequestHandler.
   - Its interface conforms to the RequestHandler's interface, typically by delegating calls to the wrapped handler.

4. **Concrete Decorators (`AuthMiddleware`, `LoggingMiddleware`, `RateLimitMiddleware`):**
   - Extend the Middleware base class.
   - Add their specific functionality before or after delegating to the wrapped handler.
   - For example, `AuthMiddleware` checks authentication before passing the request to the next handler.

This structure allows you to start with a base `RequestHandler` and dynamically wrap it with any number of middleware components to build up the processing pipeline without creating specific subclasses for each combination.

## Setup

Instructions assume you are in the `structural/decorator/request_pipeline` directory.

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

Instructions assume you are in the `structural/decorator/request_pipeline` directory.

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

Instructions assume you are in the `structural/decorator/request_pipeline` directory.

### Python

```bash
cd python
python -m unittest test_request_pipeline.py
```

### TypeScript

```bash
cd typescript
npm test
```

### Go

```bash
cd go
go mod tidy
go test -v ./...
```
