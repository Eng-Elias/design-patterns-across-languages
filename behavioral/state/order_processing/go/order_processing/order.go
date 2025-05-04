package order_processing

import (
	"fmt"
	"strings"
)

// Order struct definition is now complete (forward declared in state.go)

// NewOrder creates a new order instance.
func NewOrder(orderID string) *Order {
	o := &Order{
		orderID: orderID,
		items:   []string{},
	}
	// Set initial state
	initialState := &NewOrderState{}
	initialState.setOrder(o) // Pass the order instance to the state
	o.state = initialState
	fmt.Printf("Order %s created in state: %s\n", o.orderID, o.state.ToString())
	return o
}

// setState allows changing the Order's state.
func (o *Order) setState(state OrderState) {
	fmt.Printf("Order %s: Transitioning from %s to %s\n", o.orderID, o.state.ToString(), state.ToString())
	o.state = state
	// Ensure the new state also has a reference to the order
	// This might be redundant if states are created with the order reference,
	// but it's a safeguard.
	state.setOrder(o)
}

// GetState returns the current state.
func (o *Order) GetState() OrderState {
	return o.state
}

// Delegate actions to the current state

func (o *Order) AddItem(item string) {
	o.state.AddItem(item)
}

func (o *Order) ProcessPayment() {
	o.state.ProcessPayment()
}

func (o *Order) Ship() {
	o.state.Ship()
}

func (o *Order) Deliver() {
	o.state.Deliver()
}

func (o *Order) Cancel() {
	o.state.Cancel()
}

// ToString provides a string representation of the Order.
func (o *Order) ToString() string {
	return fmt.Sprintf("Order [ID: %s, State: %s, Items: [%s]]", o.orderID, o.state.ToString(), strings.Join(o.items, ", "))
}
