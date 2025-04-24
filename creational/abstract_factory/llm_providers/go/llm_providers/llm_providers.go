package llm_providers

import (
	"fmt"
	"time"
)

// --- Abstract Products ---

// LLMConfiguration defines configuration settings for an LLM provider.
// Abstract Product interface.
type LLMConfiguration interface {
	GetAPIKey() string
	GetModel() string
	GetBaseURL() string
}

// LLMClient defines the interface for an LLM client.
// Abstract Product interface.
type LLMClient interface {
	Generate(prompt string) (string, error) // Return error for potential issues
}

// --- Abstract Factory ---

// LLMProviderFactory declares methods for creating abstract products.
// Abstract Factory interface.
type LLMProviderFactory interface {
	CreateConfiguration() LLMConfiguration
	CreateClient() LLMClient
}

// --- Concrete Products: OpenAI ---

type openAIConfiguration struct{}

func (c *openAIConfiguration) GetAPIKey() string {
	return "DUMMY_OPENAI_KEY_12345"
}
func (c *openAIConfiguration) GetModel() string {
	return "gpt-4o"
}
func (c *openAIConfiguration) GetBaseURL() string {
	return "https://api.openai.com/v1"
}

type openAIClient struct {
	Config LLMConfiguration // Exported field
}

func newOpenAIClient(config LLMConfiguration) *openAIClient {
	fmt.Printf("  OpenAIClient initialized with config for model %s\n", config.GetModel())
	return &openAIClient{Config: config}
}

func (c *openAIClient) Generate(prompt string) (string, error) {
	fmt.Printf("  OpenAIClient: Calling %s with model %s\n", c.Config.GetBaseURL(), c.Config.GetModel())
	// Simulate API call delay
	time.Sleep(50 * time.Millisecond)
	shortPrompt := prompt
	if len(prompt) > 30 {
		shortPrompt = prompt[:30] + "..."
	}
	response := fmt.Sprintf("OpenAI completion for: \"%s\"", shortPrompt)
	fmt.Println("  OpenAIClient: Received response.")
	return response, nil // No error in dummy example
}

// --- Concrete Factory: OpenAI ---

type OpenAIFactory struct{}

func (f *OpenAIFactory) CreateConfiguration() LLMConfiguration {
	// fmt.Println("OpenAIFactory: Creating OpenAIConfiguration.") // Less verbose
	return &openAIConfiguration{}
}

func (f *OpenAIFactory) CreateClient() LLMClient {
	fmt.Println("OpenAIFactory: Creating OpenAIClient.")
	config := f.CreateConfiguration() // Use own method to ensure consistency
	return newOpenAIClient(config)
}

// --- Concrete Products: Anthropic ---

type anthropicConfiguration struct{}

func (c *anthropicConfiguration) GetAPIKey() string {
	return "DUMMY_ANTHROPIC_KEY_67890"
}
func (c *anthropicConfiguration) GetModel() string {
	return "claude-3.7-sonnet"
}
func (c *anthropicConfiguration) GetBaseURL() string {
	return "https://api.anthropic.com/v1"
}

type anthropicClient struct {
	Config LLMConfiguration // Exported field
}

func newAnthropicClient(config LLMConfiguration) *anthropicClient {
	fmt.Printf("  AnthropicClient initialized with config for model %s\n", config.GetModel())
	return &anthropicClient{Config: config}
}

func (c *anthropicClient) Generate(prompt string) (string, error) {
	fmt.Printf("  AnthropicClient: Calling %s with model %s\n", c.Config.GetBaseURL(), c.Config.GetModel())
	time.Sleep(60 * time.Millisecond)
	shortPrompt := prompt
	if len(prompt) > 30 {
		shortPrompt = prompt[:30] + "..."
	}
	response := fmt.Sprintf("Anthropic completion for: \"%s\"", shortPrompt)
	fmt.Println("  AnthropicClient: Received response.")
	return response, nil
}

// --- Concrete Factory: Anthropic ---

type AnthropicFactory struct{}

func (f *AnthropicFactory) CreateConfiguration() LLMConfiguration {
	return &anthropicConfiguration{}
}

func (f *AnthropicFactory) CreateClient() LLMClient {
	fmt.Println("AnthropicFactory: Creating AnthropicClient.")
	config := f.CreateConfiguration()
	return newAnthropicClient(config)
}

// --- Concrete Products: Gemini ---

type geminiConfiguration struct{}

func (c *geminiConfiguration) GetAPIKey() string {
	return "DUMMY_GEMINI_KEY_ABCDE"
}
func (c *geminiConfiguration) GetModel() string {
	return "gemini-2.5-pro"
}
func (c *geminiConfiguration) GetBaseURL() string {
	return "https://generativelanguage.googleapis.com/v1beta"
}

type geminiClient struct {
	Config LLMConfiguration // Exported field
}

func newGeminiClient(config LLMConfiguration) *geminiClient {
	fmt.Printf("  GeminiClient initialized with config for model %s\n", config.GetModel())
	return &geminiClient{Config: config}
}

func (c *geminiClient) Generate(prompt string) (string, error) {
	fmt.Printf("  GeminiClient: Calling %s with model %s\n", c.Config.GetBaseURL(), c.Config.GetModel())
	time.Sleep(55 * time.Millisecond)
	shortPrompt := prompt
	if len(prompt) > 30 {
		shortPrompt = prompt[:30] + "..."
	}
	response := fmt.Sprintf("Gemini completion for: \"%s\"", shortPrompt)
	fmt.Println("  GeminiClient: Received response.")
	return response, nil
}

// --- Concrete Factory: Gemini ---

type GeminiFactory struct{}

func (f *GeminiFactory) CreateConfiguration() LLMConfiguration {
	return &geminiConfiguration{}
}

func (f *GeminiFactory) CreateClient() LLMClient {
	fmt.Println("GeminiFactory: Creating GeminiClient.")
	config := f.CreateConfiguration()
	return newGeminiClient(config)
}

// --- Concrete Products: Ollama ---

type ollamaConfiguration struct{}

func (c *ollamaConfiguration) GetAPIKey() string {
	return "N/A" // Usually local
}
func (c *ollamaConfiguration) GetModel() string {
	return "llama4"
}
func (c *ollamaConfiguration) GetBaseURL() string {
	return "http://localhost:11434/api"
}

type ollamaClient struct {
	Config LLMConfiguration // Exported field
}

func newOllamaClient(config LLMConfiguration) *ollamaClient {
	fmt.Printf("  OllamaClient initialized with config for model %s\n", config.GetModel())
	return &ollamaClient{Config: config}
}

func (c *ollamaClient) Generate(prompt string) (string, error) {
	fmt.Printf("  OllamaClient: Calling %s with model %s\n", c.Config.GetBaseURL(), c.Config.GetModel())
	time.Sleep(40 * time.Millisecond)
	shortPrompt := prompt
	if len(prompt) > 30 {
		shortPrompt = prompt[:30] + "..."
	}
	response := fmt.Sprintf("Ollama completion for: \"%s\"", shortPrompt)
	fmt.Println("  OllamaClient: Received response.")
	return response, nil
}

// --- Concrete Factory: Ollama ---

type OllamaFactory struct{}

func (f *OllamaFactory) CreateConfiguration() LLMConfiguration {
	return &ollamaConfiguration{}
}

func (f *OllamaFactory) CreateClient() LLMClient {
	fmt.Println("OllamaFactory: Creating OllamaClient.")
	config := f.CreateConfiguration()
	return newOllamaClient(config)
}
