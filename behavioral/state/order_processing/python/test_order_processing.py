import unittest
from unittest.mock import patch, call  # Using patch to potentially mock print if needed later
from order import Order
from state import NewOrderState, PendingPaymentState, ShippedState, DeliveredState, CancelledState

class TestOrderProcessingState(unittest.TestCase):

    def setUp(self):
        """Set up a new order for each test."""
        # Suppress print output during tests for cleaner results
        patcher = patch('builtins.print')
        self.addCleanup(patcher.stop)
        self.mock_print = patcher.start()
        self.order = Order("TEST1")

    def test_initial_state(self):
        """Test that the order starts in NewOrderState."""
        self.assertIsInstance(self.order.get_state(), NewOrderState)
        self.assertEqual(self.order.items, [])

    def test_new_state_transitions(self):
        """Test transitions from NewOrderState."""
        # Add item -> stays New
        self.order.add_item("Book")
        self.assertIsInstance(self.order.get_state(), NewOrderState)
        self.assertEqual(self.order.items, ["Book"])

        # Process payment -> transitions to PendingPayment
        self.order.process_payment()
        self.assertIsInstance(self.order.get_state(), PendingPaymentState)

        # Reset and test cancel
        self.order = Order("TEST2")
        self.order.cancel()
        self.assertIsInstance(self.order.get_state(), CancelledState)

    def test_pending_payment_state_transitions(self):
        """Test transitions from PendingPaymentState."""
        self.order.add_item("Pen")
        self.order.process_payment() # Now in PendingPaymentState

        # Ship -> transitions to Shipped
        self.order.ship()
        self.assertIsInstance(self.order.get_state(), ShippedState)

        # Reset and test cancel
        self.order = Order("TEST3")
        self.order.add_item("Paper")
        self.order.process_payment() # PendingPaymentState
        self.order.cancel()
        self.assertIsInstance(self.order.get_state(), CancelledState)

    def test_shipped_state_transitions(self):
        """Test transitions from ShippedState."""
        self.order.add_item("Stapler")
        self.order.process_payment()
        self.order.ship() # Now in ShippedState

        # Deliver -> transitions to Delivered
        self.order.deliver()
        self.assertIsInstance(self.order.get_state(), DeliveredState)

    def test_delivered_state_no_transitions(self):
        """Test that DeliveredState is a final state (except potentially returns)."""
        self.order.add_item("Eraser")
        self.order.process_payment()
        self.order.ship()
        self.order.deliver() # Now in DeliveredState

        # Try all actions - should remain Delivered
        current_state = self.order.get_state()
        self.order.add_item("Ruler")
        self.order.process_payment()
        self.order.ship()
        self.order.deliver()
        self.order.cancel()
        self.assertIsInstance(self.order.get_state(), DeliveredState)
        self.assertIs(self.order.get_state(), current_state) # Check it's the same instance

    def test_cancelled_state_no_transitions(self):
        """Test that CancelledState is a final state."""
        self.order.cancel() # Now in CancelledState

        # Try all actions - should remain Cancelled
        current_state = self.order.get_state()
        self.order.add_item("Tape")
        self.order.process_payment()
        self.order.ship()
        self.order.deliver()
        self.order.cancel()
        self.assertIsInstance(self.order.get_state(), CancelledState)
        self.assertIs(self.order.get_state(), current_state) # Check it's the same instance

    def test_invalid_actions(self):
        """Test actions that are invalid in certain states."""
        # New state
        self.assertIsInstance(self.order.get_state(), NewOrderState)
        self.order.ship() # Invalid
        self.order.deliver() # Invalid
        self.assertIsInstance(self.order.get_state(), NewOrderState) # Should not change state

        # Process payment without items
        self.order.process_payment() # Invalid
        self.assertIsInstance(self.order.get_state(), NewOrderState)

        self.order.add_item("Marker")
        self.order.process_payment() # Valid -> PendingPayment
        self.assertIsInstance(self.order.get_state(), PendingPaymentState)

        # PendingPayment state
        self.order.add_item("Board") # Invalid
        self.order.process_payment() # Invalid (already paid)
        self.order.deliver() # Invalid
        self.assertIsInstance(self.order.get_state(), PendingPaymentState)

        self.order.ship() # Valid -> Shipped
        self.assertIsInstance(self.order.get_state(), ShippedState)

        # Shipped state
        self.order.add_item("Projector") # Invalid
        self.order.process_payment() # Invalid
        self.order.ship() # Invalid (already shipped)
        self.order.cancel() # Invalid
        self.assertIsInstance(self.order.get_state(), ShippedState)

        self.order.deliver() # Valid -> Delivered
        self.assertIsInstance(self.order.get_state(), DeliveredState)


if __name__ == '__main__':
    unittest.main()
