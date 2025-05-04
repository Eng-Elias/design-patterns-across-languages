from abc import ABC, abstractmethod
import logging

# Configure basic logging
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')

class PaymentStrategy(ABC):
    """
    Abstract Base Class for all payment strategies.
    Declares the 'pay' method that concrete strategies must implement.
    """
    @abstractmethod
    def pay(self, amount: float) -> None:
        """Process the payment of a given amount."""
        pass

class CreditCardPayment(PaymentStrategy):
    """Concrete strategy for Credit Card payments."""
    def __init__(self, card_number: str, expiry_date: str, cvv: str):
        self.card_number = card_number
        self.expiry_date = expiry_date
        self.cvv = cvv
        # Mask card number for logging, show only last 4 digits
        masked_number = f"****-****-****-{self.card_number[-4:]}"
        logging.info(f"Initialized Credit Card Payment for card ending in {self.card_number[-4:]}")

    def pay(self, amount: float) -> None:
        masked_number = f"****-****-****-{self.card_number[-4:]}"
        print(f"Processing ${amount:.2f} using Credit Card: {masked_number}")
        # Simulate payment processing logic
        print("Credit Card payment successful.")
        logging.info(f"Successfully processed ${amount:.2f} via Credit Card {masked_number}")

class PayPalPayment(PaymentStrategy):
    """Concrete strategy for PayPal payments."""
    def __init__(self, email: str):
        self.email = email
        logging.info(f"Initialized PayPal Payment for email: {self.email}")

    def pay(self, amount: float) -> None:
        print(f"Processing ${amount:.2f} using PayPal: {self.email}")
        # Simulate PayPal API call
        print("PayPal payment successful.")
        logging.info(f"Successfully processed ${amount:.2f} via PayPal account {self.email}")

class BitcoinPayment(PaymentStrategy):
    """Concrete strategy for Bitcoin payments."""
    def __init__(self, wallet_address: str):
        self.wallet_address = wallet_address
        logging.info(f"Initialized Bitcoin Payment for wallet: {self.wallet_address[:5]}...{self.wallet_address[-4:]}")

    def pay(self, amount: float) -> None:
        print(f"Processing ${amount:.2f} equivalent in BTC to wallet: {self.wallet_address}")
        # Simulate Bitcoin transaction logic
        print("Bitcoin payment initiated (waiting for confirmation).")
        logging.info(f"Initiated Bitcoin payment of ${amount:.2f} to wallet {self.wallet_address}")
