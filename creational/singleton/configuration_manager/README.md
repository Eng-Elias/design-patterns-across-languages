# Configuration Manager - Singleton Pattern

## Problem

Applications often require access to configuration settings (e.g., API keys, database URLs, feature flags) from various parts of the codebase. Loading this configuration repeatedly can be inefficient, and ensuring all parts of the application use the _same_ configuration instance is crucial for consistency.

## Solution / Implementation

The Singleton pattern is applied here to create a `ConfigurationManager`. This ensures that configuration data is loaded only once and provides a single, globally accessible point to retrieve or modify settings.

- **Python:** Implements the Singleton using a class variable (`_instance`) and `threading.Lock` with double-checked locking in `__new__` for thread-safe lazy initialization.
- **TypeScript:** Uses a `private static instance` member, a `private constructor`, and a `public static getInstance()` method for controlled, lazy instantiation.
- **Go:** Leverages `sync.Once` for guaranteed thread-safe initialization and `sync.RWMutex` for safe concurrent read/write access to the configuration map after creation.

## Setup

Instructions assume you are in the `creational/singleton/configuration_manager` directory.

### Python

```bash
# No specific setup required, uses standard libraries.
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

Instructions assume you are in the `creational/singleton/configuration_manager` directory.

### Python

```bash
cd python
python main.py
```

### TypeScript

```bash
cd typescript
npm start
# or
# ts-node main.ts
```

### Go

```bash
cd go
go run main.go
```

## How to Test

Instructions assume you are in the `creational/singleton/configuration_manager` directory.

### Python

```bash
cd python
python -m unittest test_configuration_manager.py
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
