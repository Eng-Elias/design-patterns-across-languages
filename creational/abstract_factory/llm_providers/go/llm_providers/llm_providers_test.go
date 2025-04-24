package llm_providers

import (
	"reflect" // Use reflection for type checking
	"strings"
	"testing"
)

// Helper function to test factories using reflection for type checks
func testFactory(t *testing.T, factory LLMProviderFactory, expectedClientType reflect.Type, expectedConfigType reflect.Type, expectedModel string, expectedCompletionSubstring string) {
	// Suppress verbose output during tests if needed, but keep for demonstration
	// t.Logf("Testing factory: %T", factory)

	client := factory.CreateClient()
	config := factory.CreateConfiguration() // Create separately for checks

	// Check client type using reflection
	clientType := reflect.TypeOf(client)
	if clientType != expectedClientType {
		t.Errorf("Factory %T created client of type %v, expected %v", factory, clientType, expectedClientType)
	}

	// Check config type using reflection
	configType := reflect.TypeOf(config)
	if configType != expectedConfigType {
		t.Errorf("Factory %T created config of type %v, expected %v", factory, configType, expectedConfigType)
	}

	// Check config model
	if config.GetModel() != expectedModel {
		t.Errorf("Factory %T created config with model '%s', expected '%s'", factory, config.GetModel(), expectedModel)
	}

	// Check client uses the correct config implicitly
	// We need to access the private 'Config' field via reflection for a robust test
	clientValue := reflect.ValueOf(client)
	// If client is a pointer, get the element it points to
	if clientValue.Kind() == reflect.Ptr {
		clientValue = clientValue.Elem()
	}
	// Access the 'Config' field (now exported)
	internalConfigField := clientValue.FieldByName("Config")
	if !internalConfigField.IsValid() {
		t.Fatalf("Factory %T client type %T does not have an accessible 'Config' field for testing", factory, client)
	}
	// This should no longer panic
	internalConfig := internalConfigField.Interface().(LLMConfiguration) // Assert to interface

	internalConfigType := reflect.TypeOf(internalConfig)
	if internalConfigType != expectedConfigType {
		t.Errorf("Factory %T client internal config is of type %v, expected %v", factory, internalConfigType, expectedConfigType)
	}
	if internalConfig.GetModel() != expectedModel {
		t.Errorf("Factory %T client internal config has model '%s', expected '%s'", factory, internalConfig.GetModel(), expectedModel)
	}


	// Test dummy generation
	completion, err := client.Generate("test prompt")
	if err != nil {
		t.Errorf("Factory %T client generate failed: %v", factory, err)
	}
	if !strings.Contains(completion, expectedCompletionSubstring) {
		t.Errorf("Factory %T client generated '%s', expected substring '%s'", factory, completion, expectedCompletionSubstring)
	}
}

// --- Test Cases ---

func TestOpenAIFactory(t *testing.T) {
	factory := &OpenAIFactory{}
	// Get expected types using reflection on pointers to the structs
	expectedClientType := reflect.TypeOf(&openAIClient{})
	expectedConfigType := reflect.TypeOf(&openAIConfiguration{})
	testFactory(t, factory, expectedClientType, expectedConfigType, "gpt-4o", "OpenAI completion")
}

func TestAnthropicFactory(t *testing.T) {
	factory := &AnthropicFactory{}
	expectedClientType := reflect.TypeOf(&anthropicClient{})
	expectedConfigType := reflect.TypeOf(&anthropicConfiguration{})
	testFactory(t, factory, expectedClientType, expectedConfigType, "claude-3.7-sonnet", "Anthropic completion")
}

func TestGeminiFactory(t *testing.T) {
	factory := &GeminiFactory{}
	expectedClientType := reflect.TypeOf(&geminiClient{})
	expectedConfigType := reflect.TypeOf(&geminiConfiguration{})
	testFactory(t, factory, expectedClientType, expectedConfigType, "gemini-2.5-pro", "Gemini completion")
}

func TestOllamaFactory(t *testing.T) {
	factory := &OllamaFactory{}
	expectedClientType := reflect.TypeOf(&ollamaClient{})
	expectedConfigType := reflect.TypeOf(&ollamaConfiguration{})
	testFactory(t, factory, expectedClientType, expectedConfigType, "llama4", "Ollama completion")
}
