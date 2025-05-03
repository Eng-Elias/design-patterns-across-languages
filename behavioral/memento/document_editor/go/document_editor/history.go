// behavioral/memento/document_editor/go/document_editor/history.go
package document_editor

import "fmt"

// History acts as the Caretaker. It holds the mementos.
// It doesn't know the internal details of the mementos or the originator.
type History struct {
	history   []IMemento // Stack for undo
	redoStack []IMemento // Stack for redo
	document  *Document  // Reference to the Originator
}

// NewHistory creates a new History manager for a Document.
func NewHistory(doc *Document) *History {
	h := &History{
		document: doc,
		// Initialize with empty slices, not nil, to avoid nil checks later
		history:   make([]IMemento, 0),
		redoStack: make([]IMemento, 0),
	}
	// Save the initial state of the document
	h.Save()
	return h
}

// Save captures the current state of the document and stores it.
func (h *History) Save() {
	fmt.Println("History: Saving state...")
	memento := h.document.Save() // Get memento from originator
	h.history = append(h.history, memento)
	// Clear the redo stack whenever a new state is saved directly
	h.redoStack = make([]IMemento, 0)
	fmt.Println("History: State saved.")
}

// Undo restores the document to its previous state.
func (h *History) Undo() {
	if len(h.history) <= 1 { // Need at least initial state + one more
		fmt.Println("History: Cannot undo further.")
		return
	}

	fmt.Println("History: Undoing...")
	// Pop the current state from history
	lastIndex := len(h.history) - 1
	currentMemento := h.history[lastIndex]
	h.history = h.history[:lastIndex]

	// Push the popped state onto the redo stack
	h.redoStack = append(h.redoStack, currentMemento)

	// Restore the document to the previous state (now the last in history)
	previousMemento := h.history[len(h.history)-1]
	h.document.Restore(previousMemento)
	fmt.Println("History: Undo complete.")
}

// Redo restores the document to a previously undone state.
func (h *History) Redo() {
	if len(h.redoStack) == 0 {
		fmt.Println("History: Cannot redo.")
		return
	}

	fmt.Println("History: Redoing...")
	// Pop the state to redo from the redo stack
	lastIndex := len(h.redoStack) - 1
	mementoToRestore := h.redoStack[lastIndex]
	h.redoStack = h.redoStack[:lastIndex]

	// Restore the document state
	h.document.Restore(mementoToRestore)

	// Push the restored state back onto the main history stack
	h.history = append(h.history, mementoToRestore)
	fmt.Println("History: Redo complete.")
}

// PrintHistory displays the states stored in the history and redo stacks.
func (h *History) PrintHistory() {
	fmt.Println("--- History Log ---")
	for i, memento := range h.history {
		fmt.Printf("  %d: '%s' (%s)\n", i, memento.GetState(), memento.GetName())
	}
	fmt.Println("--- Redo Stack Log ---")
	for i, memento := range h.redoStack {
		fmt.Printf("  %d: '%s' (%s)\n", i, memento.GetState(), memento.GetName())
	}
	fmt.Println("--------------------")
}

// --- Helper methods for testing ---

// GetHistoryLength returns the number of states in the undo history.
func (h *History) GetHistoryLength() int {
	return len(h.history)
}

// GetRedoStackLength returns the number of states in the redo stack.
func (h *History) GetRedoStackLength() int {
	return len(h.redoStack)
}

// GetLastHistoryState returns the state string of the last memento in history.
// Returns empty string and false if history is empty.
func (h *History) GetLastHistoryState() (string, bool) {
	if len(h.history) == 0 {
		return "", false
	}
	return h.history[len(h.history)-1].GetState(), true
}

// GetLastRedoState returns the state string of the last memento in the redo stack.
// Returns empty string and false if redo stack is empty.
func (h *History) GetLastRedoState() (string, bool) {
	if len(h.redoStack) == 0 {
		return "", false
	}
	return h.redoStack[len(h.redoStack)-1].GetState(), true
}