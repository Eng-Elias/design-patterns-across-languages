# Event Monitoring - Observer Pattern

## Problem

In many systems, when one object (the "subject" or "observable") changes its state, other dependent objects (the "observers") need to be notified and updated automatically. Directly coupling the subject to its observers makes the system inflexible and hard to maintain, as adding new observers or modifying existing ones requires changes to the subject itself.

## Solution / Implementation

The Observer pattern defines a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically. It promotes loose coupling between the subject and its observers.

Key components:

1.  **Subject (Observable):** The object that maintains a list of its dependents (observers) and provides methods to attach, detach, and notify them. When its state changes, it notifies all registered observers.
2.  **Observer:** An interface or abstract class defining an update method that the subject calls when its state changes.
3.  **Concrete Observer:** Implements the Observer interface and defines the specific action to take when notified of a change in the subject's state.
4.  **Event/State Data (Optional):** Sometimes, the notification includes data about the change (e.g., the new state or specific event details).

In this Event Monitoring example:

- The `EventSource` acts as the Subject/Observable. It generates events (e.g., system alerts, user actions).
- The `Observer` interface (or base class) defines the `update` method.
- Concrete Observers (e.g., `Logger`, `Notifier`, `DashboardDisplay`) implement the `Observer` interface and react to events (e.g., log the event, send a notification, update a UI).
- An `Event` class/struct may be used to encapsulate data passed during notification.

Language-specific details:

- **Python:** Uses classes `EventSource` (Subject), `Observer` (ABC or base class), and concrete observer classes. `EventSource` typically holds a list of observers. Can pass event data dictionaries or specific `Event` objects.
- **TypeScript:** Defines `EventSource` (Subject class), `IObserver` (interface), and concrete observer classes implementing `IObserver`. Uses arrays to manage observers. Event data often passed as objects conforming to an interface.
- **Go:** Defines `EventSource` (Subject struct), `Observer` (interface), and concrete observer structs implementing `Observer`. Uses slices (`[]Observer`) to manage observers. Event data often passed as structs.

## Setup

Instructions assume you are in the `behavioral/observer/event_monitoring` directory.

### Python

```bash
# No specific setup required
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

## How to Run

Instructions assume you are in the `behavioral/observer/event_monitoring` directory.

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

Instructions assume you are in the `behavioral/observer/event_monitoring` directory.

### Python

```bash
cd python
python -m unittest test_event_monitoring.py
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
