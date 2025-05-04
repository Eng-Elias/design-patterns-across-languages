import unittest
from io import StringIO
import sys
import logging

# Add the current directory to sys.path to find strategy and context
import os
sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))

from strategy import CreditCardPayment, PayPalPayment, BitcoinPayment
from context import PaymentContext

# Suppress logging output during tests unless specifically needed
logging.disable(logging.CRITICAL)

class TestPaymentProcessing(unittest.TestCase):

    def setUp(self):
        """Redirect stdout to capture print statements."""
        self.held_stdout = sys.stdout
        sys.stdout = self.captured_output = StringIO()
        # You might want to capture logging too if it's relevant
        # self.log_stream = StringIO()
        # self.log_handler = logging.StreamHandler(self.log_stream)
        # logging.getLogger().addHandler(self.log_handler)
        # logging.getLogger().setLevel(logging.INFO) # Ensure logs are captured

    def tearDown(self):
        """Restore stdout."""
        sys.stdout = self.held_stdout
        # logging.getLogger().removeHandler(self.log_handler) # Clean up logger
        logging.disable(logging.NOTSET) # Re-enable logging

    def test_credit_card_payment(self):
        strategy = CreditCardPayment("1111222233334444", "01/28", "987")
        context = PaymentContext(strategy)
        context.process_payment(150.00)
        output = self.captured_output.getvalue()
        self.assertIn("Processing $150.00 using Credit Card: ****-****-****-4444", output)
        self.assertIn("Credit Card payment successful.", output)

    def test_paypal_payment(self):
        strategy = PayPalPayment("test@domain.com")
        context = PaymentContext(strategy)
        context.process_payment(75.25)
        output = self.captured_output.getvalue()
        self.assertIn("Processing $75.25 using PayPal: test@domain.com", output)
        self.assertIn("PayPal payment successful.", output)

    def test_bitcoin_payment(self):
        strategy = BitcoinPayment("3AnotherBitcoinAddressExample")
        context = PaymentContext(strategy)
        context.process_payment(300.00)
        output = self.captured_output.getvalue()
        self.assertIn("Processing $300.00 equivalent in BTC to wallet: 3AnotherBitcoinAddressExample", output)
        self.assertIn("Bitcoin payment initiated", output) # Check for initiation message

    def test_strategy_switching(self):
        credit_card = CreditCardPayment("9999888877776666", "11/26", "555")
        paypal = PayPalPayment("switch@test.org")
        bitcoin = BitcoinPayment("1SwitchAddressExample")

        context = PaymentContext(credit_card)
        context.process_payment(10.00) # Pay with CC

        context.set_strategy(paypal)
        context.process_payment(20.00) # Pay with PayPal

        context.set_strategy(bitcoin)
        context.process_payment(30.00) # Pay with Bitcoin

        output = self.captured_output.getvalue()
        # Check if all payment methods were used
        self.assertIn("****-****-****-6666", output)
        self.assertIn("switch@test.org", output)
        self.assertIn("1SwitchAddressExample", output)

        # Check if the last payment was Bitcoin
        # Simple check: find the last "Processing" line
        last_processing_line = [line for line in output.strip().split('\n') if "Processing $" in line][-1]
        self.assertIn("BTC", last_processing_line)
        self.assertIn("$30.00", last_processing_line)

if __name__ == '__main__':
    unittest.main()
