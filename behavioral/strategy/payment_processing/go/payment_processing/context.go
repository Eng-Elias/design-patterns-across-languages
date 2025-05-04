package payment_processing

import (
	"fmt"
	"log"
	"reflect" // Used to get strategy type name for logging
)

// PaymentContext holds a reference to a payment strategy and delegates work.
type PaymentContext struct {
	strategy PaymentStrategy
}

// NewPaymentContext creates a new context with an initial strategy.
func NewPaymentContext(strategy PaymentStrategy) *PaymentContext {
	log.Printf("Payment context initialized with strategy: %s", getStrategyName(strategy))
	return &PaymentContext{
		strategy: strategy,
	}
}

// SetStrategy allows changing the strategy at runtime.
func (ctx *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	log.Printf("Changing payment strategy from %s to %s", getStrategyName(ctx.strategy), getStrategyName(strategy))
	ctx.strategy = strategy
}

// ProcessPayment delegates the payment processing to the current strategy.
func (ctx *PaymentContext) ProcessPayment(amount float64) {
	fmt.Printf("\nAttempting to process payment of $%.2f...\n", amount)
	ctx.strategy.Pay(amount)
}

// Helper function to get the struct name for logging
func getStrategyName(strategy PaymentStrategy) string {
	if strategy == nil {
		return "<nil>"
	}
	// Use reflection to get the type name
	t := reflect.TypeOf(strategy)
	if t.Kind() == reflect.Ptr {
		return t.Elem().Name() // Get name of the pointed-to type (e.g., CreditCardPayment)
	}
	return t.Name() // Get name of the type itself (less likely for interface implementations)
}
