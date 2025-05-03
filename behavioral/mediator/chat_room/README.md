# Chat Room - Mediator Pattern

## Problem

In systems with many interacting components (like users in a chat application), direct communication between components can lead to high coupling. Each component needs to know about many others, making the system complex, hard to understand, and difficult to maintain or modify. How can we reduce these direct dependencies and manage interactions centrally?

## Solution / Implementation

The Mediator pattern defines an object (the Mediator) that encapsulates how a set of objects (Colleagues) interact. Colleagues communicate through the Mediator instead of directly with each other, promoting loose coupling.

In this Chat Room example:

- The ChatRoom acts as the Concrete Mediator. It manages a list of users and routes messages between them.
- The ChatUser acts as a Concrete Colleague. Users register with the ChatRoom and send/receive messages only through it.

Language-specific details:

- **Python:** Uses standard classes (ChatRoom, ChatUser) to represent the Mediator and Colleagues. Communication happens via direct method calls defined in the classes.
- **TypeScript:** Defines `IChatMediator` and `IUser` interfaces for abstraction. Concrete classes ChatRoom and ChatUser implement these. Jest is used for testing, including mocks/spies to verify interactions.
- **Go:** Defines `ChatMediator` and `User` interfaces. Concrete structs ChatRoom and ChatUser implement these. Goroutines were initially used for concurrent message receiving but were later simplified to synchronous calls for clearer demo output. Go's standard `testing` package is used, along with a MockUser struct to verify behavior.

## Setup

Instructions assume you are in the `behavioral/mediator/chat_room` directory.

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

Instructions assume you are in the `behavioral/mediator/chat_room` directory.

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

Instructions assume you are in the `behavioral/mediator/chat_room` directory.

### Python

```bash
cd python
python -m unittest test_chat_room.py
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
