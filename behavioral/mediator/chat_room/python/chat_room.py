"""Mediator Pattern: Chat Room Example"""

from abc import ABC, abstractmethod
from typing import List, Dict

# --- Mediator Interface ---

class ChatMediator(ABC):
    """
    The Mediator interface declares a method used by components (Users)
    to notify the mediator about various events. The Mediator may react
    to these events and pass the execution to other components.
    """
    @abstractmethod
    def send_message(self, message: str, sender: 'User') -> None:
        pass

    @abstractmethod
    def add_user(self, user: 'User') -> None:
        pass

    @abstractmethod
    def remove_user(self, user: 'User') -> None:
        pass


# --- Colleague Interface ---

class User(ABC):
    """
    The Colleague interface. Users communicate only through the Mediator.
    """
    def __init__(self, mediator: ChatMediator, name: str):
        self._mediator = mediator
        self._name = name

    @property
    def name(self) -> str:
        return self._name

    @abstractmethod
    def send(self, message: str) -> None:
        pass

    @abstractmethod
    def receive(self, message: str, sender_name: str) -> None:
        pass


# --- Concrete Mediator ---

class ChatRoom(ChatMediator):
    """
    Concrete Mediator implements cooperative behavior by coordinating Concrete
    Colleagues (ChatUsers). It knows and maintains its colleagues.
    """
    def __init__(self):
        self._users: Dict[str, User] = {}
        print("ChatRoom created.")

    def add_user(self, user: User) -> None:
        if user.name not in self._users:
            self._users[user.name] = user
            print(f"'{user.name}' joined the chat room.")
            # Optionally notify others, but keeping it simple here
        else:
            print(f"User '{user.name}' is already in the chat room.")

    def remove_user(self, user: User) -> None:
        if user.name in self._users:
            del self._users[user.name]
            print(f"'{user.name}' left the chat room.")
            # Optionally notify others
        else:
            print(f"User '{user.name}' is not in the chat room.")

    def send_message(self, message: str, sender: User) -> None:
        print(f"'{sender.name}' sends message: '{message}'")
        for name, user in self._users.items():
            # Send message to everyone except the sender
            if user != sender:
                user.receive(message, sender.name)


# --- Concrete Colleague ---

class ChatUser(User):
    """
    Concrete Colleague communicates with the Mediator when it needs to
    interact with other colleagues.
    """
    def send(self, message: str) -> None:
        print(f"'{self.name}' preparing to send message: '{message}'")
        self._mediator.send_message(message, self)

    def receive(self, message: str, sender_name: str) -> None:
        print(f"'{self.name}' received message from '{sender_name}': '{message}'")