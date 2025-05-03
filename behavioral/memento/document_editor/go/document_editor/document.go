// behavioral/memento/document_editor/go/document_editor/document.go
package document_editor

import (
	"fmt"
)

// Document represents the Originator which holds the state.
// It can create Mementos to save its state and restore from them.
type Document struct {
	content string
}

// NewDocument creates a new Document with optional initial content.
func NewDocument(initialContent string) *Document {
	d := &Document{content: initialContent}
	fmt.Printf("Document initialized with: '%s'\n", d.content)
	return d
}

// Write appends text to the document's content.
func (d *Document) Write(text string) {
	fmt.Printf("Writing: '%s'\n", text)
	d.content += text
	fmt.Printf("Current content: '%s'\n", d.content)
}

// SetContent directly sets the document's content (used for restoring).
func (d *Document) SetContent(content string) {
    d.content = content
    fmt.Printf("Content set to: '%s'\n", d.content)
}

// GetContent returns the current content of the document.
func (d *Document) GetContent() string {
	return d.content
}

// Save creates a Memento containing the current state of the document.
func (d *Document) Save() IMemento {
	fmt.Printf("Saving state: '%s'\n", d.content)
	// We return the interface type, but create the concrete implementation
	return newConcreteMemento(d.content)
}

// Restore sets the document's state from a given Memento.
func (d *Document) Restore(m IMemento) {
	d.content = m.GetState()
	fmt.Printf("Restoring state to: '%s'\n", d.content)
}