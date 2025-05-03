# behavioral/memento/document_editor/python/history.py
from typing import List
from memento import Memento
from document import Document

class History:
    """
    The Caretaker doesn't depend on the Concrete Memento class. Therefore, it
    doesn't have access to the originator's state, stored inside the memento.
    It works with all mementos via the base Memento interface.
    """
    def __init__(self, document: Document):
        self._document = document
        self._history: List[Memento] = []
        self._redo_stack: List[Memento] = []
        # Save initial state
        self.save()

    def save(self):
        """Saves the current document state to history."""
        print("History: Saving state...")
        memento = self._document.save()
        self._history.append(memento)
        # Clear redo stack whenever a new state is saved
        self._redo_stack.clear()
        print("History: State saved.")

    def undo(self):
        """Restores the previous state."""
        if len(self._history) <= 1:  # Need at least initial state + one more
            print("History: Cannot undo further.")
            return

        print("History: Undoing...")
        # Pop current state and move it to redo stack
        current_memento = self._history.pop()
        self._redo_stack.append(current_memento)

        # Restore to the previous state (now the last in history)
        previous_memento = self._history[-1]
        self._document.restore(previous_memento)
        print("History: Undo complete.")

    def redo(self):
        """Restores a previously undone state."""
        if not self._redo_stack:
            print("History: Cannot redo.")
            return

        print("History: Redoing...")
        # Pop the state to redo from the redo stack
        memento_to_restore = self._redo_stack.pop()

        # Restore the document state
        self._document.restore(memento_to_restore)

        # Add the restored state back to the main history
        self._history.append(memento_to_restore)
        print("History: Redo complete.")

    def print_history(self):
        """Prints the states stored in history (for demo)."""
        print("History Log:")
        for i, memento in enumerate(self._history):
            print(f"  {i}: '{memento.get_state()}'")
        print("Redo Stack Log:")
        for i, memento in enumerate(self._redo_stack):
            print(f"  {i}: '{memento.get_state()}'")
