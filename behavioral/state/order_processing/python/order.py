from __future__ import annotations
from state import OrderState, NewOrderState
from typing import List

class Order:
    """The Context class representing an order."""
    def __init__(self, order_id: str):
        self.order_id = order_id
        self.items: List[str] = []
        # Default state is NewOrderState
        self._state: OrderState = NewOrderState(self)
        print(f"Order {self.order_id} created in state: {self._state}")

    def set_state(self, state: OrderState):
        """Allows changing the state."""
        print(f"Order {self.order_id}: Transitioning from {self._state} to {state}")
        self._state = state

    def get_state(self) -> OrderState:
        """Returns the current state."""
        return self._state

    def add_item(self, item: str):
        """Delegate action to the current state."""
        self._state.add_item(item)

    def process_payment(self):
        """Delegate action to the current state."""
        self._state.process_payment()

    def ship(self):
        """Delegate action to the current state."""
        self._state.ship()

    def deliver(self):
        """Delegate action to the current state."""
        self._state.deliver()

    def cancel(self):
        """Delegate action to the current state."""
        self._state.cancel()

    def __str__(self) -> str:
        return f"Order [ID: {self.order_id}, State: {self._state}, Items: {self.items}]"