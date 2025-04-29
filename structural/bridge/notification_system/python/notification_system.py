import abc

# --- Implementation Interface ---
class MessageSender(abc.ABC):
    """Defines the interface for implementation classes (message sending mechanisms)."""
    @abc.abstractmethod
    def send_message(self, subject: str, body: str):
        pass

# --- Concrete Implementations ---
class EmailSender(MessageSender):
    """Concrete implementation for sending messages via Email."""
    def send_message(self, subject: str, body: str):
        print(f"--- Sending Email ---")
        print(f"Subject: {subject}")
        print(f"Body: {body}")
        print(f"---------------------")

class SmsSender(MessageSender):
    """Concrete implementation for sending messages via SMS."""
    def send_message(self, subject: str, body: str):
        # SMS usually doesn't have a subject, so we combine them
        print(f"--- Sending SMS ---")
        print(f"To: [PhoneNumber]") # Placeholder
        print(f"Message: {subject} - {body}")
        print(f"-----------------")

class PushNotificationSender(MessageSender):
    """Concrete implementation for sending messages via Push Notification."""
    def send_message(self, subject: str, body: str):
        print(f"--- Sending Push Notification ---")
        print(f"Title: {subject}")
        print(f"Body: {body}")
        print(f"-------------------------------")

# --- Abstraction ---
class Notification(abc.ABC):
    """Defines the abstraction's interface and maintains a reference to an implementation object."""
    def __init__(self, sender: MessageSender):
        self._sender = sender # Bridge

    @abc.abstractmethod
    def send(self, message: str):
        pass

# --- Refined Abstractions ---
class InfoNotification(Notification):
    """A specific type of notification (refined abstraction)."""
    def send(self, message: str):
        subject = "Info"
        body = f"[INFO] {message}"
        print(f"Preparing '{subject}' notification.")
        self._sender.send_message(subject, body)

class WarningNotification(Notification):
    """Another specific type of notification."""
    def send(self, message: str):
        subject = "Warning"
        body = f"[WARNING] {message}"
        print(f"Preparing '{subject}' notification.")
        self._sender.send_message(subject, body)

class UrgentNotification(Notification):
    """An urgent notification type."""
    def send(self, message: str):
        subject = "** URGENT **"
        body = f"[URGENT ACTION REQUIRED] {message}"
        print(f"Preparing '{subject}' notification.")
        self._sender.send_message(subject, body)
