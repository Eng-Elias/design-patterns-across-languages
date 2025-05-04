from strategy import PaymentStrategy  # Import the abstract strategy class
import logging
import inspect # To get class names for logging

class PaymentContext:
    """
    The Context defines the interface of interest to clients.
    It maintains a reference to one of the Strategy objects.
    """
    def __init__(self, strategy: PaymentStrategy):
        self._strategy = strategy
        logging.info(f"Payment context initialized with strategy: {self._strategy.__class__.__name__}")

    def set_strategy(self, strategy: PaymentStrategy):
        """Allows changing the strategy at runtime."""
        old_strategy_name = self._strategy.__class__.__name__
        new_strategy_name = strategy.__class__.__name__
        logging.info(f"Changing payment strategy from {old_strategy_name} to {new_strategy_name}")
        self._strategy = strategy

    def process_payment(self, amount: float):
        """Delegates the payment processing to the current strategy."""
        print(f"\nAttempting to process payment of ${amount:.2f}...")
        self._strategy.pay(amount)
