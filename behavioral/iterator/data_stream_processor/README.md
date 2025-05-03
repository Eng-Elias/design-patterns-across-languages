# Data Stream Processor - Iterator Pattern

## Problem

Many applications deal with collections of data, such as lists, streams, or datasets. Providing access to the elements of these collections without exposing their internal structure (like arrays, linked lists, etc.) is crucial for maintaining encapsulation and flexibility. How can we provide a uniform way to traverse the elements of a collection, regardless of its underlying implementation?

## Solution / Implementation

The Iterator pattern provides a way to access the elements of an aggregate object sequentially without exposing its underlying representation. It separates the traversal logic (the Iterator) from the collection itself (the Aggregate).

In this example, a `DataStream` acts as the Aggregate, holding chunks of data. A `StreamIterator` acts as the Iterator, providing methods like `hasNext` and `next` to traverse the chunks in the `DataStream`.

- **Python:** Leverages Python's built-in iterator protocol (`__iter__` and `__next__` methods in the `Iterable` and `Iterator` abstract base classes from `collections.abc`).

- **TypeScript:** Defines explicit `IAggregate` and `IIterator` interfaces with methods like `createIterator()`, `hasNext()`, and `next()`, using generics for type safety.

- **Go:** Defines `Aggregate` and `Iterator` interfaces. `DataStream` implements `Aggregate`, and `StreamIterator` implements `Iterator`, using standard Go interface patterns.

## Setup

Instructions assume you are in the `behavioral/iterator/data_stream_processor` directory.

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

Instructions assume you are in the `behavioral/iterator/data_stream_processor` directory.

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

Instructions assume you are in the `behavioral/iterator/data_stream_processor` directory.

### Python

```bash
cd python
python -m unittest test_data_stream_processor.py
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
