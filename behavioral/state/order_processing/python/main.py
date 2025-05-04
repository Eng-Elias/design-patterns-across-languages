from order import Order

def main():
    """Demonstrates the State pattern with Order processing."""
    print("--- Python State Pattern Demo: Order Processing ---")

    order = Order("ORD123")
    print(f"\nCurrent Status: {order}")

    print("\n# --- Attempting actions in New state ---")
    order.add_item("Laptop")
    order.add_item("Mouse")
    print(f"Current Status: {order}")
    order.ship() # Invalid action
    order.process_payment()
    print(f"Current Status: {order}")

    print("\n# --- Attempting actions in PendingPayment state ---")
    order.add_item("Keyboard") # Invalid action
    order.deliver() # Invalid action
    order.ship()
    print(f"Current Status: {order}")

    print("\n# --- Attempting actions in Shipped state ---")
    order.process_payment() # Invalid action
    order.cancel() # Invalid action
    order.deliver()
    print(f"Current Status: {order}")

    print("\n# --- Attempting actions in Delivered state ---")
    order.add_item("Monitor") # Invalid action
    order.ship() # Invalid action
    order.cancel() # Invalid action
    print(f"Current Status: {order}")

    print("\n# --- Demonstrating Cancellation (from New) ---")
    order2 = Order("ORD456")
    order2.add_item("Webcam")
    print(f"Current Status: {order2}")
    order2.cancel()
    print(f"Current Status: {order2}")
    order2.add_item("Microphone") # Invalid action
    order2.process_payment() # Invalid action

    print("\n--- Demo Finished ---")

if __name__ == "__main__":
    main()