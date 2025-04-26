# Molecular Simulation - Prototype Pattern

## Problem

A system needs to create multiple instances of complex objects, like a `MolecularSimulation`, where the initial setup (e.g., loading large datasets, performing lengthy precomputations) is computationally expensive. Creating each new instance from scratch would be inefficient. Furthermore, we often need slightly modified versions (e.g., different simulation parameters like temperature or duration) of an existing simulation setup without repeating the expensive initialization.

## Solution / Implementation

The Prototype pattern is used. A fully initialized `MolecularSimulation` object serves as the prototype. Instead of creating new instances via a constructor that triggers the expensive setup, we `clone` the existing prototype. The cloning process is designed to be much faster than the initial setup.

Key aspects of the implementation:

- **Expensive Setup:** A method (`_performExpensiveSetup` or similar) simulates a time-consuming process, which is called only when creating the _initial_ prototype instance.
- **Cloning Method:** A `clone()` method creates a new instance. Crucially, it:

  - **Shares Large Data:** References to large, immutable (or effectively immutable) data structures (like `_precomputedStates`) are copied by reference, avoiding costly duplication.
  - **Deep Copies Mutable State:** Mutable parameters (like the `parameters` dictionary/map) are deep-copied to ensure that modifications to a clone do not affect the original prototype or other clones.
  - **Skips Setup:** The cloning mechanism bypasses the expensive setup process.

- **Python:** Uses `copy.deepcopy` for parameter cloning and explicit reference sharing for the large data list within a custom `clone` method (using `__new__` to bypass `__init__`).
- **TypeScript:** Implements a `clone()` method that manually constructs a new instance, uses `structuredClone` for deep-copying parameters, and copies the reference for the precomputed states array.
- **Go:** Implements a `Clone()` method that creates a new struct instance, performs a deep copy of the parameter map, and shares the slice reference for precomputed states.

## Setup

Instructions assume you are in the `creational/prototype/molecular_simulation` directory.

### Python

```bash
# No specific setup required, uses standard libraries (copy, time).
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

Instructions assume you are in the `creational/prototype/molecular_simulation` directory.

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

Instructions assume you are in the `creational/prototype/molecular_simulation` directory.

### Python

```bash
cd python
python -m unittest discover -v
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
