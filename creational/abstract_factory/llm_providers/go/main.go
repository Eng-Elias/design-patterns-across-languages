package main

import (
	"abstract_factory_llm_providers_go/llm_providers"
	"fmt"
)

// clientCode works with factories and products only through abstract interfaces.
func clientCode(factory llm_providers.LLMProviderFactory) {
	fmt.Printf("\nClient: Using factory: %T\n", factory)

	// Create the family of objects - just the client in this simplified version
	// The client internally gets its configuration via the factory.
	llmClient := factory.CreateClient()

	// Use the client
	prompt := "Explain the Abstract Factory design pattern in simple terms using Go."
	fmt.Printf("Client: Requesting generation for prompt: '%s...'\n", prompt[:40])
	completion, err := llmClient.Generate(prompt)
	if err != nil {
		fmt.Printf("Client: Error during generation: %v\n", err)
		return
	}
	fmt.Printf("Client: Received completion: '%s'\n", completion)
}

func main() {
	fmt.Println("--- Abstract Factory - LLM Providers Demo (Go) ---")

	// Example usage with different factories
	// We pass pointers to the factory structs
	clientCode(&llm_providers.OpenAIFactory{})
	clientCode(&llm_providers.AnthropicFactory{})
	clientCode(&llm_providers.GeminiFactory{})
	clientCode(&llm_providers.OllamaFactory{})

	fmt.Println("\n--- Demo Complete ---")
}
