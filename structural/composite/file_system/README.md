# File System - Composite Pattern

## Problem

You need to represent part-whole hierarchies of objects, like a file system where directories can contain files or other directories. You want clients to treat individual objects (files) and compositions of objects (directories) uniformly. For instance, calculating the total size of a directory should involve summing the sizes of its contents, including nested subdirectories, using the same interface as getting the size of a single file.

## Solution / Implementation

The **Composite** pattern allows you to compose objects into tree structures to represent part-whole hierarchies. Composite lets clients treat individual objects and compositions of objects uniformly.

1.  **Component (`FileSystemComponent`):**

    - Declares the common interface for objects in the composition (both leaves and composites).
    - Includes methods like `get_name()`, `get_size()`, `display()`.
    - May optionally include methods for managing children (`add`, `remove`, `get_child`), possibly with default implementations that raise errors or do nothing.

2.  **Leaf (`File`):**

    - Represents leaf objects in the composition (objects that have no children).
    - Implements the Component interface.
    - Defines behavior for primitive objects (e.g., `get_size()` returns the file's inherent size).
    - Child management methods typically raise errors.

3.  **Composite (`Directory`):**
    - Represents composite objects (objects that can have children).
    - Stores child components.
    - Implements the Component interface.
    - Implements child-related operations (`add`, `remove`, `get_child`).
    - Implements operations defined in the Component interface by delegating to its children (e.g., `get_size()` iterates over children and sums their sizes; `display()` iterates and calls `display()` on children).

Clients interact with objects through the Component interface. If the object is a Leaf, the request is handled directly. If it's a Composite, it usually forwards the request to its child components, potentially performing additional operations before or after forwarding.

- **Python:** Uses an abstract base class (`abc.ABC`) for the Component interface. `File` (Leaf) and `Directory` (Composite) inherit from it.
- **TypeScript:** Leverages an abstract class or interface for the Component. `File` (Leaf) and `Directory` (Composite) implement/extend it.
- **Go:** Uses an interface for the Component. `File` (Leaf struct) and `Directory` (Composite struct) implement the interface. The `Directory` struct contains a slice of Component interfaces.

## Setup

Instructions assume you are in the `structural/composite/file_system` directory.

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

Instructions assume you are in the `structural/composite/file_system` directory.

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

Instructions assume you are in the `structural/composite/file_system` directory.

### Python

```bash
cd python
python -m unittest test_file_system.py
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
