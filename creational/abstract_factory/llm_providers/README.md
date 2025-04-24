# Abstract Factory Pattern: LLM Providers Client

## Problem

Imagine an application that needs to interact with various Large Language Model (LLM) providers (like OpenAI, Anthropic, Google Gemini, Ollama). Each provider might require a specific client object to make API calls and a specific configuration object (containing API keys, model names, base URLs, etc.).

We want our application's client code to be able to work with any of these providers without being tightly coupled to the specific classes of each provider's client and configuration. We also want to ensure that the client object and configuration object used together are compatible (e.g., an OpenAI client should always be used with an OpenAI configuration).

## Solution / Implementation

The Abstract Factory pattern provides an interface for creating _families_ of related or dependent objects without specifying their concrete classes.

1.  **Abstract Products (`LLMClient`, `LLMConfiguration`):** Define interfaces for the distinct products we need (the client and its configuration).
2.  **Concrete Products (`OpenAIClient`, `OpenAIConfiguration`, `AnthropicClient`, etc.):** Create specific implementations of the abstract product interfaces for each LLM provider.
3.  **Abstract Factory (`LLMProviderFactory`):** Define an interface with methods for creating each abstract product (e.g., `create_client()`, `create_configuration()`).
4.  **Concrete Factories (`OpenAIFactory`, `AnthropicFactory`, etc.):** Implement the abstract factory interface. Each concrete factory is responsible for creating the specific family of products for one provider (e.g., `OpenAIFactory` creates `OpenAIClient` and `OpenAIConfiguration`).

Client code interacts only with the `LLMProviderFactory` interface and the abstract product interfaces (`LLMClient`, `LLMConfiguration`). By choosing a specific concrete factory (like `OpenAIFactory`), the client code receives a matched set of client and configuration objects, ensuring compatibility without needing to know the concrete types.

- **Python:** Uses abstract base classes (`abc.ABC`, `@abc.abstractmethod`) for the abstract factory and products. Concrete classes inherit and implement these interfaces. The `main.py` script demonstrates selecting and using different factories.
- **TypeScript:** Leverages interfaces (`LLMProviderFactory`, `LLMClient`, `LLMConfiguration`). Concrete classes implement these interfaces. The `main.ts` script demonstrates selecting and using different factories asynchronously.
- **Go:** Uses interfaces (`LLMProviderFactory`, `LLMClient`, `LLMConfiguration`). Concrete structs implement these interfaces. The `main.go` script demonstrates selecting and using different factories by passing pointers to factory structs.

## Setup

Instructions assume you are in the `creational/abstract_factory/llm_providers` directory.

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
```

## How to Run

Instructions assume you are in the `creational/abstract_factory/llm_providers` directory.

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

Instructions assume you are in the `creational/abstract_factory/llm_providers` directory.

### Python

```bash
cd python
python -m unittest test_llm_providers.py
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
