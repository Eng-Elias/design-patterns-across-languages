package main

import (
	"decorator_pattern_request_pipeline_go/request_pipeline"
	"fmt"
)

func printResponse(response request_pipeline.Response) {
	fmt.Printf("Status Code: %d\n", response.StatusCode)
	fmt.Printf("Body: %v\n\n", response.Body)
}

func main() {
	// Create a sample request
	request := request_pipeline.Request{
		Headers: map[string]string{"Content-Type": "application/json"},
		Body:    map[string]interface{}{"data": "test"},
		Path:    "/api/test",
		Method:  "POST",
	}

	fmt.Println("--- Simple Request Processing ---")
	// Process request with base handler only
	baseHandler := &request_pipeline.BaseHandler{}
	response := baseHandler.Handle(request)
	printResponse(response)

	fmt.Println("--- Request with Auth Middleware ---")
	// Process request with auth middleware
	authHandler := request_pipeline.NewAuthMiddleware(baseHandler)
	response = authHandler.Handle(request)
	printResponse(response)

	// Add authorization header and try again
	request.Headers["Authorization"] = "Bearer valid-token"
	response = authHandler.Handle(request)
	printResponse(response)

	fmt.Println("--- Request with Rate Limiting ---")
	// Process request with rate limiting
	rateLimitHandler := request_pipeline.NewRateLimitMiddleware(baseHandler, 2)
	for i := 0; i < 3; i++ {
		fmt.Printf("Request %d:\n", i+1)
		response = rateLimitHandler.Handle(request)
		printResponse(response)
	}

	fmt.Println("--- Request with Multiple Middleware ---")
	// Create a chain of middleware: Auth -> RateLimit -> Logging -> BaseHandler
	handler := request_pipeline.NewAuthMiddleware(
		request_pipeline.NewRateLimitMiddleware(
			request_pipeline.NewLoggingMiddleware(
				baseHandler,
			),
			2,
		),
	)

	// Process request through the middleware chain
	response = handler.Handle(request)
	printResponse(response)
} 