// behavioral/memento/document_editor/go/main.go
package main

import (
	"fmt"
	// Use the module path defined in go.mod
	"memento_pattern_document_editor_go/document_editor"
)

func main() {
	fmt.Println("--- Memento Pattern Document Editor Demo (Go) ---")

	// Create Originator (Document)
	doc := document_editor.NewDocument("")

	// Create Caretaker (History), passing the originator
	history := document_editor.NewHistory(doc)

	// User starts typing and saving states
	doc.Write("Hello")
	history.Save()
	history.PrintHistory()

	doc.Write(" World")
	history.Save()
	history.PrintHistory()

	doc.Write("!")
	history.Save()
	history.PrintHistory()

	// Undo operations
	fmt.Println("\n--- Undoing ---")
	history.Undo() // Undo "!"
	history.PrintHistory()

	history.Undo() // Undo " World"
	history.PrintHistory()

	// Redo operation
	fmt.Println("\n--- Redoing ---")
	history.Redo() // Redo " World"
	history.PrintHistory()

	// Make another change
	fmt.Println("\n--- Making another change ---")
	doc.Write(", How are you?")
	history.Save()
	history.PrintHistory()

	// Try to redo (should fail as redo stack was cleared)
	fmt.Println("\n--- Trying Redo again ---")
	history.Redo()
	history.PrintHistory()

	// Undo back to the beginning
	fmt.Println("\n--- Undoing to the start ---")
	history.Undo() // Undo ", How are you?"
	history.Undo() // Undo " World"
	history.Undo() // Undo "Hello"
	history.PrintHistory()

	// Try to undo further (should fail)
	fmt.Println("\n--- Trying Undo again ---")
	history.Undo()
	history.PrintHistory()

	fmt.Println("\n--- Demo Complete ---")
}