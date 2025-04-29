import abc
from typing import List

# --- Component Interface ---
class FileSystemComponent(abc.ABC):
    """The common interface for both leaves (files) and composites (directories)."""

    def __init__(self, name: str):
        self._name = name

    def get_name(self) -> str:
        return self._name

    @abc.abstractmethod
    def get_size(self) -> int:
        """Returns the size of the component in bytes."""
        pass

    @abc.abstractmethod
    def display(self, indent: str = "") -> None:
        """Displays the component's structure."""
        pass

    # Optional: Methods for managing children (might raise errors in Leaf)
    def add(self, component: 'FileSystemComponent') -> None:
        raise NotImplementedError("Cannot add to this component")

    def remove(self, component: 'FileSystemComponent') -> None:
        raise NotImplementedError("Cannot remove from this component")

    def get_child(self, index: int) -> 'FileSystemComponent':
        raise NotImplementedError("Cannot get child from this component")


# --- Leaf Class ---
class File(FileSystemComponent):
    """Represents a leaf object (a file) in the composition."""

    def __init__(self, name: str, size: int):
        super().__init__(name)
        self._size = size

    def get_size(self) -> int:
        return self._size

    def display(self, indent: str = "") -> None:
        print(f"{indent}- {self.get_name()} ({self.get_size()} bytes)")


# --- Composite Class ---
class Directory(FileSystemComponent):
    """Represents a composite object (a directory) that can contain other components."""

    def __init__(self, name: str):
        super().__init__(name)
        self._children: List[FileSystemComponent] = []

    def add(self, component: FileSystemComponent) -> None:
        """Adds a child component (file or subdirectory)."""
        self._children.append(component)

    def remove(self, component: FileSystemComponent) -> None:
        """Removes a child component."""
        self._children.remove(component)

    def get_child(self, index: int) -> FileSystemComponent:
        """Gets a specific child component."""
        return self._children[index]

    def get_size(self) -> int:
        """Calculates the total size by summing the sizes of all children."""
        total_size = 0
        for child in self._children:
            total_size += child.get_size()
        return total_size

    def display(self, indent: str = "") -> None:
        """Displays the directory and recursively displays its children."""
        print(f"{indent}+ {self.get_name()} ({self.get_size()} bytes total)")
        for child in self._children:
            child.display(indent + "  ")
