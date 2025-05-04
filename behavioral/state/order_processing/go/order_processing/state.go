package order_processing

import "fmt"

// Forward declaration of Order struct to resolve circular dependency
type Order struct {
	items []string
	state OrderState
	orderID string
}

// OrderState defines the interface for different order states.
type OrderState interface {
	AddItem(item string)
	ProcessPayment()
	Ship()
	Deliver()
	Cancel()
	ToString() string
	setOrder(order *Order) // Method to set the context (Order)
}

// --- Base State (Optional but can be useful for common logic) ---
// We'll embed this in concrete states if needed, or just implement the interface directly.

// --- Concrete States --- 

// NewOrderState represents the state when an order is newly created.
type NewOrderState struct {
	order *Order
}

func (s *NewOrderState) setOrder(order *Order) { s.order = order }

func (s *NewOrderState) AddItem(item string) {
	fmt.Printf("Adding item '%s' to the order.\n", item)
	s.order.items = append(s.order.items, item)
}

func (s *NewOrderState) ProcessPayment() {
	if len(s.order.items) == 0 {
		fmt.Println("Cannot process payment: Order is empty.")
		return
	}
	fmt.Println("Processing payment...")
	// Simulate payment success
	fmt.Println("Payment successful. Transitioning to PendingPaymentState.")
	s.order.setState(&PendingPaymentState{order: s.order})
}

func (s *NewOrderState) Ship() {
	fmt.Println("Cannot ship: Payment not processed yet.")
}

func (s *NewOrderState) Deliver() {
	fmt.Println("Cannot deliver: Order not shipped yet.")
}

func (s *NewOrderState) Cancel() {
	fmt.Println("Cancelling the order.")
	s.order.setState(&CancelledState{order: s.order})
}

func (s *NewOrderState) ToString() string {
	return "New"
}

// PendingPaymentState represents the state after payment is processed but before shipping.
type PendingPaymentState struct {
	order *Order
}

func (s *PendingPaymentState) setOrder(order *Order) { s.order = order }

func (s *PendingPaymentState) AddItem(item string) {
	fmt.Println("Cannot add item: Order payment is pending.")
}

func (s *PendingPaymentState) ProcessPayment() {
	fmt.Println("Payment already processed.")
}

func (s *PendingPaymentState) Ship() {
	fmt.Println("Shipping the order...")
	s.order.setState(&ShippedState{order: s.order})
}

func (s *PendingPaymentState) Deliver() {
	fmt.Println("Cannot deliver: Order not shipped yet.")
}

func (s *PendingPaymentState) Cancel() {
	fmt.Println("Cancelling the order and issuing refund.")
	// Simulate refund logic here if needed
	s.order.setState(&CancelledState{order: s.order})
}

func (s *PendingPaymentState) ToString() string {
	return "Pending Payment"
}

// ShippedState represents the state after the order has been shipped.
type ShippedState struct {
	order *Order
}

func (s *ShippedState) setOrder(order *Order) { s.order = order }

func (s *ShippedState) AddItem(item string) {
	fmt.Println("Cannot add item: Order already shipped.")
}

func (s *ShippedState) ProcessPayment() {
	fmt.Println("Cannot process payment: Order already shipped.")
}

func (s *ShippedState) Ship() {
	fmt.Println("Order already shipped.")
}

func (s *ShippedState) Deliver() {
	fmt.Println("Delivering the order...")
	s.order.setState(&DeliveredState{order: s.order})
}

func (s *ShippedState) Cancel() {
	fmt.Println("Cannot cancel: Order already shipped.")
}

func (s *ShippedState) ToString() string {
	return "Shipped"
}

// DeliveredState represents the final state after the order is delivered.
type DeliveredState struct {
	order *Order
}

func (s *DeliveredState) setOrder(order *Order) { s.order = order }

func (s *DeliveredState) AddItem(item string) {
	fmt.Println("Cannot add item: Order already delivered.")
}

func (s *DeliveredState) ProcessPayment() {
	fmt.Println("Cannot process payment: Order already delivered.")
}

func (s *DeliveredState) Ship() {
	fmt.Println("Cannot ship: Order already delivered.")
}

func (s *DeliveredState) Deliver() {
	fmt.Println("Order already delivered.")
}

func (s *DeliveredState) Cancel() {
	fmt.Println("Cannot cancel: Order already delivered.")
}

func (s *DeliveredState) ToString() string {
	return "Delivered"
}

// CancelledState represents the state when an order is cancelled.
type CancelledState struct {
	order *Order
}

func (s *CancelledState) setOrder(order *Order) { s.order = order }

func (s *CancelledState) AddItem(item string) {
	fmt.Println("Cannot add item: Order is cancelled.")
}

func (s *CancelledState) ProcessPayment() {
	fmt.Println("Cannot process payment: Order is cancelled.")
}

func (s *CancelledState) Ship() {
	fmt.Println("Cannot ship: Order is cancelled.")
}

func (s *CancelledState) Deliver() {
	fmt.Println("Cannot deliver: Order is cancelled.")
}

func (s *CancelledState) Cancel() {
	fmt.Println("Order already cancelled.")
}

func (s *CancelledState) ToString() string {
	return "Cancelled"
}
