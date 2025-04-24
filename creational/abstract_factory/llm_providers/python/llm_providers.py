from abc import ABC, abstractmethod

# --- Abstract Products ---

class LLMConfiguration(ABC):
    """Abstract Product: Defines configuration settings for an LLM provider."""
    @abstractmethod
    def get_api_key(self) -> str:
        pass

    @abstractmethod
    def get_model(self) -> str:
        pass

    @abstractmethod
    def get_base_url(self) -> str:
        pass

class LLMClient(ABC):
    """Abstract Product: Defines the interface for an LLM client."""
    def __init__(self, config: LLMConfiguration):
        self.config = config # Store config for internal use
        print(f"  {self.__class__.__name__} initialized with config for model {config.get_model()}")

    @abstractmethod
    def generate(self, prompt: str) -> str:
        """Generates a completion for the given prompt."""
        pass

# --- Abstract Factory ---

class LLMProviderFactory(ABC):
    """Abstract Factory: Declares methods for creating abstract products."""
    @abstractmethod
    def create_configuration(self) -> LLMConfiguration:
        pass

    @abstractmethod
    def create_client(self) -> LLMClient:
        pass

# --- Concrete Products & Factories ---

# --- OpenAI ---
class OpenAIConfiguration(LLMConfiguration):
    def get_api_key(self) -> str: return "DUMMY_OPENAI_KEY_12345"
    def get_model(self) -> str: return "gpt-4o"
    def get_base_url(self) -> str: return "https://api.openai.com/v1"

class OpenAIClient(LLMClient):
    def generate(self, prompt: str) -> str:
        print(f"  OpenAIClient: Calling {self.config.get_base_url()} with model {self.config.get_model()}")
        response = f'OpenAI completion for: "{prompt[:30]}..."'
        print(f"  OpenAIClient: Received response.")
        return response

class OpenAIFactory(LLMProviderFactory):
    def create_configuration(self) -> LLMConfiguration:
        print("OpenAIFactory: Creating OpenAIConfiguration.")
        return OpenAIConfiguration()
    def create_client(self) -> LLMClient:
        print("OpenAIFactory: Creating OpenAIClient.")
        config = self.create_configuration()
        return OpenAIClient(config)

# --- Anthropic ---
class AnthropicConfiguration(LLMConfiguration):
    def get_api_key(self) -> str: return "DUMMY_ANTHROPIC_KEY_67890"
    def get_model(self) -> str: return "claude-3.7-sonnet"
    def get_base_url(self) -> str: return "https://api.anthropic.com/v1"

class AnthropicClient(LLMClient):
    def generate(self, prompt: str) -> str:
        print(f"  AnthropicClient: Calling {self.config.get_base_url()} with model {self.config.get_model()}")
        response = f'Anthropic completion for: "{prompt[:30]}..."'
        print(f"  AnthropicClient: Received response.")
        return response

class AnthropicFactory(LLMProviderFactory):
    def create_configuration(self) -> LLMConfiguration:
        print("AnthropicFactory: Creating AnthropicConfiguration.")
        return AnthropicConfiguration()
    def create_client(self) -> LLMClient:
        print("AnthropicFactory: Creating AnthropicClient.")
        config = self.create_configuration()
        return AnthropicClient(config)

# --- Gemini ---
class GeminiConfiguration(LLMConfiguration):
    def get_api_key(self) -> str: return "DUMMY_GEMINI_KEY_ABCDE"
    def get_model(self) -> str: return "gemini-2.5-pro"
    def get_base_url(self) -> str: return "https://generativelanguage.googleapis.com/v1beta"

class GeminiClient(LLMClient):
    def generate(self, prompt: str) -> str:
        print(f"  GeminiClient: Calling {self.config.get_base_url()} with model {self.config.get_model()}")
        response = f'Gemini completion for: "{prompt[:30]}..."'
        print(f"  GeminiClient: Received response.")
        return response

class GeminiFactory(LLMProviderFactory):
    def create_configuration(self) -> LLMConfiguration:
        print("GeminiFactory: Creating GeminiConfiguration.")
        return GeminiConfiguration()
    def create_client(self) -> LLMClient:
        print("GeminiFactory: Creating GeminiClient.")
        config = self.create_configuration()
        return GeminiClient(config)

# --- Ollama ---
class OllamaConfiguration(LLMConfiguration):
    def get_api_key(self) -> str: return "N/A"
    def get_model(self) -> str: return "llama4"
    def get_base_url(self) -> str: return "http://localhost:11434/api"

class OllamaClient(LLMClient):
    def generate(self, prompt: str) -> str:
        print(f"  OllamaClient: Calling {self.config.get_base_url()} with model {self.config.get_model()}")
        response = f'Ollama completion for: "{prompt[:30]}..."'
        print(f"  OllamaClient: Received response.")
        return response

class OllamaFactory(LLMProviderFactory):
    def create_configuration(self) -> LLMConfiguration:
        print("OllamaFactory: Creating OllamaConfiguration.")
        return OllamaConfiguration()
    def create_client(self) -> LLMClient:
        print("OllamaFactory: Creating OllamaClient.")
        config = self.create_configuration()
        return OllamaClient(config)
