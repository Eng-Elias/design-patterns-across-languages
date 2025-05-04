package main

import (
	"fmt"
	"time"

	// Import the local package using the module path
	event_monitoring "observer_pattern_event_monitoring_go/event_monitoring"
)

func main() {
	fmt.Println("--- Go Observer Pattern Demo ---")

	// Create EventSource (Subject)
	eventSource := event_monitoring.NewEventSource()

	// Create Observers
	logger := event_monitoring.NewLoggerObserver()
	notifier := event_monitoring.NewNotifierObserver()

	// Attach Observers
	fmt.Println("\nAttaching observers...")
	eventSource.Attach(logger)
	eventSource.Attach(notifier)

	fmt.Println("\nGenerating specific events...")
	eventSource.GenerateEvent(event_monitoring.LogInfo, map[string]interface{}{"msg": "User logged in"})
	time.Sleep(50 * time.Millisecond) // Small delay
	eventSource.GenerateEvent(event_monitoring.LogWarn, map[string]interface{}{"msg": "Disk space low"})
	time.Sleep(50 * time.Millisecond)
	eventSource.GenerateEvent(event_monitoring.LogError, map[string]interface{}{"code": 500, "error": "Database connection failed"})
	time.Sleep(50 * time.Millisecond)
	eventSource.GenerateEvent(event_monitoring.LogCritical, map[string]interface{}{"code": 999, "error": "System meltdown imminent"})
	time.Sleep(50 * time.Millisecond)

	// Detach an observer
	fmt.Println("\nDetaching LoggerObserver...")
	eventSource.Detach(logger)
	time.Sleep(50 * time.Millisecond)

	fmt.Println("\nGenerating another event...")
	eventSource.GenerateEvent(event_monitoring.LogInfo, map[string]interface{}{"msg": "User logged out"})
	time.Sleep(50 * time.Millisecond)

	fmt.Println("\n--- Demo Finished ---")
}