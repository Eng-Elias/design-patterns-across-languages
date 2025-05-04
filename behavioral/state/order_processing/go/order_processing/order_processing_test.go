package order_processing

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// setup redirects stdout to prevent test logs from cluttering output
func setup(t *testing.T) func() {
	// Keep backup of the real stdout
	old := os.Stdout
	// Create a pipe to capture stdout
	_, w, _ := os.Pipe()
	os.Stdout = w
	log.SetOutput(w) // Also redirect standard logger if used

	// Return a function to restore stdout and close the write pipe
	return func() {
		w.Close()
		os.Stdout = old // Restore real stdout
		log.SetOutput(old)
		// Optionally read and print captured output for debugging:
		// out, _ := io.ReadAll(r)
		// fmt.Println("Captured Output:\n", string(out))
	}
}

// assertState checks if the order is in the expected state type.
func assertState(t *testing.T, order *Order, expectedStateType interface{}, msgAndArgs ...interface{}) bool {
	return assert.IsType(t, expectedStateType, order.GetState(), msgAndArgs...)
}

func TestOrderProcessing_InitialState(t *testing.T) {
	cleanup := setup(t) // Suppress console output for this test
	defer cleanup()

	order := NewOrder("TEST_INIT")
	assertState(t, order, &NewOrderState{}, "Order should start in NewOrderState")
	assert.Empty(t, order.items, "Initial order items should be empty")
}

func TestOrderProcessing_TransitionsFromNew(t *testing.T) {
	cleanup := setup(t)
	defer cleanup()

	order := NewOrder("TEST_NEW")

	// Add item -> stays New
	order.AddItem("Book")
	assertState(t, order, &NewOrderState{}, "State should remain New after adding item")
	assert.Equal(t, []string{"Book"}, order.items, "Item should be added")

	// Process payment -> transitions to PendingPayment
	order.ProcessPayment()
	assertState(t, order, &PendingPaymentState{}, "State should transition to PendingPayment after payment")

	// Reset and test cancel
	order2 := NewOrder("TEST_NEW_CANCEL")
	order2.Cancel()
	assertState(t, order2, &CancelledState{}, "State should transition to Cancelled after cancel")
}

func TestOrderProcessing_TransitionsFromPendingPayment(t *testing.T) {
	cleanup := setup(t)
	defer cleanup()

	order := NewOrder("TEST_PENDING")
	order.AddItem("Pen")
	order.ProcessPayment() // Now in PendingPaymentState
	assertState(t, order, &PendingPaymentState{}, "Order should be in PendingPaymentState")

	// Ship -> transitions to Shipped
	order.Ship()
	assertState(t, order, &ShippedState{}, "State should transition to Shipped after shipping")

	// Reset and test cancel
	order2 := NewOrder("TEST_PENDING_CANCEL")
	order2.AddItem("Paper")
	order2.ProcessPayment() // PendingPaymentState
	order2.Cancel()
	assertState(t, order2, &CancelledState{}, "State should transition to Cancelled from PendingPayment")
}

func TestOrderProcessing_TransitionsFromShipped(t *testing.T) {
	cleanup := setup(t)
	defer cleanup()

	order := NewOrder("TEST_SHIPPED")
	order.AddItem("Stapler")
	order.ProcessPayment()
	order.Ship() // Now in ShippedState
	assertState(t, order, &ShippedState{}, "Order should be in ShippedState")

	// Deliver -> transitions to Delivered
	order.Deliver()
	assertState(t, order, &DeliveredState{}, "State should transition to Delivered after delivery")
}

func TestOrderProcessing_DeliveredStateIsFinal(t *testing.T) {
	cleanup := setup(t)
	defer cleanup()

	order := NewOrder("TEST_DELIVERED")
	order.AddItem("Eraser")
	order.ProcessPayment()
	order.Ship()
	order.Deliver() // Now in DeliveredState
	assertState(t, order, &DeliveredState{}, "Order should be in DeliveredState")
	initialStatePtr := fmt.Sprintf("%p", order.GetState())

	// Try all actions
	order.AddItem("Ruler")
	order.ProcessPayment()
	order.Ship()
	order.Deliver()
	order.Cancel()
	assertState(t, order, &DeliveredState{}, "State should remain Delivered")
	finalStatePtr := fmt.Sprintf("%p", order.GetState())
	assert.Equal(t, initialStatePtr, finalStatePtr, "State instance should not change in Delivered state")
}

func TestOrderProcessing_CancelledStateIsFinal(t *testing.T) {
	cleanup := setup(t)
	defer cleanup()

	order := NewOrder("TEST_CANCELLED")
	order.Cancel() // Now in CancelledState
	assertState(t, order, &CancelledState{}, "Order should be in CancelledState")
	initialStatePtr := fmt.Sprintf("%p", order.GetState())

	// Try all actions
	order.AddItem("Tape")
	order.ProcessPayment()
	order.Ship()
	order.Deliver()
	order.Cancel()
	assertState(t, order, &CancelledState{}, "State should remain Cancelled")
	finalStatePtr := fmt.Sprintf("%p", order.GetState())
	assert.Equal(t, initialStatePtr, finalStatePtr, "State instance should not change in Cancelled state")
}

func TestOrderProcessing_InvalidActions(t *testing.T) {
	cleanup := setup(t)
	defer cleanup()

	// New state
	order := NewOrder("TEST_INVALID")
	assertState(t, order, &NewOrderState{}, "Initial state check")
	order.Ship() // Invalid
	order.Deliver() // Invalid
	assertState(t, order, &NewOrderState{}, "State unchanged after invalid ship/deliver")
	order.ProcessPayment() // Invalid (no items)
	assertState(t, order, &NewOrderState{}, "State unchanged after invalid payment")
	order.AddItem("Marker")
	order.ProcessPayment() // Valid -> PendingPayment
	assertState(t, order, &PendingPaymentState{}, "State changed to PendingPayment")

	// PendingPayment state
	order.AddItem("Board") // Invalid
	order.ProcessPayment() // Invalid (already paid)
	order.Deliver() // Invalid
	assertState(t, order, &PendingPaymentState{}, "State unchanged after invalid actions in PendingPayment")
	order.Ship() // Valid -> Shipped
	assertState(t, order, &ShippedState{}, "State changed to Shipped")

	// Shipped state
	order.AddItem("Projector") // Invalid
	order.ProcessPayment() // Invalid
	order.Ship() // Invalid (already shipped)
	order.Cancel() // Invalid
	assertState(t, order, &ShippedState{}, "State unchanged after invalid actions in Shipped")
	order.Deliver() // Valid -> Delivered
	assertState(t, order, &DeliveredState{}, "State changed to Delivered")
}
