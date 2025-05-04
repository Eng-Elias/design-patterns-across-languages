# Order Processing - State Pattern

## Problem

An object's behavior often depends on its internal state. For example, an order in an e-commerce system behaves differently depending on whether it's `New`, `PendingPayment`, `Shipped`, or `Delivered`. Implementing this behavior using large conditional statements (if/else or switch) within the object's methods can lead to complex, hard-to-maintain code, especially as the number of states and state-dependent behaviors grows.

## Solution / Implementation

The State pattern allows an object to alter its behavior when its internal state changes. The object appears to change its class. It encapsulates state-specific behavior into separate state objects and delegates the behavior execution to the current state object.

Key components:

1.  **Context:** The object whose behavior changes based on its state (e.g., `Order`). It maintains an instance of a ConcreteState subclass that defines the current state and delegates state-specific requests to it.
2.  **State:** An interface or abstract class defining the methods representing state-specific behaviors (e.g., `processPayment()`, `shipOrder()`, `deliverOrder()`).
3.  **ConcreteState:** Subclasses that implement the State interface, providing the specific behavior for each state. Each ConcreteState handles requests differently and can transition the Context to a new state.

In this Order Processing example:

- The `Order` acts as the Context. Its behavior (e.g., ability to accept payment, ship, deliver) depends on its current status.
- The `OrderState` interface defines methods like `processPayment`, `ship`, `deliver`, etc.
- Concrete States (`NewOrderState`, `PendingPaymentState`, `ShippedState`, `DeliveredState`, `CancelledState`) implement `OrderState`. Each handles the methods based on the rules of that state (e.g., you can only ship an order that is in the `PendingPayment` state after payment is processed).

Language-specific details:

- **Python:** Uses an `Order` class (Context), an `OrderState` abstract base class (State), and concrete state classes inheriting from `OrderState`. The `Order` holds a reference to its current state object.
- **TypeScript:** Defines an `Order` class (Context), an `IOrderState` interface (State), and concrete state classes implementing `IOrderState`. The `Order` holds a reference to its current state object.
- **Go:** Defines an `Order` struct (Context), an `OrderState` interface (State), and concrete state structs implementing `OrderState`. The `Order` holds a reference to its current state object.

## Setup

Instructions assume you are in the `behavioral/state/order_processing` directory.

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

Instructions assume you are in the `behavioral/state/order_processing` directory.

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

Instructions assume you are in the `behavioral/state/order_processing` directory.

### Python

```bash
cd python
python -m unittest test_order_processing.py
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
