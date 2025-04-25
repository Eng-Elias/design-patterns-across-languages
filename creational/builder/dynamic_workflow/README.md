# Builder Pattern: Dynamic Workflow

## Problem

Imagine needing to construct a complex automation workflow composed of various steps (like sending emails, running scripts, posting Slack notifications). The sequence and parameters of these steps might vary significantly for different workflows. Directly constructing such a workflow object with a complex constructor or numerous setter methods can become cumbersome and error-prone.

We want a way to create different representations of a complex object (the workflow) using the same construction process. The construction process should allow steps to be added dynamically and should hide the complex internal structure of the final workflow object from the client code.

## Solution / Implementation

The Builder pattern separates the construction of a complex object from its representation, so that the same construction process can create different representations.

1.  **Product (`DynamicWorkflow`):** The complex object being built. It contains a list of steps and methods to execute them based on registered handlers.
2.  **Builder (`DynamicWorkflowBuilder`):** An interface or abstract class for creating the parts of the Product. It provides methods like `add_step` to configure the product.
3.  **Concrete Builder (`DynamicWorkflowBuilder` implementation):** Implements the Builder interface. It keeps track of the steps being added and provides a method (`build()`) to return the final Product.
4.  **Director (Implicit):** The client code that uses the Concrete Builder. It calls the builder's methods (`add_step`) in a specific sequence to construct the desired Product and then calls `build()`.

This pattern allows for step-by-step construction, makes the construction process independent of the parts that make up the object, and provides better control over the construction process. Method chaining (`builder.add_step(...).add_step(...)`) is often used for a fluent interface.

- **Python:** Uses a `DynamicWorkflowBuilder` class with an `add_step` method returning `self` for chaining. The `build` method creates a `DynamicWorkflow` instance which holds steps and handlers (implemented as methods mapped in a dictionary).
- **TypeScript:** Similar structure using classes (`DynamicWorkflowBuilder`, `DynamicWorkflow`). `addStep` returns `this` for chaining. Handlers are methods stored in a `Map`.
- **Go:** Uses a `DynamicWorkflowBuilder` struct with methods like `AddStep` that return the builder pointer (`*DynamicWorkflowBuilder`) for chaining. The `Build` method returns a `*DynamicWorkflow` struct. Handlers are functions stored in a map.

## Setup

Instructions assume you are in the `creational/builder/dynamic_workflow` directory.

### Python

```bash
# No specific setup required, uses standard libraries (unittest).
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

Instructions assume you are in the `creational/builder/dynamic_workflow` directory.

### Python

```bash
cd python
python main.py
```

### TypeScript

```bash
cd typescript
npm start
# or (after npm install)
# ts-node main.ts
```

### Go

```bash
cd go
go run main.go
```

## How to Test

Instructions assume you are in the `creational/builder/dynamic_workflow` directory.

### Python

```bash
cd python
python -m unittest test_dynamic_workflow.py
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
