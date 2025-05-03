class Memento:
    """
    The Memento interface provides a way to retrieve the memento's metadata,
    such as creation date or name. However, it doesn't expose the
    Originator's state.
    """
    def __init__(self, state: str):
        # In Python, there's no strict way to prevent access,
        # but conceptually, only Document should access _state directly.
        # Caretaker (History) should not.
        self._state = state

    def get_state(self) -> str:
        """
        The Originator uses this method when restoring its state.
        """
        return self._state
