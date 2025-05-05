# Visitor Pattern: Code Analysis Example

## Problem

When working with complex object structures, like an Abstract Syntax Tree (AST) representing code, we often need to perform various operations or analyses on the nodes (e.g., calculate complexity, check syntax, pretty-print). Directly adding methods for each new operation into the AST node classes makes these classes bulky, harder to maintain, and violates the Open/Closed Principle, as modifying the node classes is required for every new analysis.

## Solution / Implementation

The Visitor pattern addresses this by separating the operations from the object structure. It allows defining new operations (Visitors) without changing the classes of the elements (AST Nodes) on which they operate.

Key components:

1.  **Element (`Node` interface/ABC):** Declares an `accept` method that takes a visitor as an argument. (Implemented by AST nodes).
2.  **Concrete Elements (`FunctionDefinitionNode`, `IfStatementNode`, etc.):** Implement the Element interface. Their `accept` method typically calls the appropriate `visit` method on the visitor, passing itself (`visitor.visitConcreteElement(this)`).
3.  **Visitor (`Visitor` interface/ABC):** Declares a `visit` method for each type of Concrete Element in the structure (e.g., `visitFunctionDefinition`, `visitIfStatement`).
4.  **Concrete Visitors (`ComplexityVisitor`, `SyntaxCheckVisitor`, `PrettyPrintVisitor`):** Implement the Visitor interface. Each `visit` method contains the logic for a specific operation on the corresponding Concrete Element.
5.  **Object Structure (The AST):** The collection of Element objects. The client typically traverses this structure.
6.  **Client (e.g., `main.py`/`main.ts`/`main.go`):** Creates Concrete Visitor objects and initiates the traversal of the Object Structure (AST), calling `accept` on elements and passing the visitor.

This pattern enables adding new analysis capabilities simply by creating new Concrete Visitor classes.

Language-specific details:

- **Python:** Uses a `Visitor` abstract base class (Strategy) and concrete visitor classes inheriting from it. An abstract `Node` class defines `accept`. Concrete node classes implement `accept`.
- **TypeScript:** Defines a `Visitor` interface (Strategy) and concrete visitor classes implementing it. An abstract `Node` class defines `accept`. Concrete node classes implement `accept`.
- **Go:** Defines a `Visitor` interface (Strategy) and concrete visitor structs implementing it. A `Node` interface defines `Accept`. Concrete node structs implement `Accept`.

## Setup

Instructions assume you are in the `behavioral/visitor/code_analyzer` directory.

### Python

```bash
# No specific setup required
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

Instructions assume you are in the `behavioral/visitor/code_analyzer` directory.

### Python

```bash
cd python
python main.py
```

### TypeScript

```bash
cd typescript
npm run start
```

### Go

```bash
cd go
go run main.go
```

## How to Test

Instructions assume you are in the `behavioral/visitor/code_analyzer` directory.

### Python

```bash
cd python
python -m unittest code_analyzer.test_code_analyzer.py
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
