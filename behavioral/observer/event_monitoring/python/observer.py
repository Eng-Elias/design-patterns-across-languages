from abc import ABC, abstractmethod
from event import Event, EventType

class Observer(ABC):
    """Abstract base class for Observers."""
    @abstractmethod
    def update(self, event: Event):
        """Receive update from subject."""
        pass

class LoggerObserver(Observer):
    """Concrete observer that logs events."""
    def __init__(self, name: str):
        self._name = name
        self._log_messages = [] # Store logs for testing

    def update(self, event: Event):
        log_message = f"{event.event_type}: {event.data} at {event.timestamp.strftime('%Y-%m-%d %H:%M:%S')}"
        print(f"LoggerObserver: Received event - {log_message}") # Added print for demo
        self._log_messages.append(str(event)) # Store the event string

    # Helper for testing
    def get_logs(self):
        return self._log_messages

class NotifierObserver(Observer):
    """Concrete observer that sends notifications (simulated)."""
    def __init__(self, name: str):
        self._name = name
        self._notifications = [] # Store notifications for testing

    def update(self, event: Event):
        if event.event_type == EventType.LOG_ERROR or event.event_type == EventType.LOG_CRITICAL:
            notification_message = f"Notification: Critical event - Type: {event.event_type}, Data: {event.data}, Time: {event.timestamp.strftime('%Y-%m-%d %H:%M:%S')}"
            print(f"NotifierObserver: {notification_message}") # Added print for demo
            self._notifications.append(f"Notify: {event.event_type} - {event.data}") # Store notification string

    # Helper for testing
    def get_notifications(self):
        return self._notifications