import unittest
from llm_providers import (
    OpenAIFactory,
    AnthropicFactory,
    GeminiFactory,
    OllamaFactory,
    OpenAIClient, OpenAIConfiguration,
    AnthropicClient, AnthropicConfiguration,
    GeminiClient, GeminiConfiguration,
    OllamaClient, OllamaConfiguration
)

class TestLLMProviderFactories(unittest.TestCase):

    def test_openai_factory(self):
        """Verify that OpenAIFactory creates the correct family of products."""
        factory = OpenAIFactory()
        client = factory.create_client()
        config = factory.create_configuration() # Create config separately for assertion
        
        self.assertIsInstance(client, OpenAIClient, "Factory should create OpenAIClient")
        self.assertIsInstance(config, OpenAIConfiguration, "Factory should create OpenAIConfiguration")
        self.assertEqual(config.get_model(), "gpt-4o", "Config should have the correct model")
        # Check client uses the config implicitly created by factory
        self.assertIsInstance(client.config, OpenAIConfiguration, "Client should hold the correct config type")
        self.assertEqual(client.config.get_model(), "gpt-4o", "Client's config should have the correct model")
        
        # Dummy generation test
        completion = client.generate("test prompt openai")
        self.assertIn("OpenAI completion", completion)

    def test_anthropic_factory(self):
        """Verify that AnthropicFactory creates the correct family of products."""
        factory = AnthropicFactory()
        client = factory.create_client()
        config = factory.create_configuration()

        self.assertIsInstance(client, AnthropicClient)
        self.assertIsInstance(config, AnthropicConfiguration)
        self.assertEqual(config.get_model(), "claude-3.7-sonnet")
        self.assertIsInstance(client.config, AnthropicConfiguration)
        completion = client.generate("test prompt anthropic")
        self.assertIn("Anthropic completion", completion)

    def test_gemini_factory(self):
        """Verify that GeminiFactory creates the correct family of products."""
        factory = GeminiFactory()
        client = factory.create_client()
        config = factory.create_configuration()

        self.assertIsInstance(client, GeminiClient)
        self.assertIsInstance(config, GeminiConfiguration)
        self.assertEqual(config.get_model(), "gemini-2.5-pro")
        self.assertIsInstance(client.config, GeminiConfiguration)
        completion = client.generate("test prompt gemini")
        self.assertIn("Gemini completion", completion)

    def test_ollama_factory(self):
        """Verify that OllamaFactory creates the correct family of products."""
        factory = OllamaFactory()
        client = factory.create_client()
        config = factory.create_configuration()

        self.assertIsInstance(client, OllamaClient)
        self.assertIsInstance(config, OllamaConfiguration)
        self.assertEqual(config.get_model(), "llama4")
        self.assertIsInstance(client.config, OllamaConfiguration)
        completion = client.generate("test prompt ollama")
        self.assertIn("Ollama completion", completion)

if __name__ == '__main__':
    unittest.main()
