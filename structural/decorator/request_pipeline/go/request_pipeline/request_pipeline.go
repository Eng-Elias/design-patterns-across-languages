package request_pipeline

import "fmt"

// --- Types ---
type Request struct {
	Headers map[string]string
	Body    map[string]interface{}
	Path    string
	Method  string
}

type Response struct {
	StatusCode int
	Body       map[string]interface{}
}

// --- Component Interface ---
type RequestHandler interface {
	Handle(request Request) Response
}

// --- Concrete Component ---
type BaseHandler struct{}

func (h *BaseHandler) Handle(request Request) Response {
	// In a real application, this would contain the core business logic
	return Response{
		StatusCode: 200,
		Body: map[string]interface{}{
			"message": "Request processed successfully",
		},
	}
}

// --- Decorator Base Class ---
type Middleware struct {
	handler RequestHandler
}

func NewMiddleware(handler RequestHandler) *Middleware {
	return &Middleware{handler: handler}
}

// --- Concrete Decorators ---
type AuthMiddleware struct {
	*Middleware
}

func NewAuthMiddleware(handler RequestHandler) *AuthMiddleware {
	return &AuthMiddleware{NewMiddleware(handler)}
}

func (m *AuthMiddleware) Handle(request Request) Response {
	// Check if the request has a valid auth token
	if _, ok := request.Headers["Authorization"]; !ok {
		return Response{
			StatusCode: 401,
			Body: map[string]interface{}{
				"error": "Unauthorized",
			},
		}
	}

	// If authenticated, pass to the next handler
	return m.handler.Handle(request)
}

type LoggingMiddleware struct {
	*Middleware
}

func NewLoggingMiddleware(handler RequestHandler) *LoggingMiddleware {
	return &LoggingMiddleware{NewMiddleware(handler)}
}

func (m *LoggingMiddleware) Handle(request Request) Response {
	// Log request details
	println("Request received:", request.Method, request.Path)
	fmt.Printf("Headers: %v \n", request.Headers)
	fmt.Printf("Body: %v \n", request.Body)

	// Process the request
	response := m.handler.Handle(request)

	// Log response
	println("Response:", response.StatusCode)
	fmt.Printf("Body: %v \n", response.Body)

	return response
}

type RateLimitMiddleware struct {
	*Middleware
	maxRequests   int
	requestCount  int
}

func NewRateLimitMiddleware(handler RequestHandler, maxRequests int) *RateLimitMiddleware {
	return &RateLimitMiddleware{
		Middleware:   NewMiddleware(handler),
		maxRequests:  maxRequests,
		requestCount: 0,
	}
}

func (m *RateLimitMiddleware) Handle(request Request) Response {
	m.requestCount++

	if m.requestCount > m.maxRequests {
		return Response{
			StatusCode: 429,
			Body: map[string]interface{}{
				"error": "Too many requests",
			},
		}
	}

	return m.handler.Handle(request)
} 