package main

import (
	la "adapter_pattern_logging_adapter_go/logging_adapter" // Alias for clarity
	"fmt"
)

func main() {
	fmt.Println("--- Using the Adapter for the Third-Party Logger ---")

	// Create the Adaptee (the incompatible logger)
	thirdPartyLogger := &la.ThirdPartyLogger{}

	// Create the Adapter, wrapping the Adaptee
	// Ensure the adapter implements the Logger interface
	var logger la.Logger = la.NewLoggerAdapter(thirdPartyLogger)

	// The client code (ApplicationService) uses the standard Logger interface
	// It doesn't know it's talking to an adapter or a third-party logger.
	appService := la.NewApplicationService(logger)

	fmt.Println("\nPerforming operations:")
	appService.PerformOperation("ImportantData123")
	fmt.Println("---")
	appService.PerformOperation("abc") // Should trigger a warning
	fmt.Println("---")
	appService.PerformOperation("")    // Should trigger an error
}
