package event_monitoring

import (
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"
)

// MockObserver for testing purposes
type MockObserver struct {
	UpdateFunc func(event IEvent)
	ID         string // To differentiate mock observers if needed
}

func (mo *MockObserver) Update(event IEvent) {
	if mo.UpdateFunc != nil {
		mo.UpdateFunc(event)
	}
}

func TestEventSource_AttachDetach(t *testing.T) {
	eventSource := NewEventSource()
	observer1 := &MockObserver{ID: "mock1"}
	observer2 := &MockObserver{ID: "mock2"}

	if len(eventSource.observers) != 0 {
		t.Errorf("Expected initial observer count to be 0, got %d", len(eventSource.observers))
	}

	eventSource.Attach(observer1)
	if len(eventSource.observers) != 1 {
		t.Errorf("Expected observer count to be 1 after attaching, got %d", len(eventSource.observers))
	}

	eventSource.Attach(observer2)
	if len(eventSource.observers) != 2 {
		t.Errorf("Expected observer count to be 2 after attaching another, got %d", len(eventSource.observers))
	}

	// Test attaching the same observer again (should not increase count)
	eventSource.Attach(observer1)
	if len(eventSource.observers) != 2 {
		t.Errorf("Expected observer count to remain 2 after re-attaching, got %d", len(eventSource.observers))
	}

	eventSource.Detach(observer1)
	if len(eventSource.observers) != 1 {
		t.Errorf("Expected observer count to be 1 after detaching, got %d", len(eventSource.observers))
	}
	// Check if the correct observer is remaining
	if _, exists := eventSource.observers[observer2]; !exists {
		t.Errorf("Expected observer2 to remain after detaching observer1")
	}

	eventSource.Detach(observer2)
	if len(eventSource.observers) != 0 {
		t.Errorf("Expected observer count to be 0 after detaching all, got %d", len(eventSource.observers))
	}

	// Test detaching a non-existent observer
	nonExistentObserver := &MockObserver{ID: "nonexistent"}
	eventSource.Detach(nonExistentObserver)
	if len(eventSource.observers) != 0 {
		t.Errorf("Expected observer count to remain 0 after detaching non-existent observer, got %d", len(eventSource.observers))
	}
}

func TestEventSource_Notify(t *testing.T) {
	eventSource := NewEventSource()
	var wg sync.WaitGroup
	update1Called := false
	update2Called := false

	observer1 := &MockObserver{
		ID: "mock1",
		UpdateFunc: func(event IEvent) {
			update1Called = true
			wg.Done()
		},
	}
	observer2 := &MockObserver{
		ID: "mock2",
		UpdateFunc: func(event IEvent) {
			update2Called = true
			wg.Done()
		},
	}

	eventSource.Attach(observer1)
	eventSource.Attach(observer2)

	testEvent := NewEvent(LogInfo, map[string]interface{}{"test": "data"})

	wg.Add(2) // Expect two updates
	eventSource.Notify(testEvent)

	// Wait for observers to finish processing or timeout
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// Proceed with assertions
	case <-time.After(1 * time.Second): // Timeout to prevent test hanging
		t.Fatal("Timeout waiting for observers to update")
	}

	if !update1Called {
		t.Errorf("Expected observer1's Update method to be called")
	}
	if !update2Called {
		t.Errorf("Expected observer2's Update method to be called")
	}
}

func TestLoggerObserver(t *testing.T) {
	logger := NewLoggerObserver()
	event1 := NewEvent(LogInfo, map[string]interface{}{"msg": "Info 1"})
	event2 := NewEvent(LogWarn, map[string]interface{}{"msg": "Warn 1"})

	logger.Update(event1)
	logger.Update(event2)

	logs := logger.GetLogs()
	if len(logs) != 2 {
		t.Fatalf("Expected 2 log entries, got %d", len(logs))
	}

	// Note: Timestamp makes exact match difficult, check for essential parts
	if !strings.Contains(logs[0], `LOG_INFO: {"msg":"Info 1"}`) {
		t.Errorf("Log entry 1 does not contain expected info data: %s", logs[0])
	}
	if !strings.Contains(logs[1], `LOG_WARN: {"msg":"Warn 1"}`) {
		t.Errorf("Log entry 2 does not contain expected warn data: %s", logs[1])
	}
}

func TestNotifierObserver(t *testing.T) {
	notifier := NewNotifierObserver()
	eventInfo := NewEvent(LogInfo, map[string]interface{}{"msg": "Just info"})
	eventWarn := NewEvent(LogWarn, map[string]interface{}{"msg": "Just warning"})
	eventError := NewEvent(LogError, map[string]interface{}{"code": 500})
	eventCritical := NewEvent(LogCritical, map[string]interface{}{"code": 999})

	notifier.Update(eventInfo)
	notifier.Update(eventWarn)
	notifier.Update(eventError)
	notifier.Update(eventCritical)

	notifications := notifier.GetNotifications()
	if len(notifications) != 2 {
		t.Fatalf("Expected 2 notification entries, got %d", len(notifications))
	}

	// Check if notifications contain the correct event types
	if !strings.Contains(notifications[0], "Type: LOG_ERROR") {
		t.Errorf("Notification 1 should be for LOG_ERROR, got: %s", notifications[0])
	}
	if !strings.Contains(notifications[1], "Type: LOG_CRITICAL") {
		t.Errorf("Notification 2 should be for LOG_CRITICAL, got: %s", notifications[1])
	}
}

func TestEventSource_GenerateEvent(t *testing.T) {
	eventSource := NewEventSource()
	var wg sync.WaitGroup
	updateCalled := false
	var receivedEvent IEvent

	observer := &MockObserver{
		ID: "mock",
		UpdateFunc: func(event IEvent) {
			updateCalled = true
			receivedEvent = event
			wg.Done()
		},
	}

	eventSource.Attach(observer)

	dataType := LogInfo
	dataMap := map[string]interface{}{"generated": true}

	wg.Add(1)
	eventSource.GenerateEvent(dataType, dataMap)

	// Wait or timeout
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// Proceed
	case <-time.After(1 * time.Second):
		t.Fatal("Timeout waiting for GenerateEvent notification")
	}

	if !updateCalled {
		t.Fatal("GenerateEvent did not trigger observer Update")
	}
	if receivedEvent == nil {
		t.Fatal("Received event is nil")
	}
	if receivedEvent.GetType() != dataType {
		t.Errorf("Expected event type %s, got %s", dataType, receivedEvent.GetType())
	}
	if fmt.Sprintf("%v", receivedEvent.GetData()) != fmt.Sprintf("%v", dataMap) {
		t.Errorf("Expected event data %v, got %v", dataMap, receivedEvent.GetData())
	}
}
