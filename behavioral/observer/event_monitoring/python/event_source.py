from observer import Observer
from event import Event, EventType
from typing import List
import time # For simulating events

class EventSource:
    """The Subject (Observable) class."""
    def __init__(self):
        self._observers: List[Observer] = []

    def attach(self, observer: Observer):
        """Attach an observer to the subject."""
        if observer not in self._observers:
            print(f"EventSource: Attaching {observer.__class__.__name__}")
            self._observers.append(observer)
        else:
            print(f"EventSource: {observer.__class__.__name__} already attached.")

    def detach(self, observer: Observer):
        """Detach an observer from the subject."""
        try:
            print(f"EventSource: Detaching {observer.__class__.__name__}")
            self._observers.remove(observer)
        except ValueError:
            print(f"EventSource: {observer.__class__.__name__} not found for detachment.")

    def _notify(self, event: Event):
        """Notify all observers about an event."""
        print(f"EventSource: Notifying observers about event: {event.event_type}")
        for observer in self._observers:
            observer.update(event)

    def generate_event(self, event_type: EventType, data: dict):
        """Simulate generating an event and notify observers."""
        print(f"\nEventSource: Generating event '{event_type}'...")
        event = Event(event_type=event_type, data=data)
        time.sleep(0.1) # Simulate time passing
        self._notify(event)

    # Helper for testing
    def get_observer_count(self) -> int:
        return len(self._observers)