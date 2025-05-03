# behavioral/memento/document_editor/python/main.py
from document import Document
from history import History

def main():
    print("--- Memento Pattern Document Editor Demo ---")

    # Create Originator (Document)
    doc = Document()

    # Create Caretaker (History), passing the originator
    history = History(doc)

    # User starts typing and saving states
    doc.write("Hello")
    history.save()
    history.print_history()

    doc.write(" World")
    history.save()
    history.print_history()

    doc.write("!")
    history.save()
    history.print_history()

    # Undo operations
    print("\n--- Undoing ---")
    history.undo() # Undo "!"
    history.print_history()

    history.undo() # Undo " World"
    history.print_history()

    # Redo operation
    print("\n--- Redoing ---")
    history.redo() # Redo " World"
    history.print_history()

    # Make another change
    print("\n--- Making another change ---")
    doc.write(", How are you?")
    history.save()
    history.print_history()

    # Try to redo (should fail as redo stack was cleared)
    print("\n--- Trying Redo again ---")
    history.redo()
    history.print_history()

    # Undo back to the beginning
    print("\n--- Undoing to the start ---")
    history.undo() # Undo ", How are you?"
    history.undo() # Undo " World"
    history.undo() # Undo "Hello"
    history.print_history()

    # Try to undo further (should fail)
    print("\n--- Trying Undo again ---")
    history.undo()
    history.print_history()

    print("\n--- Demo Complete ---")

if __name__ == "__main__":
    main()
