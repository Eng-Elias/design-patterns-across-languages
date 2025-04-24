# Document Processor - Factory Method Pattern

## Problem

A system needs to create different types of documents (like TXT, JSON, HTML) but wants to decouple the client code that _uses_ the documents from the specific _creation_ logic of each document type. The core system shouldn't need to know all concrete document types upfront, allowing for easy extension with new types later.

## Solution / Implementation

The Factory Method pattern is used. An abstract `DocumentProcessor` (Creator) defines a `create_document` method (the Factory Method). Concrete subclasses like `TextProcessor`, `JSONProcessor`, `HTMLProcessor` implement this method to return instances of specific `Document` types (`TextDocument`, `JSONDocument`, `HTMLDocument` - the Products). Client code interacts with the `DocumentProcessor` interface to get documents, delegating the actual instantiation to the appropriate subclass.

- **Python:** Uses abstract base classes (`abc.ABC`, `@abc.abstractmethod`) to define the `Document` and `DocumentProcessor` interfaces. Concrete classes inherit and implement the interfaces and the factory method.
- **TypeScript:** Leverages interfaces (`Document`, `DocumentProcessor`) for defining contracts. Concrete classes implement these interfaces and the factory method.
- **Go:** Uses interfaces (`Document`, `DocumentProcessor`) to define behavior. Structs representing concrete documents and processors implement these interfaces. The factory method in concrete processors returns the specific document type satisfying the `Document` interface.

## Setup

Instructions assume you are in the `creational/factory_method/document_processor` directory.

### Python

```bash
# No specific setup required, uses standard libraries (abc).
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

Instructions assume you are in the `creational/factory_method/document_processor` directory.

### Python

```bash
cd python
python main.py
```

### TypeScript

```bash
cd typescript
npm start
# or
# ts-node main.ts
```

### Go

```bash
cd go
go run main.go
```

## How to Test

Instructions assume you are in the `creational/factory_method/document_processor` directory.

### Python

```bash
cd python
python -m unittest test_document_processor.py
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
