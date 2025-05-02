# Task Scheduler - Command Pattern

## Problem

We need to build a task scheduler that can execute various types of tasks at a later time (e.g., sending emails, generating reports, running database backups). The scheduler itself should not need to know the specific details of _how_ each task is performed. It should be easy to add new types of tasks without modifying the scheduler core. We also want to queue these tasks and execute them sequentially.

## Solution / Implementation

The Command pattern is ideal here. We encapsulate each task request as a Command object.

1.  **Command:** An interface (`Command`) defining an `execute` method.
2.  **Concrete Commands:** Specific task classes (e.g., `SendEmailCommand`, `GenerateReportCommand`, `RunDatabaseBackupCommand`) that implement the `Command` interface. Each command holds:
    - A reference to a **Receiver** object that knows how to perform the actual work.
    - Any parameters needed for the task (e.g., email address, report type).
    - The `execute` method calls the appropriate method on its Receiver.
3.  **Receiver:** Objects that perform the actual work (e.g., `EmailService`, `ReportGenerator`, `DatabaseService`).
4.  **Invoker:** The `TaskScheduler` class. It holds a collection of `Command` objects. It has methods like `add_task` and `run_pending_tasks`. The `run_pending_tasks` method simply iterates through its commands and calls `execute()` on each. The scheduler is decoupled from the concrete tasks.
5.  **Client:** The main application code that creates Receiver objects, creates Concrete Command objects (linking receivers and parameters), and adds these commands to the Task Scheduler (Invoker).

- **Python:** Uses an ABC for `Command`. Receivers are simple classes. Concrete commands store receiver and args. The `TaskScheduler` holds a list of commands.
- **TypeScript:** Uses an interface for `Command`. Classes for receivers and concrete commands. The `TaskScheduler` holds an array of `Command` objects.
- **Go:** Uses an interface for `Command`. Structs for receivers and concrete commands (implementing the interface). `TaskScheduler` uses a slice of `Command` interface types.

## Setup

Instructions assume you are in the `behavioral/command/task_scheduler` directory.

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

Instructions assume you are in the `behavioral/command/task_scheduler` directory.

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

Instructions assume you are in the `behavioral/command/task_scheduler` directory.

### Python

```bash
cd python
python -m unittest test_task_scheduler.py
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
