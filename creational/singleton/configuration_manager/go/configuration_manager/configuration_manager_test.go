package configuration_manager

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	t.Run("should return the same instance on multiple calls", func(t *testing.T) {
		instance1 := GetInstance()
		instance2 := GetInstance()

		if instance1 != instance2 {
			t.Errorf("Expected same instance, but got different instances: %p != %p", instance1, instance2)
		}
	})

	t.Run("should initialize only once concurrently", func(t *testing.T) {
		// Reset instance and once for testing (this is tricky in Go without helpers)
		// For a robust test, you might need build tags or separate test helpers.
		// Here, we'll rely on the GetInstance logic and test concurrent access.
		instance = nil         // Reset global instance (use with caution, not ideal for parallel tests)
		once = sync.Once{} // Reset sync.Once

		numGoroutines := 100
		instances := make([]*ConfigurationManager, numGoroutines)
		var wg sync.WaitGroup
		wg.Add(numGoroutines)

		// Simulate concurrent GetInstance calls
		for i := 0; i < numGoroutines; i++ {
			go func(index int) {
				defer wg.Done()
				instances[index] = GetInstance()
			}(i)
		}

		wg.Wait() // Wait for all goroutines to complete

		// Check if all goroutines got the same instance
		firstInstance := instances[0]
		if firstInstance == nil {
			t.Fatal("First instance is nil after concurrent initialization")
		}
		for i := 1; i < numGoroutines; i++ {
			if instances[i] != firstInstance {
				t.Errorf("Goroutine %d got a different instance (%p) than the first (%p)", i, instances[i], firstInstance)
			}
		}

		// Clean up after reset for subsequent tests if needed
		instance = nil
		once = sync.Once{}
	})
}

func TestConfigurationSettings(t *testing.T) {
	// Ensure a clean state for this test sequence
	instance = nil
	once = sync.Once{}

	config := GetInstance()

	t.Run("should get initial values correctly", func(t *testing.T) {
		dbHost, ok := config.GetSetting("dbHost")
		if !ok || dbHost != "db.example.go.com" {
			t.Errorf("Expected dbHost 'db.example.go.com', got '%v' (found: %t)", dbHost, ok)
		}
	})

	t.Run("should set and get values correctly", func(t *testing.T) {
		config.SetSetting("timeout", 5000)
		timeout, ok := config.GetSetting("timeout")
		if !ok || timeout != 5000 {
			t.Errorf("Expected timeout 5000, got '%v' (found: %t)", timeout, ok)
		}
	})

	t.Run("GetAllSettings should return a copy", func(t *testing.T) {
		settings1 := config.GetAllSettings()
		settings1["newKey"] = "newValue" // Modify the returned map

		settings2 := config.GetAllSettings()
		_, exists := settings2["newKey"]
		if exists {
			t.Error("Modification of returned map affected the internal configuration")
		}
	})
}

// Example of how you might run the main function as part of a test
// (though typically main logic isn't tested this way directly)
func ExampleMain() {
	// Redirect stdout for testing if needed
	fmt.Println("Running ExampleMain (simulates main execution)")
	// main() // Cannot call main directly in test package
	// Output: Running ExampleMain (simulates main execution)
}
