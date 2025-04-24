from llm_providers import (
    LLMProviderFactory,
    OpenAIFactory,
    AnthropicFactory,
    GeminiFactory,
    OllamaFactory
)

def client_code(factory: LLMProviderFactory):
    """The client code works with factories and products only through abstract types."""
    print(f"\nClient: Using factory: {factory.__class__.__name__}")
    
    # Create the client (which internally gets its correct configuration)
    llm_client = factory.create_client()
    
    # Use the client
    prompt = "Explain the Abstract Factory design pattern in simple terms."
    print(f"Client: Requesting generation for prompt: '{prompt[:40]}...'" )
    completion = llm_client.generate(prompt)
    print(f"Client: Received completion: '{completion}'")

if __name__ == "__main__":
    print("--- Abstract Factory - LLM Providers Demo (Python) ---")

    # Example usage with different factories
    client_code(OpenAIFactory())
    client_code(AnthropicFactory())
    client_code(GeminiFactory())
    client_code(OllamaFactory())

    print("\n--- Demo Complete ---")
