# Logging System - Adapter Pattern

## Problem

You have an existing application codebase that relies on a specific logging interface (e.g., `Logger` with methods `log_info`, `log_warning`, `log_error`). You want to integrate a new, powerful third-party logging library, but its interface is different (e.g., `ThirdPartyLogger` with a single `record(severity, message)` method). Modifying the existing application code to use the new interface directly would be intrusive and potentially require widespread changes. Modifying the third-party library is often not feasible.

## Solution / Implementation

The **Adapter** pattern provides a solution by creating a wrapper class (the Adapter) that translates calls between the incompatible interfaces.

1.  **Target Interface:** Define the interface your application expects (`Logger` in this example).
2.  **Adaptee:** The existing class or system with the incompatible interface (`ThirdPartyLogger`).
3.  **Adapter Class:** Create a class (`LoggerAdapter`) that:
    - Implements the Target interface (`Logger`).
    - Holds an instance of the Adaptee (`ThirdPartyLogger`).
    - Translates calls to the Target interface methods (`log_info`, `log_warning`, `log_error`) into calls to the Adaptee's methods (`record`).

Client code (like `ApplicationService`) interacts solely with the Target interface (`Logger`), unaware of whether it's using the original logger implementation or the adapter connected to the third-party logger.

- **Python:** Uses abstract base classes (`abc.ABC`) to define the Target interface. The Adapter class inherits from the Target ABC and delegates calls to the Adaptee instance.
- **TypeScript:** Leverages interfaces to define the Target (`Logger`). The Adapter class implements the Target interface and wraps the Adaptee.
- **Go:** Uses interfaces to define the Target (`Logger`). The Adapter struct embeds or holds an instance of the Adaptee struct and implements the Target interface methods by calling the Adaptee's methods.

## Setup

Instructions assume you are in the `structural/adapter/logging_adapter` directory.

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
```

## How to Run

Instructions assume you are in the `structural/adapter/logging_adapter` directory.

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

Instructions assume you are in the `structural/adapter/logging_adapter` directory.

### Python

```bash
cd python
python -m unittest test_logging_adapter.py
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
