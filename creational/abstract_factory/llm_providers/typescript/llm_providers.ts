// --- Abstract Products ---

export interface LLMConfiguration {
  getAPIKey(): string;
  getModel(): string;
  getBaseURL(): string;
}

export interface LLMClient {
  readonly config: LLMConfiguration; // Expose config if needed
  generate(prompt: string): Promise<string>; // Use Promise for async nature
}

// --- Abstract Factory ---

export interface LLMProviderFactory {
  createConfiguration(): LLMConfiguration;
  createClient(): LLMClient;
}

// --- Utility for simulated delay ---
const delay = (ms: number) => new Promise((resolve) => setTimeout(resolve, ms));

// --- Concrete Products & Factories ---

// --- OpenAI ---
class OpenAIConfiguration implements LLMConfiguration {
  getAPIKey(): string {
    return "DUMMY_OPENAI_KEY_12345";
  }
  getModel(): string {
    return "gpt-4o";
  }
  getBaseURL(): string {
    return "https://api.openai.com/v1";
  }
}

class OpenAIClient implements LLMClient {
  readonly config: LLMConfiguration;
  constructor(config: LLMConfiguration) {
    this.config = config;
    console.log(
      `  OpenAIClient initialized with config for model ${config.getModel()}`
    );
  }
  async generate(prompt: string): Promise<string> {
    console.log(
      `  OpenAIClient: Calling ${this.config.getBaseURL()} with model ${this.config.getModel()}`
    );
    await delay(50); // Simulate API call
    const response = `OpenAI completion for: "${prompt.substring(0, 30)}..."`;
    console.log(`  OpenAIClient: Received response.`);
    return response;
  }
}

export class OpenAIFactory implements LLMProviderFactory {
  createConfiguration(): LLMConfiguration {
    // console.log("OpenAIFactory: Creating OpenAIConfiguration."); // Less verbose
    return new OpenAIConfiguration();
  }
  createClient(): LLMClient {
    console.log("OpenAIFactory: Creating OpenAIClient.");
    const config = this.createConfiguration();
    return new OpenAIClient(config);
  }
}

// --- Anthropic ---
class AnthropicConfiguration implements LLMConfiguration {
  getAPIKey(): string {
    return "DUMMY_ANTHROPIC_KEY_67890";
  }
  getModel(): string {
    return "claude-3.7-sonnet";
  }
  getBaseURL(): string {
    return "https://api.anthropic.com/v1";
  }
}

class AnthropicClient implements LLMClient {
  readonly config: LLMConfiguration;
  constructor(config: LLMConfiguration) {
    this.config = config;
    console.log(
      `  AnthropicClient initialized with config for model ${config.getModel()}`
    );
  }
  async generate(prompt: string): Promise<string> {
    console.log(
      `  AnthropicClient: Calling ${this.config.getBaseURL()} with model ${this.config.getModel()}`
    );
    await delay(60);
    const response = `Anthropic completion for: "${prompt.substring(
      0,
      30
    )}..."`;
    console.log(`  AnthropicClient: Received response.`);
    return response;
  }
}

export class AnthropicFactory implements LLMProviderFactory {
  createConfiguration(): LLMConfiguration {
    return new AnthropicConfiguration();
  }
  createClient(): LLMClient {
    console.log("AnthropicFactory: Creating AnthropicClient.");
    const config = this.createConfiguration();
    return new AnthropicClient(config);
  }
}

// --- Gemini ---
class GeminiConfiguration implements LLMConfiguration {
  getAPIKey(): string {
    return "DUMMY_GEMINI_KEY_ABCDE";
  }
  getModel(): string {
    return "gemini-2.5-pro";
  }
  getBaseURL(): string {
    return "https://generativelanguage.googleapis.com/v1beta";
  }
}

class GeminiClient implements LLMClient {
  readonly config: LLMConfiguration;
  constructor(config: LLMConfiguration) {
    this.config = config;
    console.log(
      `  GeminiClient initialized with config for model ${config.getModel()}`
    );
  }
  async generate(prompt: string): Promise<string> {
    console.log(
      `  GeminiClient: Calling ${this.config.getBaseURL()} with model ${this.config.getModel()}`
    );
    await delay(55);
    const response = `Gemini completion for: "${prompt.substring(0, 30)}..."`;
    console.log(`  GeminiClient: Received response.`);
    return response;
  }
}

export class GeminiFactory implements LLMProviderFactory {
  createConfiguration(): LLMConfiguration {
    return new GeminiConfiguration();
  }
  createClient(): LLMClient {
    console.log("GeminiFactory: Creating GeminiClient.");
    const config = this.createConfiguration();
    return new GeminiClient(config);
  }
}

// --- Ollama ---
class OllamaConfiguration implements LLMConfiguration {
  getAPIKey(): string {
    return "N/A";
  }
  getModel(): string {
    return "llama4";
  }
  getBaseURL(): string {
    return "http://localhost:11434/api";
  }
}

class OllamaClient implements LLMClient {
  readonly config: LLMConfiguration;
  constructor(config: LLMConfiguration) {
    this.config = config;
    console.log(
      `  OllamaClient initialized with config for model ${config.getModel()}`
    );
  }
  async generate(prompt: string): Promise<string> {
    console.log(
      `  OllamaClient: Calling ${this.config.getBaseURL()} with model ${this.config.getModel()}`
    );
    await delay(40);
    const response = `Ollama completion for: "${prompt.substring(0, 30)}..."`;
    console.log(`  OllamaClient: Received response.`);
    return response;
  }
}

export class OllamaFactory implements LLMProviderFactory {
  createConfiguration(): LLMConfiguration {
    return new OllamaConfiguration();
  }
  createClient(): LLMClient {
    console.log("OllamaFactory: Creating OllamaClient.");
    const config = this.createConfiguration();
    return new OllamaClient(config);
  }
}
