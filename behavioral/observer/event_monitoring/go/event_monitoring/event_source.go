package event_monitoring

import (
	"log"
	"sync"
)

// EventSource manages observers and notifies them of events.
type EventSource struct {
	observers map[IObserver]struct{} // Using a map as a set for efficient add/remove
	mu        sync.RWMutex
}

// NewEventSource creates a new EventSource.
func NewEventSource() *EventSource {
	return &EventSource{
		observers: make(map[IObserver]struct{}),
	}
}

// Attach adds an observer to the EventSource.
func (es *EventSource) Attach(observer IObserver) {
	es.mu.Lock()
	defer es.mu.Unlock()
	es.observers[observer] = struct{}{}
	log.Printf("EventSource: Attached observer %T", observer)
}

// Detach removes an observer from the EventSource.
func (es *EventSource) Detach(observer IObserver) {
	es.mu.Lock()
	defer es.mu.Unlock()
	delete(es.observers, observer)
	log.Printf("EventSource: Detached observer %T", observer)
}

// Notify sends an event to all attached observers.
func (es *EventSource) Notify(event IEvent) {
	es.mu.RLock()
	defer es.mu.RUnlock()

	log.Printf("EventSource: Notifying %d observers about event type %s", len(es.observers), event.GetType())
	for observer := range es.observers {
		// Run updates in separate goroutines for potential concurrency
		// Note: In a real-world scenario, consider error handling and worker pools
		go observer.Update(event)
	}
}

// GenerateEvent simulates creating an event and notifying observers.
// This is a helper method for demonstration.
func (es *EventSource) GenerateEvent(eventType EventType, data map[string]interface{}) {
	event := NewEvent(eventType, data)
	log.Printf("EventSource: Generating event %s with data %v", eventType, data)
	es.Notify(event)
}
