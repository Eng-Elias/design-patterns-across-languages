# Notification System - Bridge Pattern

## Problem

You need to send different types of notifications (e.g., Info, Warning, Urgent) using various delivery mechanisms (e.g., Email, SMS, Push Notification). If you create a class hierarchy where each notification type directly includes the sending logic, or if you create subclasses for every combination (like `UrgentEmailNotification`, `InfoSmsNotification`), you face several issues:

- **Class Explosion:** The number of classes grows multiplicatively (Notification Types \* Delivery Methods).
- **Tight Coupling:** The notification logic (abstraction) is tightly bound to the delivery mechanism (implementation).
- **Poor Extensibility:** Adding a new notification type requires creating new classes for _all_ existing delivery methods, and adding a new delivery method requires modifying _all_ existing notification types.

## Solution / Implementation

The **Bridge** pattern decouples an abstraction from its implementation so that the two can vary independently.

1.  **Abstraction (`Notification`):**

    - Defines the high-level interface for the operations (e.g., `send(message)`).
    - Contains a reference (the "bridge") to an object of the Implementation interface (`MessageSender`).
    - Delegates the actual work to the implementation object.

2.  **Refined Abstraction (`InfoNotification`, `WarningNotification`, `UrgentNotification`):**

    - Extends the Abstraction interface.
    - Implements the high-level logic (e.g., formatting the subject/body based on the notification type) before delegating to the implementation.

3.  **Implementation (`MessageSender`):**

    - Defines the interface for the low-level implementation classes (e.g., `send_message(subject, body)`).
    - This interface should focus on primitive operations required by the Abstraction.

4.  **Concrete Implementation (`EmailSender`, `SmsSender`, `PushNotificationSender`):**
    - Implements the Implementation interface.
    - Provides the specific logic for each delivery mechanism.

This way, you can add new notification types (Refined Abstractions) without touching the sender implementations, and add new delivery methods (Concrete Implementations) without touching the notification types.

- **Python:** Uses abstract base classes (`abc.ABC`) for both the Abstraction (`Notification`) and Implementation (`MessageSender`) interfaces. Concrete classes inherit and implement these.
- **TypeScript:** Leverages interfaces for both Abstraction (`Notification`) and Implementation (`MessageSender`). Concrete classes implement these interfaces.
- **Go:** Uses interfaces for both Abstraction (`Notification`) and Implementation (`MessageSender`). Structs implement these interfaces. The Abstraction struct holds an instance of the Implementation interface.

## Setup

Instructions assume you are in the `structural/bridge/notification_system` directory.

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

Instructions assume you are in the `structural/bridge/notification_system` directory.

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

Instructions assume you are in the `structural/bridge/notification_system` directory.

### Python

```bash
cd python
python -m unittest test_notification_system.py
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
