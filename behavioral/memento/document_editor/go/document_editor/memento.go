// behavioral/memento/document_editor/go/document_editor/memento.go
package document_editor

import (
	"fmt"
	"time"
)

// Memento interface defines the method for retrieving the saved state.
// It doesn't expose the Originator's state directly.
type IMemento interface {
	GetState() string
	GetName() string // Typically used for metadata like timestamp
	GetDate() time.Time
}

// concreteMemento is the internal implementation of the Memento.
// It stores the state of the Originator.
type concreteMemento struct {
	state string
	date  time.Time
}

// newConcreteMemento creates a new Memento.
// This function is usually package-private or used by the Originator.
func newConcreteMemento(state string) *concreteMemento {
	return &concreteMemento{
		state: state,
		date:  time.Now(),
	}
}

// GetState returns the saved state.
func (m *concreteMemento) GetState() string {
	return m.state
}

// GetName returns metadata about the memento (e.g., timestamp).
func (m *concreteMemento) GetName() string {
	return fmt.Sprintf("%s / (%s...)", m.date.Format(time.RFC3339), m.state[:min(10, len(m.state))])
}

// GetDate returns the creation date of the memento.
func (m *concreteMemento) GetDate() time.Time {
	return m.date
}

// Helper function (not strictly needed for pattern, utility)
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}