// go/main.go
package main

import (
	"fmt"
	"log"
	"os" // Needed to configure logger

	// Adjust the import path based on your go.mod module name
	"strategy_pattern_payment_processing_go/payment_processing"
)

func main() {
	// Configure logger to output to stdout without timestamp/prefix for cleaner demo output matching others
	log.SetOutput(os.Stdout)
	log.SetFlags(0) // Remove default flags (date, time)

	fmt.Println("--- Go Strategy Pattern Demo: Payment Processing ---")

	// Create concrete strategies
	creditCard := payment_processing.NewCreditCardPayment("1234567890123456", "12/25", "123")
	paypal := payment_processing.NewPayPalPayment("user@example.com")
	bitcoin := payment_processing.NewBitcoinPayment("1AbcDeFgHiJkLmNoPqRsTuVwXyZ") // Example Bitcoin address

	// Create context with initial strategy (Credit Card)
	context := payment_processing.NewPaymentContext(creditCard)

	// Process payment using the current strategy
	context.ProcessPayment(100.00)

	// Change strategy to PayPal
	context.SetStrategy(paypal)
	context.ProcessPayment(50.50)

	// Change strategy to Bitcoin
	context.SetStrategy(bitcoin)
	context.ProcessPayment(250.75)

	fmt.Println("\n--- Demo Finished ---")
}
