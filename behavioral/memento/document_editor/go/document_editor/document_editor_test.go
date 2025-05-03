// behavioral/memento/document_editor/go/document_editor/document_editor_test.go
package document_editor

import (
	"testing"
	"time"
)

// TestDocumentWrite checks if writing to the document updates its content.
func TestDocumentWrite(t *testing.T) {
	doc := NewDocument("")
	doc.Write("Initial content")
	if doc.GetContent() != "Initial content" {
		t.Errorf("Expected 'Initial content', got '%s'", doc.GetContent())
	}
	doc.Write(" additional")
	if doc.GetContent() != "Initial content additional" {
		t.Errorf("Expected 'Initial content additional', got '%s'", doc.GetContent())
	}
}

// TestMementoState checks if the Memento correctly stores and retrieves the state.
func TestMementoState(t *testing.T) {
	doc := NewDocument("Test state")
	memento := doc.Save()
	if memento.GetState() != "Test state" {
		t.Errorf("Expected Memento state 'Test state', got '%s'", memento.GetState())
	}
	// Check if the name format is reasonable (e.g., contains date)
	expectedDate := time.Now().Format("2006-01-02")
	if len(memento.GetName()) < len(expectedDate) || memento.GetName()[:len(expectedDate)] != expectedDate {
		t.Errorf("Memento name '%s' does not seem to start with the expected date format '%s'", memento.GetName(), expectedDate)
	}
}

// TestDocumentRestore checks if the document can be restored from a memento.
func TestDocumentRestore(t *testing.T) {
	doc := NewDocument("State 1")
	memento1 := doc.Save()

	doc.Write(" plus State 2")
	memento2 := doc.Save() // Save State 2

	doc.Write(" and State 3") // Current is State 3

	doc.Restore(memento2) // Restore to State 2
	if doc.GetContent() != "State 1 plus State 2" {
		t.Errorf("Expected restored state 'State 1 plus State 2', got '%s'", doc.GetContent())
	}

	doc.Restore(memento1) // Restore to State 1
	if doc.GetContent() != "State 1" {
		t.Errorf("Expected restored state 'State 1', got '%s'", doc.GetContent())
	}
}

// TestHistorySave checks if saving state adds mementos to history.
func TestHistorySave(t *testing.T) {
	doc := NewDocument("")
	history := NewHistory(doc) // Saves initial state

	if history.GetHistoryLength() != 1 {
		t.Errorf("Expected initial history length 1, got %d", history.GetHistoryLength())
	}

	doc.Write("First change")
	history.Save()
	if history.GetHistoryLength() != 2 {
		t.Errorf("Expected history length 2 after save, got %d", history.GetHistoryLength())
	}

	doc.Write(" Second change")
	history.Save()
	if history.GetHistoryLength() != 3 {
		t.Errorf("Expected history length 3 after save, got %d", history.GetHistoryLength())
	}
}

// TestHistoryUndo checks the undo functionality.
func TestHistoryUndo(t *testing.T) {
	doc := NewDocument("A")
	history := NewHistory(doc) // History: [Memento(A)]

	doc.Write("B")
	history.Save() // History: [Memento(A), Memento(AB)], Redo: []

	doc.Write("C")
	history.Save() // History: [Memento(A), Memento(AB), Memento(ABC)], Redo: []

	if doc.GetContent() != "ABC" {
		t.Errorf("Expected content 'ABC' before undo, got '%s'", doc.GetContent())
	}
	if history.GetHistoryLength() != 3 {
		t.Errorf("Expected history length 3 before undo, got %d", history.GetHistoryLength())
	}
	if history.GetRedoStackLength() != 0 {
		t.Errorf("Expected redo stack length 0 before undo, got %d", history.GetRedoStackLength())
	}

	history.Undo() // Restore AB. History: [Memento(A), Memento(AB)], Redo: [Memento(ABC)]
	if doc.GetContent() != "AB" {
		t.Errorf("Expected content 'AB' after first undo, got '%s'", doc.GetContent())
	}
	if history.GetHistoryLength() != 2 {
		t.Errorf("Expected history length 2 after first undo, got %d", history.GetHistoryLength())
	}
	if history.GetRedoStackLength() != 1 {
		t.Errorf("Expected redo stack length 1 after first undo, got %d", history.GetRedoStackLength())
	}
	lastRedoState, ok := history.GetLastRedoState()
	if !ok || lastRedoState != "ABC" {
		t.Errorf("Expected last redo state 'ABC', got '%s' (ok: %v)", lastRedoState, ok)
	}

	history.Undo() // Restore A. History: [Memento(A)], Redo: [Memento(ABC), Memento(AB)]
	if doc.GetContent() != "A" {
		t.Errorf("Expected content 'A' after second undo, got '%s'", doc.GetContent())
	}
	if history.GetHistoryLength() != 1 {
		t.Errorf("Expected history length 1 after second undo, got %d", history.GetHistoryLength())
	}
	if history.GetRedoStackLength() != 2 {
		t.Errorf("Expected redo stack length 2 after second undo, got %d", history.GetRedoStackLength())
	}

	// Try undoing further (should do nothing)
	history.Undo()
	if doc.GetContent() != "A" {
		t.Errorf("Expected content 'A' after third undo attempt, got '%s'", doc.GetContent())
	}
	if history.GetHistoryLength() != 1 {
		t.Errorf("Expected history length 1 after third undo attempt, got %d", history.GetHistoryLength())
	}
	if history.GetRedoStackLength() != 2 {
		t.Errorf("Expected redo stack length 2 after third undo attempt, got %d", history.GetRedoStackLength())
	}
}

// TestHistoryRedo checks the redo functionality.
func TestHistoryRedo(t *testing.T) {
	doc := NewDocument("1")
	history := NewHistory(doc) // History: [Memento(1)], Redo: []

	doc.Write("2")
	history.Save() // History: [Memento(1), Memento(12)], Redo: []

	doc.Write("3")
	history.Save() // History: [Memento(1), Memento(12), Memento(123)], Redo: []

	history.Undo() // Restore 12. History: [Memento(1), Memento(12)], Redo: [Memento(123)]
	history.Undo() // Restore 1. History: [Memento(1)], Redo: [Memento(123), Memento(12)]

	if doc.GetContent() != "1" {
		t.Errorf("Expected content '1' before redo, got '%s'", doc.GetContent())
	}
	if history.GetHistoryLength() != 1 {
		t.Errorf("Expected history length 1 before redo, got %d", history.GetHistoryLength())
	}
	if history.GetRedoStackLength() != 2 {
		t.Errorf("Expected redo stack length 2 before redo, got %d", history.GetRedoStackLength())
	}

	history.Redo() // Restore 12. History: [Memento(1), Memento(12)], Redo: [Memento(123)]
	if doc.GetContent() != "12" {
		t.Errorf("Expected content '12' after first redo, got '%s'", doc.GetContent())
	}
	if history.GetHistoryLength() != 2 {
		t.Errorf("Expected history length 2 after first redo, got %d", history.GetHistoryLength())
	}
	if history.GetRedoStackLength() != 1 {
		t.Errorf("Expected redo stack length 1 after first redo, got %d", history.GetRedoStackLength())
	}

	history.Redo() // Restore 123. History: [Memento(1), Memento(12), Memento(123)], Redo: []
	if doc.GetContent() != "123" {
		t.Errorf("Expected content '123' after second redo, got '%s'", doc.GetContent())
	}
	if history.GetHistoryLength() != 3 {
		t.Errorf("Expected history length 3 after second redo, got %d", history.GetHistoryLength())
	}
	if history.GetRedoStackLength() != 0 {
		t.Errorf("Expected redo stack length 0 after second redo, got %d", history.GetRedoStackLength())
	}

	// Try redoing further (should do nothing)
	history.Redo()
	if doc.GetContent() != "123" {
		t.Errorf("Expected content '123' after third redo attempt, got '%s'", doc.GetContent())
	}
	if history.GetHistoryLength() != 3 {
		t.Errorf("Expected history length 3 after third redo attempt, got %d", history.GetHistoryLength())
	}
	if history.GetRedoStackLength() != 0 {
		t.Errorf("Expected redo stack length 0 after third redo attempt, got %d", history.GetRedoStackLength())
	}
}

// TestHistoryRedoClear checks if redo stack is cleared on new save.
func TestHistoryRedoClear(t *testing.T) {
	doc := NewDocument("X")
	history := NewHistory(doc) // History: [Memento(X)], Redo: []

	doc.Write("Y")
	history.Save() // History: [Memento(X), Memento(XY)], Redo: []

	doc.Write("Z")
	history.Save() // History: [Memento(X), Memento(XY), Memento(XYZ)], Redo: []

	history.Undo() // Restore XY. History: [Memento(X), Memento(XY)], Redo: [Memento(XYZ)]
	if history.GetRedoStackLength() != 1 {
		t.Errorf("Expected redo stack length 1 after undo, got %d", history.GetRedoStackLength())
	}

	// Make a new change instead of redoing
	doc.Write("W") // Content is now XYW
	history.Save() // History: [Memento(X), Memento(XY), Memento(XYW)], Redo: [] (should be cleared)

	if doc.GetContent() != "XYW" {
		t.Errorf("Expected content 'XYW' after new save, got '%s'", doc.GetContent())
	}
	if history.GetHistoryLength() != 3 {
		t.Errorf("Expected history length 3 after new save, got %d", history.GetHistoryLength())
	}
	if history.GetRedoStackLength() != 0 {
		t.Errorf("Expected redo stack to be cleared after new save, got length %d", history.GetRedoStackLength())
	}

	// Try redoing (should do nothing as stack is clear)
	history.Redo()
	if doc.GetContent() != "XYW" {
		t.Errorf("Expected content 'XYW' after redo attempt, got '%s'", doc.GetContent())
	}
}