package event_monitoring

import "time"

// EventType defines the type of event
type EventType string

const (
	LogInfo     EventType = "LOG_INFO"
	LogWarn     EventType = "LOG_WARN"
	LogError    EventType = "LOG_ERROR"
	LogCritical EventType = "LOG_CRITICAL"
)

// IEvent defines the interface for events
type IEvent interface {
	GetType() EventType
	GetData() map[string]interface{}
	GetTimestamp() time.Time
}

// Event represents a specific event that occurred
type Event struct {
	EventType EventType
	Data      map[string]interface{}
	Timestamp time.Time
}

// NewEvent creates a new Event
func NewEvent(eventType EventType, data map[string]interface{}) *Event {
	return &Event{
		EventType: eventType,
		Data:      data,
		Timestamp: time.Now(),
	}
}

// GetType returns the event type
func (e *Event) GetType() EventType {
	return e.EventType
}

// GetData returns the event data
func (e *Event) GetData() map[string]interface{} {
	return e.Data
}

// GetTimestamp returns the event timestamp
func (e *Event) GetTimestamp() time.Time {
	return e.Timestamp
}
