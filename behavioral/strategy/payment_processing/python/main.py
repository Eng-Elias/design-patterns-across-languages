from strategy import CreditCardPayment, PayPalPayment, BitcoinPayment
from context import PaymentContext

def main():
    """Demonstrates the use of the Strategy pattern for payment processing."""
    print("--- Python Strategy Pattern Demo: Payment Processing ---")

    # Create concrete strategies
    credit_card = CreditCardPayment("1234567890123456", "12/25", "123")
    paypal = PayPalPayment("user@example.com")
    bitcoin = BitcoinPayment("1AbcDeFgHiJkLmNoPqRsTuVwXyZ") # Example Bitcoin address

    # Create context with initial strategy (Credit Card)
    context = PaymentContext(credit_card)

    # Process payment using the current strategy
    context.process_payment(100.00)

    # Change strategy to PayPal
    context.set_strategy(paypal)
    context.process_payment(50.50)

    # Change strategy to Bitcoin
    context.set_strategy(bitcoin)
    context.process_payment(250.75)

    print("\n--- Demo Finished ---")

if __name__ == "__main__":
    main()
