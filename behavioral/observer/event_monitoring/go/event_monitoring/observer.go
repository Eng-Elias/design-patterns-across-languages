package event_monitoring

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

// IObserver defines the interface for observers
type IObserver interface {
	Update(event IEvent)
}

// --- LoggerObserver ---

// LoggerObserver logs events
type LoggerObserver struct {
	mu   sync.Mutex
	logs []string
}

// NewLoggerObserver creates a new LoggerObserver
func NewLoggerObserver() *LoggerObserver {
	return &LoggerObserver{
		logs: make([]string, 0),
	}
}

// Update logs the event details
func (lo *LoggerObserver) Update(event IEvent) {
	lo.mu.Lock()
	defer lo.mu.Unlock()

	dataBytes, err := json.Marshal(event.GetData())
	if err != nil {
		// Log internal error, but format a message for the demo log anyway
		log.Printf("LoggerObserver internal error marshalling data: %v", err)
		dataBytes = []byte("{'error': 'marshalling failed'}")
	}

	// Format log message
	// Using Printf for direct output without standard log prefixes for cleaner demo output
	logMsg := fmt.Sprintf("%s: %s at %s", event.GetType(), string(dataBytes), event.GetTimestamp().Format("2006-01-02 15:04:05"))
	fmt.Printf("LoggerObserver: Received event - %s\n", logMsg) // Use fmt.Printf for demo
	lo.logs = append(lo.logs, logMsg) // Still store for testing GetLogs
}

// GetLogs returns the recorded logs (for testing)
func (lo *LoggerObserver) GetLogs() []string {
	lo.mu.Lock()
	defer lo.mu.Unlock()
	// Return a copy to prevent external modification
	logsCopy := make([]string, len(lo.logs))
	copy(logsCopy, lo.logs)
	return logsCopy
}

// --- NotifierObserver ---

// NotifierObserver sends notifications for specific event types
type NotifierObserver struct {
	mu           sync.Mutex
	notifications []string
}

// NewNotifierObserver creates a new NotifierObserver
func NewNotifierObserver() *NotifierObserver {
	return &NotifierObserver{
		notifications: make([]string, 0),
	}
}

// Update sends a notification if the event is ERROR or CRITICAL
func (no *NotifierObserver) Update(event IEvent) {
	no.mu.Lock()
	defer no.mu.Unlock()

	if event.GetType() == LogError || event.GetType() == LogCritical {
		notificationMsg := fmt.Sprintf("Notification: Critical event - Type: %s, Data: %v, Time: %s",
			event.GetType(),
			event.GetData(), // Keep simple map representation for demo
			event.GetTimestamp().Format("2006-01-02 15:04:05"))
		fmt.Printf("NotifierObserver: %s\n", notificationMsg) // Use fmt.Printf for demo
		no.notifications = append(no.notifications, notificationMsg) // Still store for testing
	}
}

// GetNotifications returns the recorded notifications (for testing)
func (no *NotifierObserver) GetNotifications() []string {
	no.mu.Lock()
	defer no.mu.Unlock()
	// Return a copy
	notificationsCopy := make([]string, len(no.notifications))
	copy(notificationsCopy, no.notifications)
	return notificationsCopy
}
