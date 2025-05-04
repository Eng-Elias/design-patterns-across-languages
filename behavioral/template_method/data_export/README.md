# Template Method Pattern: Data Export Example

## Problem

Imagine you need to implement a system for exporting data in various formats (CSV, JSON, XML, etc.). The overall export process follows a similar sequence of steps: fetch the data, format it according to the target specification, and then save or send the formatted data. Implementing this logic separately for each format can lead to code duplication, especially for the common steps or the overall orchestration logic. If the sequence of steps needs to change, you'd have to modify multiple classes.

## Solution / Implementation

The Template Method pattern defines the skeleton of an algorithm in a base class (superclass) but lets subclasses override specific steps of the algorithm without changing its structure.

Key components:

1.  **Abstract Class (`DataExporter`)**:

    - Defines the `templateMethod()` (`exportData` in our example) which calls a series of abstract or concrete "primitive" operations. This method outlines the overall algorithm structure.
    - Declares abstract `primitiveOperation()` methods (e.g., `fetchData`, `formatData`, `saveData`) that concrete subclasses must implement.
    - May include `hook()` methods, which are optional steps with default (often empty) implementations, allowing subclasses to provide custom behavior at specific points in the algorithm.

2.  **Concrete Class (`CsvExporter`, `JsonExporter`)**:
    - Inherits from the Abstract Class.
    - Implements the abstract `primitiveOperation()` methods required by the `templateMethod`.
    - Optionally overrides `hook()` methods to customize the algorithm.

In this Data Export example:

- The `DataExporter` acts as the Abstract Class. It defines the `exportData` template method which orchestrates the fetching, formatting, and saving steps.
- `fetchData` (can be concrete or abstract), `formatData`, and `saveData` are the essential primitive operations (abstract in this case for format/save).
- `CsvExporter` and `JsonExporter` are Concrete Classes providing specific implementations for formatting data as CSV or JSON and saving it (e.g., printing to console or simulating file save).

Language-specific details:

- **Python:** Uses an abstract base class (`DataExporter`) with `abc.ABC` and `@abc.abstractmethod`. Concrete classes inherit and implement abstract methods.
- **TypeScript:** Uses an `abstract class` (`DataExporter`) with `abstract` methods. Concrete classes `extend` the abstract class and implement abstract methods.
- **Go:** Go doesn't have classes or inheritance. The pattern is often simulated using interfaces and composition. A common approach is to define an interface for the varying steps (`FormatterSaver`) and a struct (`DataExporter`) that uses this interface within its template method (`Export`). Concrete types (`CsvExporter`, `JsonExporter`) implement the interface.

## Setup

Instructions assume you are in the `behavioral/template_method/data_export` directory.

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

Instructions assume you are in the `behavioral/template_method/data_export` directory.

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

Instructions assume you are in the `behavioral/template_method/data_export` directory.

### Python

```bash
cd python
python -m unittest test_data_exporter.py
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
