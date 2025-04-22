package main

import (
	"fmt"
	"singleton_pattern_configuration_manager_go/configuration_manager"
	"sync"
)

func main() {
	fmt.Println("--- Singleton Pattern - Configuration Manager Demo (Go) ---")

	// Get the singleton instance
	fmt.Println("\nGetting Configuration Manager instance 1...")
	configManager1 := configuration_manager.GetInstance()

	// Access configuration
dbHost, _ := configManager1.GetSetting("dbHost")
	fmt.Printf("DB Host (Instance 1): %v\n", dbHost)

	// Modify configuration through the first instance
	fmt.Println("\nSetting 'connectionPoolSize' via Instance 1...")
	configManager1.SetSetting("connectionPoolSize", 10)
	poolSize, _ := configManager1.GetSetting("connectionPoolSize")
	fmt.Printf("Connection Pool Size (Instance 1): %v\n", poolSize)

	// Get the singleton instance again (possibly concurrently)
	fmt.Println("\nGetting Configuration Manager instance 2 (concurrently)...")
	var configManager2 *configuration_manager.ConfigurationManager
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		configManager2 = configuration_manager.GetInstance()
	}()
	wg.Wait() // Wait for the goroutine to finish

	// Verify it's the same instance
	fmt.Printf("\nInstance 1 and Instance 2 are the same: %t\n", configManager1 == configManager2)

	// Access the modified configuration through the second instance
poolSize2, _ := configManager2.GetSetting("connectionPoolSize")
	fmt.Printf("Connection Pool Size (Instance 2): %v\n", poolSize2)
dbHost2, _ := configManager2.GetSetting("dbHost")
	fmt.Printf("DB Host (Instance 2): %v\n", dbHost2)

	fmt.Println("\nAll settings:")
	allSettings := configManager2.GetAllSettings()
	for key, value := range allSettings {
		fmt.Printf("  %s: %v\n", key, value)
	}
}
