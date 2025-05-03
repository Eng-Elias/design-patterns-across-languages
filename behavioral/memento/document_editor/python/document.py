from memento import Memento

class Document:
    """
    The Originator holds some important state that may change over time.
    It also defines a method for saving the state inside a memento and
    another method for restoring the state from it.
    """
    def __init__(self, initial_content: str = ""):
        self._content = initial_content
        print(f"Document initialized with: '{self._content}'")

    def write(self, text: str):
        """Appends text to the document content."""
        print(f"Writing: '{text}'")
        self._content += text
        print(f"Current content: '{self._content}'")

    def set_content(self, content: str):
        """Sets the document content directly (used for restoring)."""
        self._content = content
        print(f"Content set to: '{self._content}'")

    def get_content(self) -> str:
        """Returns the current content."""
        return self._content

    def save(self) -> Memento:
        """Saves the current state inside a memento."""
        memento = Memento(self._content)
        print(f"Saving state: '{self._content}'")
        return memento

    def restore(self, memento: Memento):
        """Restores the Originator's state from a memento object."""
        self._content = memento.get_state()
        print(f"Restoring state to: '{self._content}'")
