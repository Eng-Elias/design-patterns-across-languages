// Import Abstract Factory and Concrete Factories
import {
  LLMProviderFactory,
  OpenAIFactory,
  AnthropicFactory,
  GeminiFactory,
  OllamaFactory,
} from "./llm_providers";

// Client code works with factories and products through interfaces
async function clientCode(factory: LLMProviderFactory): Promise<void> {
  console.log(`\nClient: Using factory: ${factory.constructor.name}`);

  // Create the client (which gets its configuration internally)
  const llmClient = factory.createClient();

  // Use the client (awaiting the async generation)
  const prompt =
    "Explain the Abstract Factory design pattern in simple terms using TypeScript.";
  console.log(
    `Client: Requesting generation for prompt: '${prompt.substring(0, 40)}...'`
  );
  try {
    const completion = await llmClient.generate(prompt);
    console.log(`Client: Received completion: '${completion}'`);
  } catch (error) {
    console.error(`Client: Error during generation: ${error}`);
  }
}

// Main execution function
async function main() {
  console.log("--- Abstract Factory - LLM Providers Demo (TypeScript) ---");

  // Example usage with different factories
  await clientCode(new OpenAIFactory());
  await clientCode(new AnthropicFactory());
  await clientCode(new GeminiFactory());
  await clientCode(new OllamaFactory());

  console.log("\n--- Demo Complete ---");
}

// Run the main function
main().catch((error) => {
  console.error("An unexpected error occurred:", error);
});
