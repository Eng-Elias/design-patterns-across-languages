package main

import (
	"fmt"

	"order_processing_pattern_go/order_processing" // Import our local package
)

func main() {
	fmt.Println("--- Go State Pattern Demo: Order Processing ---")

	order := order_processing.NewOrder("ORD123")
	fmt.Printf("\nCurrent Status: %s\n", order.ToString())

	fmt.Println("\n# --- Attempting actions in New state ---")
	order.AddItem("Laptop")
	order.AddItem("Mouse")
	fmt.Printf("Current Status: %s\n", order.ToString())
	order.Ship() // Invalid action
	order.ProcessPayment()
	fmt.Printf("Current Status: %s\n", order.ToString())

	fmt.Println("\n# --- Attempting actions in PendingPayment state ---")
	order.AddItem("Keyboard") // Invalid action
	order.Deliver() // Invalid action
	order.Ship()
	fmt.Printf("Current Status: %s\n", order.ToString())

	fmt.Println("\n# --- Attempting actions in Shipped state ---")
	order.ProcessPayment() // Invalid action
	order.Cancel() // Invalid action
	order.Deliver()
	fmt.Printf("Current Status: %s\n", order.ToString())

	fmt.Println("\n# --- Attempting actions in Delivered state ---")
	order.AddItem("Monitor") // Invalid action
	order.Ship() // Invalid action
	order.Cancel() // Invalid action
	fmt.Printf("Current Status: %s\n", order.ToString())

	fmt.Println("\n# --- Demonstrating Cancellation (from New) ---")
	order2 := order_processing.NewOrder("ORD456")
	order2.AddItem("Webcam")
	fmt.Printf("Current Status: %s\n", order2.ToString())
	order2.Cancel()
	fmt.Printf("Current Status: %s\n", order2.ToString())
	order2.AddItem("Microphone") // Invalid action
	order2.ProcessPayment() // Invalid action

	fmt.Println("\n--- Demo Finished ---")
}