from __future__ import annotations
from abc import ABC, abstractmethod
import typing

# Forward declaration for type hinting
if typing.TYPE_CHECKING:
    from order import Order

class OrderState(ABC):
    """Abstract base class for all order states."""
    def __init__(self, order: Order):
        self._order = order

    @abstractmethod
    def add_item(self, item: str):
        """Attempt to add an item to the order."""
        pass

    @abstractmethod
    def process_payment(self):
        """Attempt to process payment for the order."""
        pass

    @abstractmethod
    def ship(self):
        """Attempt to ship the order."""
        pass

    @abstractmethod
    def deliver(self):
        """Attempt to mark the order as delivered."""
        pass

    @abstractmethod
    def cancel(self):
        """Attempt to cancel the order."""
        pass

    def __str__(self) -> str:
        return self.__class__.__name__

# --- Concrete States --- #

class NewOrderState(OrderState):
    """State for a newly created order."""
    def add_item(self, item: str):
        print(f"Adding item '{item}' to the order.")
        self._order.items.append(item)

    def process_payment(self):
        if not self._order.items:
            print("Cannot process payment: Order is empty.")
            return
        print("Processing payment...")
        # Simulate payment success
        print("Payment successful. Transitioning to PendingPaymentState.")
        self._order.set_state(PendingPaymentState(self._order))

    def ship(self):
        print("Cannot ship: Order payment not processed yet.")

    def deliver(self):
        print("Cannot deliver: Order not shipped yet.")

    def cancel(self):
        print("Cancelling the new order.")
        self._order.set_state(CancelledState(self._order))

class PendingPaymentState(OrderState):
    """State for an order awaiting shipment after payment."""
    def add_item(self, item: str):
        print("Cannot add items: Order payment has been processed.")

    def process_payment(self):
        print("Payment already processed for this order.")

    def ship(self):
        print("Shipping the order...")
        # Simulate shipping success
        print("Order shipped. Transitioning to ShippedState.")
        self._order.set_state(ShippedState(self._order))

    def deliver(self):
        print("Cannot deliver: Order not shipped yet.")

    def cancel(self):
        print("Cancelling the order. Refunding payment...")
        # Simulate refund
        self._order.set_state(CancelledState(self._order))

class ShippedState(OrderState):
    """State for an order that has been shipped."""
    def add_item(self, item: str):
        print("Cannot add items: Order has been shipped.")

    def process_payment(self):
        print("Payment already processed.")

    def ship(self):
        print("Order already shipped.")

    def deliver(self):
        print("Delivering the order...")
        # Simulate delivery success
        print("Order delivered. Transitioning to DeliveredState.")
        self._order.set_state(DeliveredState(self._order))

    def cancel(self):
        print("Cannot cancel: Order has already been shipped.")

class DeliveredState(OrderState):
    """State for an order that has been delivered."""
    def add_item(self, item: str):
        print("Cannot add items: Order has been delivered.")

    def process_payment(self):
        print("Payment already processed.")

    def ship(self):
        print("Order already shipped and delivered.")

    def deliver(self):
        print("Order already delivered.")

    def cancel(self):
        print("Cannot cancel: Order has already been delivered.")

class CancelledState(OrderState):
    """State for a cancelled order."""
    def add_item(self, item: str):
        print("Cannot add items: Order is cancelled.")

    def process_payment(self):
        print("Cannot process payment: Order is cancelled.")

    def ship(self):
        print("Cannot ship: Order is cancelled.")

    def deliver(self):
        print("Cannot deliver: Order is cancelled.")

    def cancel(self):
        print("Order is already cancelled.")