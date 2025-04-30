from abc import ABC, abstractmethod
import time

# --- Subject Interface --- #

class Image(ABC):
    """Subject Interface: Declares the common interface for RealSubject and Proxy."""
    @abstractmethod
    def display(self) -> None:
        """Method to display the image."""
        pass

    @abstractmethod
    def get_filename(self) -> str:
        """Returns the filename associated with the image."""
        pass

# --- Real Subject --- #

class RealImage(Image):
    """RealSubject: Defines the real object that the proxy represents.
    Loading the image is simulated here.
    """
    def __init__(self, filename: str):
        self._filename = filename
        self._load_from_disk()

    def _load_from_disk(self) -> None:
        """Private helper method to simulate loading the image data."""
        print(f"Loading image: '{self._filename}' from disk... (Simulating delay)")
        # Simulate time-consuming operation
        time.sleep(1.5)
        print(f"Finished loading image: '{self._filename}'")

    def display(self) -> None:
        """Displays the image (after it has been loaded)."""
        print(f"Displaying image: '{self._filename}'")

    def get_filename(self) -> str:
        """Returns the filename."""
        return self._filename

# --- Proxy --- #

class ProxyImage(Image):
    """Proxy: Maintains a reference that lets the proxy access the real subject.
    Implements the same interface as the RealSubject.
    Controls access to the real subject and may be responsible for its creation (lazy loading).
    """
    def __init__(self, filename: str):
        self._filename = filename
        self._real_image: RealImage | None = None # Reference to RealImage, initially None
        print(f"ProxyImage created for: '{self._filename}' (Real image not loaded yet)")

    def display(self) -> None:
        """Handles the display request.
        Loads the RealImage only if it hasn't been loaded yet (Lazy Initialization).
        """
        if self._real_image is None:
            print(f"Proxy for '{self._filename}': Real image needs loading.")
            # Lazy initialization: Create the RealImage object only when needed
            self._real_image = RealImage(self._filename)
        else:
            print(f"Proxy for '{self._filename}': Real image already loaded.")

        # Delegate the display call to the RealImage object
        self._real_image.display()

    def get_filename(self) -> str:
        """Returns the filename without loading the real image."""
        return self._filename

    # Optional: Method to check if loaded without triggering load
    def is_loaded(self) -> bool:
        """Checks if the real image has been loaded."""
        return self._real_image is not None
