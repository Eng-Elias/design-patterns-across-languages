# behavioral/memento/document_editor/python/test_document_editor.py
import unittest
from document import Document
from history import History
from memento import Memento # Needed for type checking if desired

class TestDocumentEditorMemento(unittest.TestCase):

    def test_initial_state(self):
        doc = Document("Initial")
        history = History(doc)
        self.assertEqual(doc.get_content(), "Initial")
        # History should have saved the initial state
        self.assertEqual(len(history._history), 1)
        self.assertEqual(history._history[0].get_state(), "Initial")

    def test_write_and_save(self):
        doc = Document()
        history = History(doc) # Saves initial "" state
        doc.write("First edit")
        history.save()
        self.assertEqual(doc.get_content(), "First edit")
        self.assertEqual(len(history._history), 2)
        self.assertEqual(history._history[1].get_state(), "First edit")

    def test_undo(self):
        doc = Document("A")
        history = History(doc) # History: ["A"]
        doc.write("B")
        history.save()        # History: ["A", "AB"]
        doc.write("C")
        history.save()        # History: ["A", "AB", "ABC"]

        self.assertEqual(doc.get_content(), "ABC")
        history.undo()        # Restore "AB". History: ["A", "AB"], Redo: ["ABC"]
        self.assertEqual(doc.get_content(), "AB")
        self.assertEqual(len(history._history), 2)
        self.assertEqual(len(history._redo_stack), 1)
        self.assertEqual(history._redo_stack[0].get_state(), "ABC")

        history.undo()        # Restore "A". History: ["A"], Redo: ["ABC", "AB"]
        self.assertEqual(doc.get_content(), "A")
        self.assertEqual(len(history._history), 1)
        self.assertEqual(len(history._redo_stack), 2)

        history.undo()        # Cannot undo further
        self.assertEqual(doc.get_content(), "A") # State remains "A"
        self.assertEqual(len(history._history), 1)
        self.assertEqual(len(history._redo_stack), 2) # Redo stack unchanged

    def test_redo(self):
        doc = Document("A")
        history = History(doc) # History: ["A"]
        doc.write("B")
        history.save()        # History: ["A", "AB"]
        doc.write("C")
        history.save()        # History: ["A", "AB", "ABC"]

        history.undo()        # Back to "AB". History: ["A", "AB"], Redo: ["ABC"]
        history.undo()        # Back to "A". History: ["A"], Redo: ["ABC", "AB"]

        self.assertEqual(doc.get_content(), "A")
        history.redo()        # Restore "AB". History: ["A", "AB"], Redo: ["ABC"]
        self.assertEqual(doc.get_content(), "AB")
        self.assertEqual(len(history._history), 2)
        self.assertEqual(len(history._redo_stack), 1)

        history.redo()        # Restore "ABC". History: ["A", "AB", "ABC"], Redo: []
        self.assertEqual(doc.get_content(), "ABC")
        self.assertEqual(len(history._history), 3)
        self.assertEqual(len(history._redo_stack), 0)

        history.redo()        # Cannot redo further
        self.assertEqual(doc.get_content(), "ABC") # State remains "ABC"
        self.assertEqual(len(history._history), 3)
        self.assertEqual(len(history._redo_stack), 0) # Redo stack unchanged

    def test_save_clears_redo_stack(self):
        doc = Document("A")
        history = History(doc) # History: ["A"]
        doc.write("B")
        history.save()        # History: ["A", "AB"]
        doc.write("C")
        history.save()        # History: ["A", "AB", "ABC"]

        history.undo()        # Back to "AB". History: ["A", "AB"], Redo: ["ABC"]
        self.assertEqual(len(history._redo_stack), 1)

        # Make a new edit after undo
        doc.write("D") # Content becomes "ABD"
        history.save()        # History: ["A", "AB", "ABD"], Redo: [] (cleared)
        self.assertEqual(doc.get_content(), "ABD")
        self.assertEqual(len(history._history), 3)
        self.assertEqual(len(history._redo_stack), 0) # Redo stack should be empty

        # Try redoing - should fail
        history.redo()
        self.assertEqual(doc.get_content(), "ABD") # State remains "ABD"
