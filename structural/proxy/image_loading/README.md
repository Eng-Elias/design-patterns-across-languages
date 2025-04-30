# Image Loading - Proxy Pattern (Virtual Proxy)

## Problem

Your application needs to display potentially large or numerous images (e.g., high-resolution photos, document scans). Loading these images from disk or network can be a time-consuming and resource-intensive operation. If you load all images immediately when the application starts or when the document/gallery is opened, the user might experience significant delays and high memory consumption, even if they only view a few images initially.

## Solution / Implementation

The **Proxy** pattern provides a surrogate or placeholder for another object to control access to it. A **Virtual Proxy** is a specific type of proxy used to delay the creation and initialization of a resource-intensive object until it is actually needed (Lazy Initialization).

1.  **Subject Interface (`Image`):**

    - Defines the common interface for both the `RealImage` (Real Subject) and the `ProxyImage` (Proxy).
    - Clients interact with objects through this interface.
    - Contains methods like `display()` and `get_filename()`.

2.  **Real Subject (`RealImage`):**

    - The actual object that performs the core task (loading and displaying the image data).
    - Its constructor (`__init__` or equivalent) simulates the expensive loading operation (e.g., reading from disk, network request).
    - The `display()` method shows the fully loaded image.

3.  **Proxy (`ProxyImage`):**

    - Implements the same `Image` interface as the `RealImage`.
    - Holds a reference to the `RealImage` object, but this reference is initially `null` or `None`.
    - Stores information needed to create the `RealImage` (like the `filename`).
    - When a client calls a method like `display()` on the `ProxyImage`:
      - The Proxy checks if the `RealImage` instance has already been created.
      - If not (`null`/`None`), the Proxy creates the `RealImage` instance (triggering the expensive loading operation) and stores the reference.
      - The Proxy then delegates the `display()` call to the now-available `RealImage` instance.
      - If the `RealImage` was already created (on a previous call), the Proxy simply delegates the call immediately.
    - Methods that don't require the fully loaded object (like `get_filename()`) can be handled directly by the Proxy without creating the `RealImage`.

4.  **Client (`main` script):**
    - Interacts with the `ProxyImage` through the `Image` interface.
    - The client is unaware that it might be dealing with a proxy instead of the real object initially. The lazy loading mechanism is hidden behind the common interface.

This approach ensures that the expensive image loading only happens when an image is actually requested for display, improving application startup time and reducing initial memory usage.

- **Python:** Uses an abstract base class for the Subject, a concrete class for the Real Subject, and a Proxy class holding an optional reference to the Real Subject.
- **TypeScript:** Uses an interface for the Subject, a class for the Real Subject, and a Proxy class implementing the Subject interface and managing the lazy loading.
- **Go:** Uses an interface for the Subject, a struct for the Real Subject, and a Proxy struct implementing the Subject interface, holding a pointer (initially nil) to the Real Subject struct.

## Setup

Instructions assume you are in the `structural/proxy/image_loading` directory.

### Python

```bash
# No specific setup required beyond Python itself.
# Uses standard libraries (unittest.mock for tests, time).
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
cd go
go mod tidy # To fetch dependencies like testify if needed
```

## How to Run

Instructions assume you are in the `structural/proxy/image_loading` directory.

### Python

```bash
cd python
python main.py
```

### TypeScript

```bash
cd typescript
npm start
# Or build first:
# npm run build && node dist/main.js
```

### Go

```bash
cd go
go run main.go
```

## How to Test

Instructions assume you are in the `structural/proxy/image_loading` directory.

### Python

```bash
cd python
python -m unittest test_image_proxy.py
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
