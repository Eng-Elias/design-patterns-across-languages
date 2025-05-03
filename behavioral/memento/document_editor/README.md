# Document Editor - Memento Pattern

## Problem

Applications often need to support undo/redo functionality or save/restore object states without violating encapsulation. Exposing an object's internal state directly would break encapsulation and make the system fragile. How can we capture and externalize an object's internal state so that the object can be restored to this state later, without exposing its implementation details?

## Solution / Implementation

The Memento pattern allows capturing and restoring an object's internal state without violating encapsulation. It involves three key components:

1.  **Originator:** The object whose state needs to be saved (e.g., `Document`). It creates a `Memento` containing a snapshot of its current internal state and can restore its state from a `Memento`.
2.  **Memento:** An object that stores the internal state of the `Originator` (e.g., `DocumentMemento`, `DocumentState`). It should protect against access by objects other than the `Originator`. Often, it provides only getter methods or is an immutable object/struct.
3.  **Caretaker:** An object responsible for keeping track of `Memento`s (e.g., `History`, `Editor`). It requests a `Memento` from the `Originator` when saving a state and passes a `Memento` back to the `Originator` when restoring a state. The `Caretaker` never inspects or modifies the `Memento`'s content.

In this Document Editor example:

- The `Document` acts as the Originator. It holds the text content.
- The `Memento` (e.g., `DocumentState`, `concreteMemento`) stores a snapshot of the `Document`'s text content.
- The `History` (or `Editor`) acts as the Caretaker. It maintains stacks (or lists) of Mementos to manage undo and redo operations.

Language-specific details:

- **Python:** Uses classes `Document` (Originator), `Memento` (stores state), and `History` (Caretaker). The `Memento` class typically holds the state as attributes. `History` uses lists to manage undo/redo stacks.
- **TypeScript:** Defines `Document` (Originator), `Memento` (interface/class storing state), and `History` (Caretaker) classes. Uses arrays for undo/redo stacks within `History`. Interfaces might be used for stricter type checking.
- **Go:** Defines `Document` (Originator struct), `memento` (interface), `concreteMemento` (struct implementing `memento`), and `History` (Caretaker struct). Uses slices (`[]memento`) for undo/redo stacks. The `memento` interface typically exposes only methods needed for state retrieval, protecting internal details.

## Setup

Instructions assume you are in the `behavioral/memento/document_editor` directory.

### Python

```bash
# No specific setup required, uses standard libraries (unittest).
```

### TypeScript

```bash
# Install Node.js/npm if you haven't already.
cd typescript
npm install # Installs typescript, ts-node, jest, @types/jest, etc.
```

### Go

```bash
# Ensure Go is installed.
```

_(Note: Replace the module path if your repository structure differs)_

## How to Run

Instructions assume you are in the `behavioral/memento/document_editor` directory.

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

Instructions assume you are in the `behavioral/memento/document_editor` directory.

### Python

```bash
cd python
python -m unittest test_document_editor.py
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
