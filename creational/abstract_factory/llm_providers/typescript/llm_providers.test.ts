import {
  LLMProviderFactory,
  OpenAIFactory,
  AnthropicFactory,
  GeminiFactory,
  OllamaFactory,
} from "./llm_providers";

// Helper function to get concrete class names (less brittle than direct import)
// This relies on the structure within llm_providers.ts
// It's a bit of a workaround because TS interfaces don't exist at runtime.
// A better approach in a larger project might be distinct files per class.
const getImplementationName = (obj: any): string =>
  obj?.constructor?.name ?? "unknown";

describe("LLM Provider Abstract Factories", () => {
  // Helper function to test a factory
  async function testFactory(
    factory: LLMProviderFactory,
    expectedClientClass: string,
    expectedConfigClass: string,
    expectedModel: string,
    expectedCompletionSubstring: string
  ) {
    const client = factory.createClient();
    const config = factory.createConfiguration(); // Create separately for assertions

    // Check types using constructor names as a proxy
    // expect(client).toBeInstanceOf(expectedClientClass); // This won't work directly with class vars
    expect(getImplementationName(client)).toBe(expectedClientClass);
    expect(getImplementationName(config)).toBe(expectedConfigClass);

    // Check config details
    expect(config.getModel()).toBe(expectedModel);

    // Check if client holds the correct config internally (if exposed)
    // Assuming the client has a public `config` property for this test
    const clientInternalConfig = (client as any).config;
    expect(clientInternalConfig).toBeDefined();
    expect(getImplementationName(clientInternalConfig)).toBe(
      expectedConfigClass
    );
    expect(clientInternalConfig.getModel()).toBe(expectedModel);

    // Test dummy generation (async)
    const completion = await client.generate(
      `test prompt ${expectedClientClass}`
    );
    expect(completion).toContain(expectedCompletionSubstring);
  }

  test("OpenAIFactory should create OpenAI products", async () => {
    await testFactory(
      new OpenAIFactory(),
      "OpenAIClient", // Expected constructor name
      "OpenAIConfiguration",
      "gpt-4o",
      "OpenAI completion"
    );
  });

  test("AnthropicFactory should create Anthropic products", async () => {
    await testFactory(
      new AnthropicFactory(),
      "AnthropicClient",
      "AnthropicConfiguration",
      "claude-3.7-sonnet",
      "Anthropic completion"
    );
  });

  test("GeminiFactory should create Gemini products", async () => {
    await testFactory(
      new GeminiFactory(),
      "GeminiClient",
      "GeminiConfiguration",
      "gemini-2.5-pro",
      "Gemini completion"
    );
  });

  test("OllamaFactory should create Ollama products", async () => {
    await testFactory(
      new OllamaFactory(),
      "OllamaClient",
      "OllamaConfiguration",
      "llama4",
      "Ollama completion"
    );
  });
});
