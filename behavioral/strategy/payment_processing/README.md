# Strategy Pattern: Payment Processing Example

## Problem

A system often needs to support multiple ways (algorithms or strategies) to perform a specific task, such as processing payments (Credit Card, PayPal, Bitcoin). Directly embedding the logic for all these algorithms within the client code (e.g., a checkout process) using conditional statements (if/else or switch) makes the code complex, difficult to maintain, and hard to extend with new algorithms without modifying the client.

## Solution / Implementation

The Strategy pattern solves this by defining a family of algorithms, encapsulating each one into separate classes (Strategies), and making them interchangeable. The client (Context) interacts with these algorithms through a common interface. This allows the algorithm to vary independently from the clients that use it.

Key components:

1.  **Context (`PaymentContext`):** Represents the object whose behavior uses a strategy (e.g., the shopping cart checkout). It maintains a reference to a Strategy object and delegates the task (payment processing) to it.
2.  **Strategy (`PaymentStrategy` interface/ABC):** Declares an interface common to all supported payment algorithms. The Context uses this interface to call the algorithm defined by a ConcreteStrategy.
3.  **Concrete Strategies (`CreditCardPayment`, `PayPalPayment`, `BitcoinPayment`):** Implement the Strategy interface, providing specific algorithms for processing payments via Credit Card, PayPal, and Bitcoin, respectively.

The `PaymentContext` can be configured with a concrete strategy object. When the payment needs to be processed, the context calls the execution method on its current strategy object. The client can change the strategy associated with the context at runtime.

Language-specific details:

- **Python:** Uses a `PaymentContext` class (Context), a `PaymentStrategy` abstract base class (Strategy), and concrete strategy classes inheriting from `PaymentStrategy`. The `PaymentContext` holds a reference to its current strategy object.
- **TypeScript:** Defines a `PaymentContext` class (Context), an `IPaymentStrategy` interface (Strategy), and concrete strategy classes implementing `IPaymentStrategy`. The `PaymentContext` holds a reference to its current strategy object.
- **Go:** Defines a `PaymentContext` struct (Context), a `PaymentStrategy` interface (Strategy), and concrete strategy structs implementing `PaymentStrategy`. The `PaymentContext` holds a reference to its current strategy object.

## Setup

Instructions assume you are in the `behavioral/strategy/payment_processing` directory.

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

Instructions assume you are in the `behavioral/strategy/payment_processing` directory.

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

Instructions assume you are in the `behavioral/strategy/payment_processing` directory.

### Python

```bash
cd python
python -m unittest test_payment_processing.py
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
