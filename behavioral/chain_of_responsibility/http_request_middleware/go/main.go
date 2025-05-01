package main

import (
	"fmt"
	"log"

	// Use the correct import path based on your go.mod file
	mw "chain_of_responsibility_pattern_http_request_middleware_go/http_request_middleware"
)

func setupMiddlewareChain() mw.Middleware {
	logger := &mw.LoggingMiddleware{}
	authenticator := &mw.AuthenticationMiddleware{}
	authorizer := &mw.AuthorizationMiddleware{}
	finalHandler := &mw.FinalHandler{}

	// Build chain: Log -> Auth -> Authorize -> Final
	logger.SetNext(authenticator).SetNext(authorizer).SetNext(finalHandler)
	return logger // Return the entry point
}

func main() {
	// Important: Replace the import path above with the one matching your project structure and go.mod file.
	log.Println("Setting up middleware chain...")
	middlewareChain := setupMiddlewareChain()

	fmt.Println("\n--- Simulating a successful request ---")
	request1 := mw.MakeRequest("/admin/dashboard", "valid_token", "admin")
	response1, err1 := middlewareChain.Handle(request1)
	if err1 != nil {
		fmt.Printf("Error processing request 1: %v\n", err1)
	}
	fmt.Printf("Response 1: %+v\n", response1)

	fmt.Println("\n--- Simulating a request with invalid authentication ---")
	request2 := mw.MakeRequest("/user/profile", "invalid_token", "")
	response2, err2 := middlewareChain.Handle(request2)
	if err2 != nil {
		fmt.Printf("Error processing request 2: %v\n", err2)
	}
	fmt.Printf("Response 2: %+v\n", response2)

	fmt.Println("\n--- Simulating a request with insufficient authorization ---")
	request3 := mw.MakeRequest("/admin/settings", "valid_token", "superadmin")
	response3, err3 := middlewareChain.Handle(request3)
	if err3 != nil {
		fmt.Printf("Error processing request 3: %v\n", err3)
	}
	fmt.Printf("Response 3: %+v\n", response3)

	fmt.Println("\n--- Simulating a request requiring no specific role ---")
	request4 := mw.MakeRequest("/public/info", "valid_token", "")
	response4, err4 := middlewareChain.Handle(request4)
	if err4 != nil {
		fmt.Printf("Error processing request 4: %v\n", err4)
	}
	fmt.Printf("Response 4: %+v\n", response4)
}
